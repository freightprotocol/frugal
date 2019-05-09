// Autogenerated by Frugal Compiler (3.4.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package intermediate_include

import (
	"fmt"

	"github.com/Workiva/frugal/lib/gopherjs/frugal"
	"github.com/Workiva/frugal/lib/gopherjs/thrift"
)

type FIntermediateFoo interface {
	IntermeidateFoo(ctx frugal.FContext) (err error)
}

type FIntermediateFooClient struct {
	transport       frugal.FTransport
	protocolFactory *frugal.FProtocolFactory
	methods         map[string]*frugal.Method
}

func NewFIntermediateFooClient(provider *frugal.FServiceProvider, middleware ...frugal.ServiceMiddleware) *FIntermediateFooClient {
	methods := make(map[string]*frugal.Method)
	client := &FIntermediateFooClient{
		transport:       provider.GetTransport(),
		protocolFactory: provider.GetProtocolFactory(),
		methods:         methods,
	}
	middleware = append(middleware, provider.GetMiddleware()...)
	methods["intermeidateFoo"] = frugal.NewMethod(client, client.intermeidateFoo, "intermeidateFoo", middleware)
	return client
}

func (f *FIntermediateFooClient) IntermeidateFoo(ctx frugal.FContext) (err error) {
	ret := f.methods["intermeidateFoo"].Invoke([]interface{}{ctx})
	if len(ret) != 1 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 1", len(ret)))
	}
	if ret[0] != nil {
		err = ret[0].(error)
	}
	return err
}

func (f *FIntermediateFooClient) intermeidateFoo(ctx frugal.FContext) (err error) {
	buffer := frugal.NewTMemoryOutputBuffer(f.transport.GetRequestSizeLimit())
	oprot := f.protocolFactory.GetProtocol(buffer)
	if err = oprot.WriteRequestHeader(ctx); err != nil {
		return
	}
	if err = oprot.WriteMessageBegin("intermeidateFoo", thrift.CALL, 0); err != nil {
		return
	}
	args := IntermediateFooIntermeidateFooArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	if err = oprot.Flush(); err != nil {
		return
	}
	var resultTransport thrift.TTransport
	resultTransport, err = f.transport.Request(ctx, buffer.Bytes())
	if err != nil {
		return
	}
	iprot := f.protocolFactory.GetProtocol(resultTransport)
	if err = iprot.ReadResponseHeader(ctx); err != nil {
		return
	}
	method, mTypeId, _, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if method != "intermeidateFoo" {
		err = thrift.NewTApplicationException(frugal.APPLICATION_EXCEPTION_WRONG_METHOD_NAME, "intermeidateFoo failed: wrong method name")
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error0 := thrift.NewTApplicationException(frugal.APPLICATION_EXCEPTION_UNKNOWN, "Unknown Exception")
		var error1 thrift.TApplicationException
		error1, err = error0.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		if error1.TypeId() == frugal.APPLICATION_EXCEPTION_RESPONSE_TOO_LARGE {
			err = thrift.NewTTransportException(frugal.TRANSPORT_EXCEPTION_RESPONSE_TOO_LARGE, error1.Error())
			return
		}
		err = error1
		return
	}
	if mTypeId != thrift.REPLY {
		err = thrift.NewTApplicationException(frugal.APPLICATION_EXCEPTION_INVALID_MESSAGE_TYPE, "intermeidateFoo failed: invalid message type")
		return
	}
	result := IntermediateFooIntermeidateFooResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	return
}

type FIntermediateFooProcessor struct {
	*frugal.FBaseProcessor
}

func NewFIntermediateFooProcessor(handler FIntermediateFoo, middleware ...frugal.ServiceMiddleware) *FIntermediateFooProcessor {
	p := &FIntermediateFooProcessor{frugal.NewFBaseProcessor()}
	p.AddToProcessorMap("intermeidateFoo", &intermediatefooFIntermeidateFoo{frugal.NewFBaseProcessorFunction(p.GetWriteMutex(), frugal.NewMethod(handler, handler.IntermeidateFoo, "IntermeidateFoo", middleware))})
	return p
}

type intermediatefooFIntermeidateFoo struct {
	*frugal.FBaseProcessorFunction
}

