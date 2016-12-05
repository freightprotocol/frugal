// Autogenerated by Frugal Compiler (2.0.0-RC4)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

library variety.src.f_events_scope;

import 'dart:async';

import 'package:thrift/thrift.dart' as thrift;
import 'package:frugal/frugal.dart' as frugal;

import 'package:variety/variety.dart' as t_variety;


const String delimiter = '.';

/// This docstring gets added to the generated code because it has
/// the @ sign. Prefix specifies topic prefix tokens, which can be static or
/// variable.
class EventsPublisher {
  frugal.FPublisherTransport transport;
  frugal.FProtocolFactory protocolFactory;
  Map<String, frugal.FMethod> _methods;
  EventsPublisher(frugal.FScopeProvider provider, [List<frugal.Middleware> middleware]) {
    transport = provider.publisherTransportFactory.getTransport();
    protocolFactory = provider.protocolFactory;
    this._methods = {};
    this._methods['EventCreated'] = new frugal.FMethod(this._publishEventCreated, 'Events', 'publishEventCreated', middleware);
  }

  Future open() {
    return transport.open();
  }

  Future close() {
    return transport.close();
  }

  /// This is a docstring.
  Future publishEventCreated(frugal.FContext ctx, String user, t_variety.Event req) {
    return this._methods['EventCreated']([ctx, user, req]);
  }

  Future _publishEventCreated(frugal.FContext ctx, String user, t_variety.Event req) async {
    var op = "EventCreated";
    var prefix = "foo.${user}.";
    var topic = "${prefix}Events${delimiter}${op}";
    var memoryBuffer = new frugal.TMemoryOutputBuffer(transport.publishSizeLimit);
    var oprot = protocolFactory.getProtocol(memoryBuffer);
    var msg = new thrift.TMessage(op, thrift.TMessageType.CALL, 0);
    oprot.writeRequestHeader(ctx);
    oprot.writeMessageBegin(msg);
    req.write(oprot);
    oprot.writeMessageEnd();
    await transport.publish(topic, memoryBuffer.writeBytes);
  }
}


/// This docstring gets added to the generated code because it has
/// the @ sign. Prefix specifies topic prefix tokens, which can be static or
/// variable.
class EventsSubscriber {
  final frugal.FScopeProvider provider;
  final List<frugal.Middleware> _middleware;

  EventsSubscriber(this.provider, [this._middleware]) {}

  /// This is a docstring.
  Future<frugal.FSubscription> subscribeEventCreated(String user, dynamic onEvent(frugal.FContext ctx, t_variety.Event req)) async {
    var op = "EventCreated";
    var prefix = "foo.${user}.";
    var topic = "${prefix}Events${delimiter}${op}";
    var transport = provider.subscriberTransportFactory.getTransport();
    await transport.subscribe(topic, _recvEventCreated(op, provider.protocolFactory, onEvent));
    return new frugal.FSubscription(topic, transport);
  }

  frugal.FAsyncCallback _recvEventCreated(String op, frugal.FProtocolFactory protocolFactory, dynamic onEvent(frugal.FContext ctx, t_variety.Event req)) {
    frugal.FMethod method = new frugal.FMethod(onEvent, 'Events', 'subscribeEvent', this._middleware);
    callbackEventCreated(thrift.TTransport transport) {
      var iprot = protocolFactory.getProtocol(transport);
      var ctx = iprot.readRequestHeader();
      var tMsg = iprot.readMessageBegin();
      if (tMsg.name != op) {
        thrift.TProtocolUtil.skip(iprot, thrift.TType.STRUCT);
        iprot.readMessageEnd();
        throw new thrift.TApplicationError(
        thrift.TApplicationErrorType.UNKNOWN_METHOD, tMsg.name);
      }
      var req = new t_variety.Event();
      req.read(iprot);
      iprot.readMessageEnd();
      method([ctx, req]);
    }
    return callbackEventCreated;
  }
}

