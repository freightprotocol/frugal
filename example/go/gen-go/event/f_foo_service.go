// Autogenerated by Frugal Compiler (1.1.1)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package event

import (
	"bytes"
	"fmt"
	"sync"
	"time"

	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/Workiva/frugal/example/go/gen-go/base"
	"github.com/Workiva/frugal/lib/go"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

type FFoo interface {
	base.FBaseFoo

	// Ping the server.
	Ping(ctx *frugal.FContext) (err error)
	// Blah the server.
	Blah(ctx *frugal.FContext, num int32, Str string, event *Event) (r int64, err error)
	// oneway methods don't receive a response from the server.
	OneWay(ctx *frugal.FContext, id ID, req Request) (err error)
}

type FFooClient struct {
	*base.FBaseFooClient
	transport       frugal.FTransport
	protocolFactory *frugal.FProtocolFactory
	oprot           *frugal.FProtocol
	mu              sync.Mutex
}

func NewFFooClient(t frugal.FTransport, p *frugal.FProtocolFactory) *FFooClient {
	t.SetRegistry(frugal.NewFClientRegistry())
	return &FFooClient{
		FBaseFooClient:  base.NewFBaseFooClient(t, p),
		transport:       t,
		protocolFactory: p,
		oprot:           p.GetProtocol(t),
	}
}

// Ping the server.
func (f *FFooClient) Ping(ctx *frugal.FContext) (err error) {
	errorC := make(chan error, 1)
	resultC := make(chan struct{}, 1)
	if err = f.transport.Register(ctx, f.recvPingHandler(ctx, resultC, errorC)); err != nil {
		return
	}
	defer f.transport.Unregister(ctx)
	f.mu.Lock()
	if err = f.oprot.WriteRequestHeader(ctx); err != nil {
		f.mu.Unlock()
		return
	}
	if err = f.oprot.WriteMessageBegin("ping", thrift.CALL, 0); err != nil {
		f.mu.Unlock()
		return
	}
	args := FooPingArgs{}
	if err = args.Write(f.oprot); err != nil {
		f.mu.Unlock()
		return
	}
	if err = f.oprot.WriteMessageEnd(); err != nil {
		f.mu.Unlock()
		return
	}
	if err = f.oprot.Flush(); err != nil {
		f.mu.Unlock()
		return
	}
	f.mu.Unlock()

	select {
	case err = <-errorC:
	case <-resultC:
	case <-time.After(ctx.Timeout()):
		err = frugal.ErrTimeout
	case <-f.transport.Closed():
		err = frugal.ErrTransportClosed
	}
	return
}

func (f *FFooClient) recvPingHandler(ctx *frugal.FContext, resultC chan<- struct{}, errorC chan<- error) frugal.FAsyncCallback {
	return func(tr thrift.TTransport) error {
		iprot := f.protocolFactory.GetProtocol(tr)
		if err := iprot.ReadResponseHeader(ctx); err != nil {
			errorC <- err
			return err
		}
		method, mTypeId, _, err := iprot.ReadMessageBegin()
		if err != nil {
			errorC <- err
			return err
		}
		if method != "ping" {
			err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "ping failed: wrong method name")
			errorC <- err
			return err
		}
		if mTypeId == thrift.EXCEPTION {
			error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
			var error1 thrift.TApplicationException
			error1, err = error0.Read(iprot)
			if err != nil {
				errorC <- err
				return err
			}
			if err = iprot.ReadMessageEnd(); err != nil {
				errorC <- err
				return err
			}
			if error1.TypeId() == frugal.RESPONSE_TOO_LARGE {
				err = thrift.NewTTransportException(frugal.RESPONSE_TOO_LARGE, "response too large for transport")
				errorC <- err
				return nil
			}
			err = error1
			errorC <- err
			return err
		}
		if mTypeId != thrift.REPLY {
			err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "ping failed: invalid message type")
			errorC <- err
			return err
		}
		result := FooPingResult{}
		if err = result.Read(iprot); err != nil {
			errorC <- err
			return err
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			errorC <- err
			return err
		}
		resultC <- struct{}{}
		return nil
	}
}