func (p *intermediatefooFIntermeidateFoo) Process(ctx frugal.FContext, iprot, oprot *frugal.FProtocol) error {
	args := IntermediateFooIntermeidateFooArgs{}
	var err error
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		p.GetWriteMutex().Lock()
		err = intermediatefooWriteApplicationError(ctx, oprot, frugal.APPLICATION_EXCEPTION_PROTOCOL_ERROR, "intermeidateFoo", err.Error())
		p.GetWriteMutex().Unlock()
		return err
	}

	iprot.ReadMessageEnd()
	result := IntermediateFooIntermeidateFooResult{}
	var err2 error
	ret := p.InvokeMethod([]interface{}{ctx})
	if len(ret) != 1 {
		panic(fmt.Sprintf("Middleware returned %d arguments, expected 1", len(ret)))
	}
	if ret[0] != nil {
		err2 = ret[0].(error)
	}
	if err2 != nil {
		if err3, ok := err2.(thrift.TApplicationException); ok {
			p.GetWriteMutex().Lock()
			oprot.WriteResponseHeader(ctx)
			oprot.WriteMessageBegin("intermeidateFoo", thrift.EXCEPTION, 0)
			err3.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			p.GetWriteMutex().Unlock()
			return nil
		}
		p.GetWriteMutex().Lock()
		err2 := intermediatefooWriteApplicationError(ctx, oprot, frugal.APPLICATION_EXCEPTION_INTERNAL_ERROR, "intermeidateFoo", "Internal error processing intermeidateFoo: "+err2.Error())
		p.GetWriteMutex().Unlock()
		return err2
	}
	p.GetWriteMutex().Lock()
	defer p.GetWriteMutex().Unlock()
	if err2 = oprot.WriteResponseHeader(ctx); err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			intermediatefooWriteApplicationError(ctx, oprot, frugal.APPLICATION_EXCEPTION_RESPONSE_TOO_LARGE, "intermeidateFoo", err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.WriteMessageBegin("intermeidateFoo", thrift.REPLY, 0); err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			intermediatefooWriteApplicationError(ctx, oprot, frugal.APPLICATION_EXCEPTION_RESPONSE_TOO_LARGE, "intermeidateFoo", err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			intermediatefooWriteApplicationError(ctx, oprot, frugal.APPLICATION_EXCEPTION_RESPONSE_TOO_LARGE, "intermeidateFoo", err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			intermediatefooWriteApplicationError(ctx, oprot, frugal.APPLICATION_EXCEPTION_RESPONSE_TOO_LARGE, "intermeidateFoo", err2.Error())
			return nil
		}
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		if frugal.IsErrTooLarge(err2) {
			intermediatefooWriteApplicationError(ctx, oprot, frugal.APPLICATION_EXCEPTION_RESPONSE_TOO_LARGE, "intermeidateFoo", err2.Error())
			return nil
		}
		err = err2
	}
	return err
}

func intermediatefooWriteApplicationError(ctx frugal.FContext, oprot *frugal.FProtocol, type_ int32, method, message string) error {
	x := thrift.NewTApplicationException(type_, message)
	oprot.WriteResponseHeader(ctx)
	oprot.WriteMessageBegin(method, thrift.EXCEPTION, 0)
	x.Write(oprot)
	oprot.WriteMessageEnd()
	oprot.Flush()
	return x
}

type IntermediateFooIntermeidateFooArgs struct {
}

func NewIntermediateFooIntermeidateFooArgs() *IntermediateFooIntermeidateFooArgs {
	return &IntermediateFooIntermeidateFooArgs{}
}

func (p *IntermediateFooIntermeidateFooArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IntermediateFooIntermeidateFooArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("IntermeidateFoo_args"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IntermediateFooIntermeidateFooArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IntermediateFooIntermeidateFooArgs(%+v)", *p)
}

type IntermediateFooIntermeidateFooResult struct {
}

func NewIntermediateFooIntermeidateFooResult() *IntermediateFooIntermeidateFooResult {
	return &IntermediateFooIntermeidateFooResult{}
}

func (p *IntermediateFooIntermeidateFooResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read error: ", p), err)
	}

	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return thrift.PrependError(fmt.Sprintf("%T field %d read error: ", p, fieldId), err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T read struct end error: ", p), err)
	}
	return nil
}

func (p *IntermediateFooIntermeidateFooResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("IntermeidateFoo_result"); err != nil {
		return thrift.PrependError(fmt.Sprintf("%T write struct begin error: ", p), err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return thrift.PrependError("write field stop error: ", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return thrift.PrependError("write struct stop error: ", err)
	}
	return nil
}

func (p *IntermediateFooIntermeidateFooResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("IntermediateFooIntermeidateFooResult(%+v)", *p)
}
