// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var service_tracking_pb = require('../service/tracking_pb.js');

function serialize_flightTracking_MyService_Flight_Request(arg) {
  if (!(arg instanceof service_tracking_pb.MyService.Flight.Request)) {
    throw new Error('Expected argument of type flightTracking.MyService.Flight.Request');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flightTracking_MyService_Flight_Request(buffer_arg) {
  return service_tracking_pb.MyService.Flight.Request.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_flightTracking_MyService_Flight_Response(arg) {
  if (!(arg instanceof service_tracking_pb.MyService.Flight.Response)) {
    throw new Error('Expected argument of type flightTracking.MyService.Flight.Response');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_flightTracking_MyService_Flight_Response(buffer_arg) {
  return service_tracking_pb.MyService.Flight.Response.deserializeBinary(new Uint8Array(buffer_arg));
}


var FlightTrackingService = exports.FlightTrackingService = {
  getFlight: {
    path: '/flightTracking.FlightTracking/GetFlight',
    requestStream: false,
    responseStream: false,
    requestType: service_tracking_pb.MyService.Flight.Request,
    responseType: service_tracking_pb.MyService.Flight.Response,
    requestSerialize: serialize_flightTracking_MyService_Flight_Request,
    requestDeserialize: deserialize_flightTracking_MyService_Flight_Request,
    responseSerialize: serialize_flightTracking_MyService_Flight_Response,
    responseDeserialize: deserialize_flightTracking_MyService_Flight_Response,
  },
  listFlights: {
    path: '/flightTracking.FlightTracking/ListFlights',
    requestStream: false,
    responseStream: true,
    requestType: service_tracking_pb.MyService.Flight.Request,
    responseType: service_tracking_pb.MyService.Flight.Response,
    requestSerialize: serialize_flightTracking_MyService_Flight_Request,
    requestDeserialize: deserialize_flightTracking_MyService_Flight_Request,
    responseSerialize: serialize_flightTracking_MyService_Flight_Response,
    responseDeserialize: deserialize_flightTracking_MyService_Flight_Response,
  },
};

exports.FlightTrackingClient = grpc.makeGenericClientConstructor(FlightTrackingService);