// Blah the server.
func (f *FFooClient) Blah(ctx *frugal.FContext, num int32, str string, event *Event) (r int64, err error) {
	errorC := make(chan error, 1)
	resultC := make(chan int64, 1)
	if err = f.transport.Register(ctx, f.recvBlahHandler(ctx, resultC, errorC)); err != nil {
		return
	}
	defer f.transport.Unregister(ctx)
	f.mu.Lock()
	if err = f.oprot.WriteRequestHeader(ctx); err != nil {
		f.mu.Unlock()
		return
	}
	if err = f.oprot.WriteMessageBegin("blah", thrift.CALL, 0); err != nil {
		f.mu.Unlock()
		return
	}
	args := FooBlahArgs{
		Num:   num,
		Str:   str,
		Event: event,
	}
	if err = args.Write(f.oprot); err != nil {
		f.mu.Unlock()
		return
	}
	if err = f.oprot.WriteMessageEnd(); err != nil {
		f.mu.Unlock()
		return
	}
	if err = f.oprot.Flush(); err != nil {
		f.mu.Unlock()
		return
	}
	f.mu.Unlock()

	select {
	case err = <-errorC:
	case r = <-resultC:
	case <-time.After(ctx.Timeout()):
		err = frugal.ErrTimeout
	case <-f.transport.Closed():
		err = frugal.ErrTransportClosed
	}
	return
}

func (f *FFooClient) recvBlahHandler(ctx *frugal.FContext, resultC chan<- int64, errorC chan<- error) frugal.FAsyncCallback {
	return func(tr thrift.TTransport) error {
		iprot := f.protocolFactory.GetProtocol(tr)
		if err := iprot.ReadResponseHeader(ctx); err != nil {
			errorC <- err
			return err
		}
		method, mTypeId, _, err := iprot.ReadMessageBegin()
		if err != nil {
			errorC <- err
			return err
		}
		if method != "blah" {
			err = thrift.NewTApplicationException(thrift.WRONG_METHOD_NAME, "blah failed: wrong method name")
			errorC <- err
			return err
		}
		if mTypeId == thrift.EXCEPTION {
			error0 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
			var error1 thrift.TApplicationException
			error1, err = error0.Read(iprot)
			if err != nil {
				errorC <- err
				return err
			}
			if err = iprot.ReadMessageEnd(); err != nil {
				errorC <- err
				return err
			}
			if error1.TypeId() == frugal.RESPONSE_TOO_LARGE {
				err = thrift.NewTTransportException(frugal.RESPONSE_TOO_LARGE, "response too large for transport")
				errorC <- err
				return nil
			}
			err = error1
			errorC <- err
			return err
		}
		if mTypeId != thrift.REPLY {
			err = thrift.NewTApplicationException(thrift.INVALID_MESSAGE_TYPE_EXCEPTION, "blah failed: invalid message type")
			errorC <- err
			return err
		}
		result := FooBlahResult{}
		if err = result.Read(iprot); err != nil {
			errorC <- err
			return err
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			errorC <- err
			return err
		}
		if result.Awe != nil {
			errorC <- result.Awe
			return nil
		}
		if result.API != nil {
			errorC <- result.API
			return nil
		}
		resultC <- result.GetSuccess()
		return nil
	}
}

// oneway methods don't receive a response from the server.
func (f *FFooClient) OneWay(ctx *frugal.FContext, id ID, req Request) (err error) {
	f.mu.Lock()
	if err = f.oprot.WriteRequestHeader(ctx); err != nil {
		f.mu.Unlock()
		return
	}
	if err = f.oprot.WriteMessageBegin("oneWay", thrift.ONEWAY, 0); err != nil {
		f.mu.Unlock()
		return
	}
	args := FooOneWayArgs{
		ID:  id,
		Req: req,
	}
	if err = args.Write(f.oprot); err != nil {
		f.mu.Unlock()
		return
	}
	if err = f.oprot.WriteMessageEnd(); err != nil {
		f.mu.Unlock()
		return
	}
	if err = f.oprot.Flush(); err != nil {
		f.mu.Unlock()
		return
	}
	f.mu.Unlock()

	return
}

type FFooProcessor struct {
	*base.FBaseFooProcessor
}

func NewFFooProcessor(handler FFoo, middleware ...frugal.ServiceMiddleware) *FFooProcessor {
	p := &FFooProcessor{
		base.NewFBaseFooProcessor(handler, middleware...),
	}
	p.AddToProcessorMap("ping", &fooFPing{handler: frugal.ComposeMiddleware(handler.Ping, middleware), writeMu: p.GetWriteMutex()})
	p.AddToProcessorMap("blah", &fooFBlah{handler: frugal.ComposeMiddleware(handler.Blah, middleware), writeMu: p.GetWriteMutex()})
	p.AddToProcessorMap("oneWay", &fooFOneWay{handler: frugal.ComposeMiddleware(handler.OneWay, middleware), writeMu: p.GetWriteMutex()})
	return p
}

type fooFPing struct {
	handler frugal.InvocationHandler
	writeMu *sync.Mutex
}

