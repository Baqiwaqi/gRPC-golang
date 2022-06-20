// package: flightTracking
// file: service/tracking.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import {handleClientStreamingCall} from "@grpc/grpc-js/build/src/server-call";
import * as service_tracking_pb from "../service/tracking_pb";

interface IFlightTrackingService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getFlight: IFlightTrackingService_IGetFlight;
    listFlights: IFlightTrackingService_IListFlights;
}

interface IFlightTrackingService_IGetFlight extends grpc.MethodDefinition<service_tracking_pb.MyService.Flight.Request, service_tracking_pb.MyService.Flight.Response> {
    path: "/flightTracking.FlightTracking/GetFlight";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_tracking_pb.MyService.Flight.Request>;
    requestDeserialize: grpc.deserialize<service_tracking_pb.MyService.Flight.Request>;
    responseSerialize: grpc.serialize<service_tracking_pb.MyService.Flight.Response>;
    responseDeserialize: grpc.deserialize<service_tracking_pb.MyService.Flight.Response>;
}
interface IFlightTrackingService_IListFlights extends grpc.MethodDefinition<service_tracking_pb.MyService.Flight.Request, service_tracking_pb.MyService.Flight.Response> {
    path: "/flightTracking.FlightTracking/ListFlights";
    requestStream: false;
    responseStream: true;
    requestSerialize: grpc.serialize<service_tracking_pb.MyService.Flight.Request>;
    requestDeserialize: grpc.deserialize<service_tracking_pb.MyService.Flight.Request>;
    responseSerialize: grpc.serialize<service_tracking_pb.MyService.Flight.Response>;
    responseDeserialize: grpc.deserialize<service_tracking_pb.MyService.Flight.Response>;
}

export const FlightTrackingService: IFlightTrackingService;

export interface IFlightTrackingServer {
    getFlight: grpc.handleUnaryCall<service_tracking_pb.MyService.Flight.Request, service_tracking_pb.MyService.Flight.Response>;
    listFlights: grpc.handleServerStreamingCall<service_tracking_pb.MyService.Flight.Request, service_tracking_pb.MyService.Flight.Response>;
}

export interface IFlightTrackingClient {
    getFlight(request: service_tracking_pb.MyService.Flight.Request, callback: (error: grpc.ServiceError | null, response: service_tracking_pb.MyService.Flight.Response) => void): grpc.ClientUnaryCall;
    getFlight(request: service_tracking_pb.MyService.Flight.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_tracking_pb.MyService.Flight.Response) => void): grpc.ClientUnaryCall;
    getFlight(request: service_tracking_pb.MyService.Flight.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_tracking_pb.MyService.Flight.Response) => void): grpc.ClientUnaryCall;
    listFlights(request: service_tracking_pb.MyService.Flight.Request, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<service_tracking_pb.MyService.Flight.Response>;
    listFlights(request: service_tracking_pb.MyService.Flight.Request, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<service_tracking_pb.MyService.Flight.Response>;
}

export class FlightTrackingClient extends grpc.Client implements IFlightTrackingClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: Partial<grpc.ClientOptions>);
    public getFlight(request: service_tracking_pb.MyService.Flight.Request, callback: (error: grpc.ServiceError | null, response: service_tracking_pb.MyService.Flight.Response) => void): grpc.ClientUnaryCall;
    public getFlight(request: service_tracking_pb.MyService.Flight.Request, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_tracking_pb.MyService.Flight.Response) => void): grpc.ClientUnaryCall;
    public getFlight(request: service_tracking_pb.MyService.Flight.Request, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_tracking_pb.MyService.Flight.Response) => void): grpc.ClientUnaryCall;
    public listFlights(request: service_tracking_pb.MyService.Flight.Request, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<service_tracking_pb.MyService.Flight.Response>;
    public listFlights(request: service_tracking_pb.MyService.Flight.Request, metadata?: grpc.Metadata, options?: Partial<grpc.CallOptions>): grpc.ClientReadableStream<service_tracking_pb.MyService.Flight.Response>;
}
