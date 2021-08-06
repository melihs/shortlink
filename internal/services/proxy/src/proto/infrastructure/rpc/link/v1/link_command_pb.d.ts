// package: infrastructure.rpc.link.v1
// file: infrastructure/rpc/link/v1/link_command.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as domain_link_v1_link_pb from "../../../../domain/link/v1/link_pb";

export class AddRequest extends jspb.Message { 

    hasLink(): boolean;
    clearLink(): void;
    getLink(): domain_link_v1_link_pb.Link | undefined;
    setLink(value?: domain_link_v1_link_pb.Link): AddRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AddRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AddRequest): AddRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AddRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AddRequest;
    static deserializeBinaryFromReader(message: AddRequest, reader: jspb.BinaryReader): AddRequest;
}

export namespace AddRequest {
    export type AsObject = {
        link?: domain_link_v1_link_pb.Link.AsObject,
    }
}

export class AddResponse extends jspb.Message { 

    hasLink(): boolean;
    clearLink(): void;
    getLink(): domain_link_v1_link_pb.Link | undefined;
    setLink(value?: domain_link_v1_link_pb.Link): AddResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AddResponse.AsObject;
    static toObject(includeInstance: boolean, msg: AddResponse): AddResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AddResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AddResponse;
    static deserializeBinaryFromReader(message: AddResponse, reader: jspb.BinaryReader): AddResponse;
}

export namespace AddResponse {
    export type AsObject = {
        link?: domain_link_v1_link_pb.Link.AsObject,
    }
}

export class UpdateRequest extends jspb.Message { 

    hasLink(): boolean;
    clearLink(): void;
    getLink(): domain_link_v1_link_pb.Link | undefined;
    setLink(value?: domain_link_v1_link_pb.Link): UpdateRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateRequest): UpdateRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateRequest;
    static deserializeBinaryFromReader(message: UpdateRequest, reader: jspb.BinaryReader): UpdateRequest;
}

export namespace UpdateRequest {
    export type AsObject = {
        link?: domain_link_v1_link_pb.Link.AsObject,
    }
}

export class UpdateResponse extends jspb.Message { 

    hasLink(): boolean;
    clearLink(): void;
    getLink(): domain_link_v1_link_pb.Link | undefined;
    setLink(value?: domain_link_v1_link_pb.Link): UpdateResponse;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateResponse.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateResponse): UpdateResponse.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateResponse, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateResponse;
    static deserializeBinaryFromReader(message: UpdateResponse, reader: jspb.BinaryReader): UpdateResponse;
}

export namespace UpdateResponse {
    export type AsObject = {
        link?: domain_link_v1_link_pb.Link.AsObject,
    }
}

export class DeleteRequest extends jspb.Message { 
    getHash(): string;
    setHash(value: string): DeleteRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): DeleteRequest.AsObject;
    static toObject(includeInstance: boolean, msg: DeleteRequest): DeleteRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: DeleteRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): DeleteRequest;
    static deserializeBinaryFromReader(message: DeleteRequest, reader: jspb.BinaryReader): DeleteRequest;
}

export namespace DeleteRequest {
    export type AsObject = {
        hash: string,
    }
}