func (p *fooFPing) Process(ctx *frugal.FContext, iprot, oprot *frugal.FProtocol) error {
	args := FooPingArgs{}
	var err error
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		p.writeMu.Lock()
		fooWriteApplicationError(ctx, oprot, thrift.PROTOCOL_ERROR, "ping", err.Error())
		p.writeMu.Unlock()
		return err
	}

	iprot.ReadMessageEnd()
	result := FooPingResult{}
	var err2 error
	ret := p.handler("Foo", "Ping", []interface{}{ctx})
	if len(ret) != 1 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 1", len(ret)))
	}
	if ret[0] != nil {
		err2 = ret[0].(error)
	}
	if err2 != nil {
		p.writeMu.Lock()
		fooWriteApplicationError(ctx, oprot, thrift.INTERNAL_ERROR, "ping", "Internal error processing ping: "+err2.Error())
		p.writeMu.Unlock()
		return err2
	}
	p.writeMu.Lock()
	defer p.writeMu.Unlock()
	if err2 = oprot.WriteResponseHeader(ctx); err2 != nil {
		if err2 == frugal.ErrTooLarge {
			fooWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "ping", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.WriteMessageBegin("ping", thrift.REPLY, 0); err2 != nil {
		if err2 == frugal.ErrTooLarge {
			fooWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "ping", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		if err2 == frugal.ErrTooLarge {
			fooWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "ping", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		if err2 == frugal.ErrTooLarge {
			fooWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "ping", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		if err2 == frugal.ErrTooLarge {
			fooWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "ping", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	return err
}

type fooFBlah struct {
	handler frugal.InvocationHandler
	writeMu *sync.Mutex
}

func (p *fooFBlah) Process(ctx *frugal.FContext, iprot, oprot *frugal.FProtocol) error {
	args := FooBlahArgs{}
	var err error
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		p.writeMu.Lock()
		fooWriteApplicationError(ctx, oprot, thrift.PROTOCOL_ERROR, "blah", err.Error())
		p.writeMu.Unlock()
		return err
	}

	iprot.ReadMessageEnd()
	result := FooBlahResult{}
	var err2 error
	var retval int64
	ret := p.handler("Foo", "Blah", []interface{}{ctx, args.Num, args.Str, args.Event})
	if len(ret) != 2 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 2", len(ret)))
	}
	retval = ret[0].(int64)
	if ret[1] != nil {
		err2 = ret[1].(error)
	}
	if err2 != nil {
		switch v := err2.(type) {
		case *AwesomeException:
			result.Awe = v
		case *base.APIException:
			result.API = v
		default:
			p.writeMu.Lock()
			fooWriteApplicationError(ctx, oprot, thrift.INTERNAL_ERROR, "blah", "Internal error processing blah: "+err2.Error())
			p.writeMu.Unlock()
			return err2
		}
	} else {
		result.Success = &retval
	}
	p.writeMu.Lock()
	defer p.writeMu.Unlock()
	if err2 = oprot.WriteResponseHeader(ctx); err2 != nil {
		if err2 == frugal.ErrTooLarge {
			fooWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "blah", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.WriteMessageBegin("blah", thrift.REPLY, 0); err2 != nil {
		if err2 == frugal.ErrTooLarge {
			fooWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "blah", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		if err2 == frugal.ErrTooLarge {
			fooWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "blah", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		if err2 == frugal.ErrTooLarge {
			fooWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "blah", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		if err2 == frugal.ErrTooLarge {
			fooWriteApplicationError(ctx, oprot, frugal.RESPONSE_TOO_LARGE, "blah", "response too large: "+err2.Error())
			return nil
		}
		err = err2
	}
	return err
}

type fooFOneWay struct {
	handler frugal.InvocationHandler
	writeMu *sync.Mutex
}

func (p *fooFOneWay) Process(ctx *frugal.FContext, iprot, oprot *frugal.FProtocol) error {
	args := FooOneWayArgs{}
	var err error
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		return err
	}

	iprot.ReadMessageEnd()
	var err2 error
	ret := p.handler("Foo", "OneWay", []interface{}{ctx, args.ID, args.Req})
	if len(ret) != 1 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 1", len(ret)))
	}
	if ret[0] != nil {
		err2 = ret[0].(error)
	}
	if err2 != nil {
		return err2
	}
	return err
}

func fooWriteApplicationError(ctx *frugal.FContext, oprot *frugal.FProtocol, type_ int32, method, message string) {
	x := thrift.NewTApplicationException(type_, message)
	oprot.WriteResponseHeader(ctx)
	oprot.WriteMessageBegin(method, thrift.EXCEPTION, 0)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
}
