// package: testing2
// file: echov.proto

import * as jspb from "google-protobuf";

export class Empty extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Empty.AsObject;
  static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Empty;
  static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
  export type AsObject = {
  }
}

export class EchoRequest extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EchoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: EchoRequest): EchoRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EchoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EchoRequest;
  static deserializeBinaryFromReader(message: EchoRequest, reader: jspb.BinaryReader): EchoRequest;
}

export namespace EchoRequest {
  export type AsObject = {
    message: string,
  }
}

export class EchoResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  getMessageCount(): number;
  setMessageCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EchoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: EchoResponse): EchoResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EchoResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EchoResponse;
  static deserializeBinaryFromReader(message: EchoResponse, reader: jspb.BinaryReader): EchoResponse;
}

export namespace EchoResponse {
  export type AsObject = {
    message: string,
    messageCount: number,
  }
}

export class ServerStreamingEchoRequest extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  getMessageCount(): number;
  setMessageCount(value: number): void;

  getMessageInterval(): number;
  setMessageInterval(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServerStreamingEchoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ServerStreamingEchoRequest): ServerStreamingEchoRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ServerStreamingEchoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServerStreamingEchoRequest;
  static deserializeBinaryFromReader(message: ServerStreamingEchoRequest, reader: jspb.BinaryReader): ServerStreamingEchoRequest;
}

export namespace ServerStreamingEchoRequest {
  export type AsObject = {
    message: string,
    messageCount: number,
    messageInterval: number,
  }
}

export class ServerStreamingEchoResponse extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ServerStreamingEchoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ServerStreamingEchoResponse): ServerStreamingEchoResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ServerStreamingEchoResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ServerStreamingEchoResponse;
  static deserializeBinaryFromReader(message: ServerStreamingEchoResponse, reader: jspb.BinaryReader): ServerStreamingEchoResponse;
}

export namespace ServerStreamingEchoResponse {
  export type AsObject = {
    message: string,
  }
}

export class ClientStreamingEchoRequest extends jspb.Message {
  getMessage(): string;
  setMessage(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientStreamingEchoRequest.AsObject;
  static toObject(includeInstance: boolean, msg: ClientStreamingEchoRequest): ClientStreamingEchoRequest.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientStreamingEchoRequest, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientStreamingEchoRequest;
  static deserializeBinaryFromReader(message: ClientStreamingEchoRequest, reader: jspb.BinaryReader): ClientStreamingEchoRequest;
}

export namespace ClientStreamingEchoRequest {
  export type AsObject = {
    message: string,
  }
}

export class ClientStreamingEchoResponse extends jspb.Message {
  getMessageCount(): number;
  setMessageCount(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ClientStreamingEchoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ClientStreamingEchoResponse): ClientStreamingEchoResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ClientStreamingEchoResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ClientStreamingEchoResponse;
  static deserializeBinaryFromReader(message: ClientStreamingEchoResponse, reader: jspb.BinaryReader): ClientStreamingEchoResponse;
}

export namespace ClientStreamingEchoResponse {
  export type AsObject = {
    messageCount: number,
  }
}

