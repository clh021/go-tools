// package: testing2
// file: echov.proto

import * as echov_pb from "./echov_pb";
import {grpc} from "@improbable-eng/grpc-web";

type EchoServiceEcho = {
  readonly methodName: string;
  readonly service: typeof EchoService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof echov_pb.EchoRequest;
  readonly responseType: typeof echov_pb.EchoResponse;
};

type EchoServiceEchoAbort = {
  readonly methodName: string;
  readonly service: typeof EchoService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof echov_pb.EchoRequest;
  readonly responseType: typeof echov_pb.EchoResponse;
};

type EchoServiceNoOp = {
  readonly methodName: string;
  readonly service: typeof EchoService;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof echov_pb.Empty;
  readonly responseType: typeof echov_pb.Empty;
};

type EchoServiceServerStreamingEcho = {
  readonly methodName: string;
  readonly service: typeof EchoService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof echov_pb.ServerStreamingEchoRequest;
  readonly responseType: typeof echov_pb.ServerStreamingEchoResponse;
};

type EchoServiceServerStreamingEchoAbort = {
  readonly methodName: string;
  readonly service: typeof EchoService;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof echov_pb.ServerStreamingEchoRequest;
  readonly responseType: typeof echov_pb.ServerStreamingEchoResponse;
};

type EchoServiceClientStreamingEcho = {
  readonly methodName: string;
  readonly service: typeof EchoService;
  readonly requestStream: true;
  readonly responseStream: false;
  readonly requestType: typeof echov_pb.ClientStreamingEchoRequest;
  readonly responseType: typeof echov_pb.ClientStreamingEchoResponse;
};

type EchoServiceFullDuplexEcho = {
  readonly methodName: string;
  readonly service: typeof EchoService;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof echov_pb.EchoRequest;
  readonly responseType: typeof echov_pb.EchoResponse;
};

type EchoServiceHalfDuplexEcho = {
  readonly methodName: string;
  readonly service: typeof EchoService;
  readonly requestStream: true;
  readonly responseStream: true;
  readonly requestType: typeof echov_pb.EchoRequest;
  readonly responseType: typeof echov_pb.EchoResponse;
};

export class EchoService {
  static readonly serviceName: string;
  static readonly Echo: EchoServiceEcho;
  static readonly EchoAbort: EchoServiceEchoAbort;
  static readonly NoOp: EchoServiceNoOp;
  static readonly ServerStreamingEcho: EchoServiceServerStreamingEcho;
  static readonly ServerStreamingEchoAbort: EchoServiceServerStreamingEchoAbort;
  static readonly ClientStreamingEcho: EchoServiceClientStreamingEcho;
  static readonly FullDuplexEcho: EchoServiceFullDuplexEcho;
  static readonly HalfDuplexEcho: EchoServiceHalfDuplexEcho;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class EchoServiceClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  echo(
    requestMessage: echov_pb.EchoRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: echov_pb.EchoResponse|null) => void
  ): UnaryResponse;
  echo(
    requestMessage: echov_pb.EchoRequest,
    callback: (error: ServiceError|null, responseMessage: echov_pb.EchoResponse|null) => void
  ): UnaryResponse;
  echoAbort(
    requestMessage: echov_pb.EchoRequest,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: echov_pb.EchoResponse|null) => void
  ): UnaryResponse;
  echoAbort(
    requestMessage: echov_pb.EchoRequest,
    callback: (error: ServiceError|null, responseMessage: echov_pb.EchoResponse|null) => void
  ): UnaryResponse;
  noOp(
    requestMessage: echov_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: echov_pb.Empty|null) => void
  ): UnaryResponse;
  noOp(
    requestMessage: echov_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: echov_pb.Empty|null) => void
  ): UnaryResponse;
  serverStreamingEcho(requestMessage: echov_pb.ServerStreamingEchoRequest, metadata?: grpc.Metadata): ResponseStream<echov_pb.ServerStreamingEchoResponse>;
  serverStreamingEchoAbort(requestMessage: echov_pb.ServerStreamingEchoRequest, metadata?: grpc.Metadata): ResponseStream<echov_pb.ServerStreamingEchoResponse>;
  clientStreamingEcho(metadata?: grpc.Metadata): RequestStream<echov_pb.ClientStreamingEchoRequest>;
  fullDuplexEcho(metadata?: grpc.Metadata): BidirectionalStream<echov_pb.EchoRequest, echov_pb.EchoResponse>;
  halfDuplexEcho(metadata?: grpc.Metadata): BidirectionalStream<echov_pb.EchoRequest, echov_pb.EchoResponse>;
}

