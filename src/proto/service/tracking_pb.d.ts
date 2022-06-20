// package: flightTracking
// file: service/tracking.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class NoParams extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): NoParams.AsObject;
    static toObject(includeInstance: boolean, msg: NoParams): NoParams.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: NoParams, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): NoParams;
    static deserializeBinaryFromReader(message: NoParams, reader: jspb.BinaryReader): NoParams;
}

export namespace NoParams {
    export type AsObject = {
    }
}

export class MyService extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): MyService.AsObject;
    static toObject(includeInstance: boolean, msg: MyService): MyService.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: MyService, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): MyService;
    static deserializeBinaryFromReader(message: MyService, reader: jspb.BinaryReader): MyService;
}

export namespace MyService {
    export type AsObject = {
    }


    export class Flight extends jspb.Message { 
        getId(): number;
        setId(value: number): Flight;

        getNumber(): string;
        setNumber(value: string): Flight;


        hasAirline(): boolean;
        clearAirline(): void;
        getAirline(): MyService.Airline | undefined;
        setAirline(value?: MyService.Airline): Flight;


        hasDeparture(): boolean;
        clearDeparture(): void;
        getDeparture(): MyService.Airport | undefined;
        setDeparture(value?: MyService.Airport): Flight;


        hasArrival(): boolean;
        clearArrival(): void;
        getArrival(): MyService.Airport | undefined;
        setArrival(value?: MyService.Airport): Flight;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Flight.AsObject;
        static toObject(includeInstance: boolean, msg: Flight): Flight.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Flight, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Flight;
        static deserializeBinaryFromReader(message: Flight, reader: jspb.BinaryReader): Flight;
    }

    export namespace Flight {
        export type AsObject = {
            id: number,
            number: string,
            airline?: MyService.Airline.AsObject,
            departure?: MyService.Airport.AsObject,
            arrival?: MyService.Airport.AsObject,
        }


        export class Request extends jspb.Message { 
            getId(): number;
            setId(value: number): Request;


            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): Request.AsObject;
            static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
            static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
            static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
            static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): Request;
            static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
        }

        export namespace Request {
            export type AsObject = {
                id: number,
            }
        }

        export class Response extends jspb.Message { 

            hasFlight(): boolean;
            clearFlight(): void;
            getFlight(): MyService.Flight | undefined;
            setFlight(value?: MyService.Flight): Response;


            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): Response.AsObject;
            static toObject(includeInstance: boolean, msg: Response): Response.AsObject;
            static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
            static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
            static serializeBinaryToWriter(message: Response, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): Response;
            static deserializeBinaryFromReader(message: Response, reader: jspb.BinaryReader): Response;
        }

        export namespace Response {
            export type AsObject = {
                flight?: MyService.Flight.AsObject,
            }
        }

        export class NoParams extends jspb.Message { 

            serializeBinary(): Uint8Array;
            toObject(includeInstance?: boolean): NoParams.AsObject;
            static toObject(includeInstance: boolean, msg: NoParams): NoParams.AsObject;
            static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
            static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
            static serializeBinaryToWriter(message: NoParams, writer: jspb.BinaryWriter): void;
            static deserializeBinary(bytes: Uint8Array): NoParams;
            static deserializeBinaryFromReader(message: NoParams, reader: jspb.BinaryReader): NoParams;
        }

        export namespace NoParams {
            export type AsObject = {
            }
        }

    }

    export class Airport extends jspb.Message { 
        getCode(): string;
        setCode(value: string): Airport;

        getName(): string;
        setName(value: string): Airport;

        getCity(): string;
        setCity(value: string): Airport;

        getCountry(): string;
        setCountry(value: string): Airport;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Airport.AsObject;
        static toObject(includeInstance: boolean, msg: Airport): Airport.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Airport, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Airport;
        static deserializeBinaryFromReader(message: Airport, reader: jspb.BinaryReader): Airport;
    }

    export namespace Airport {
        export type AsObject = {
            code: string,
            name: string,
            city: string,
            country: string,
        }
    }

    export class Airline extends jspb.Message { 
        getCode(): string;
        setCode(value: string): Airline;

        getName(): string;
        setName(value: string): Airline;

        getCountry(): string;
        setCountry(value: string): Airline;


        serializeBinary(): Uint8Array;
        toObject(includeInstance?: boolean): Airline.AsObject;
        static toObject(includeInstance: boolean, msg: Airline): Airline.AsObject;
        static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
        static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
        static serializeBinaryToWriter(message: Airline, writer: jspb.BinaryWriter): void;
        static deserializeBinary(bytes: Uint8Array): Airline;
        static deserializeBinaryFromReader(message: Airline, reader: jspb.BinaryReader): Airline;
    }

    export namespace Airline {
        export type AsObject = {
            code: string,
            name: string,
            country: string,
        }
    }

}
