/**
 * @fileoverview gRPC-Web generated client stub for testing
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.testing = require('./echo_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.testing.EchoServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.testing.EchoServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.testing.EchoRequest,
 *   !proto.testing.EchoResponse>}
 */
const methodDescriptor_EchoService_Echo = new grpc.web.MethodDescriptor(
  '/testing.EchoService/Echo',
  grpc.web.MethodType.UNARY,
  proto.testing.EchoRequest,
  proto.testing.EchoResponse,
  /**
   * @param {!proto.testing.EchoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.testing.EchoResponse.deserializeBinary
);


/**
 * @param {!proto.testing.EchoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.testing.EchoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.testing.EchoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.testing.EchoServiceClient.prototype.echo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/testing.EchoService/Echo',
      request,
      metadata || {},
      methodDescriptor_EchoService_Echo,
      callback);
};


/**
 * @param {!proto.testing.EchoRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.testing.EchoResponse>}
 *     Promise that resolves to the response
 */
proto.testing.EchoServicePromiseClient.prototype.echo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/testing.EchoService/Echo',
      request,
      metadata || {},
      methodDescriptor_EchoService_Echo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.testing.EchoRequest,
 *   !proto.testing.EchoResponse>}
 */
const methodDescriptor_EchoService_EchoAbort = new grpc.web.MethodDescriptor(
  '/testing.EchoService/EchoAbort',
  grpc.web.MethodType.UNARY,
  proto.testing.EchoRequest,
  proto.testing.EchoResponse,
  /**
   * @param {!proto.testing.EchoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.testing.EchoResponse.deserializeBinary
);


/**
 * @param {!proto.testing.EchoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.testing.EchoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.testing.EchoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.testing.EchoServiceClient.prototype.echoAbort =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/testing.EchoService/EchoAbort',
      request,
      metadata || {},
      methodDescriptor_EchoService_EchoAbort,
      callback);
};


/**
 * @param {!proto.testing.EchoRequest} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.testing.EchoResponse>}
 *     Promise that resolves to the response
 */
proto.testing.EchoServicePromiseClient.prototype.echoAbort =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/testing.EchoService/EchoAbort',
      request,
      metadata || {},
      methodDescriptor_EchoService_EchoAbort);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.testing.Empty,
 *   !proto.testing.Empty>}
 */
const methodDescriptor_EchoService_NoOp = new grpc.web.MethodDescriptor(
  '/testing.EchoService/NoOp',
  grpc.web.MethodType.UNARY,
  proto.testing.Empty,
  proto.testing.Empty,
  /**
   * @param {!proto.testing.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.testing.Empty.deserializeBinary
);


/**
 * @param {!proto.testing.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.testing.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.testing.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.testing.EchoServiceClient.prototype.noOp =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/testing.EchoService/NoOp',
      request,
      metadata || {},
      methodDescriptor_EchoService_NoOp,
      callback);
};


/**
 * @param {!proto.testing.Empty} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.testing.Empty>}
 *     Promise that resolves to the response
 */
proto.testing.EchoServicePromiseClient.prototype.noOp =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/testing.EchoService/NoOp',
      request,
      metadata || {},
      methodDescriptor_EchoService_NoOp);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.testing.ServerStreamingEchoRequest,
 *   !proto.testing.ServerStreamingEchoResponse>}
 */
const methodDescriptor_EchoService_ServerStreamingEcho = new grpc.web.MethodDescriptor(
  '/testing.EchoService/ServerStreamingEcho',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.testing.ServerStreamingEchoRequest,
  proto.testing.ServerStreamingEchoResponse,
  /**
   * @param {!proto.testing.ServerStreamingEchoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.testing.ServerStreamingEchoResponse.deserializeBinary
);


/**
 * @param {!proto.testing.ServerStreamingEchoRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.testing.ServerStreamingEchoResponse>}
 *     The XHR Node Readable Stream
 */
proto.testing.EchoServiceClient.prototype.serverStreamingEcho =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/testing.EchoService/ServerStreamingEcho',
      request,
      metadata || {},
      methodDescriptor_EchoService_ServerStreamingEcho);
};


/**
 * @param {!proto.testing.ServerStreamingEchoRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.testing.ServerStreamingEchoResponse>}
 *     The XHR Node Readable Stream
 */
proto.testing.EchoServicePromiseClient.prototype.serverStreamingEcho =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/testing.EchoService/ServerStreamingEcho',
      request,
      metadata || {},
      methodDescriptor_EchoService_ServerStreamingEcho);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.testing.ServerStreamingEchoRequest,
 *   !proto.testing.ServerStreamingEchoResponse>}
 */
const methodDescriptor_EchoService_ServerStreamingEchoAbort = new grpc.web.MethodDescriptor(
  '/testing.EchoService/ServerStreamingEchoAbort',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.testing.ServerStreamingEchoRequest,
  proto.testing.ServerStreamingEchoResponse,
  /**
   * @param {!proto.testing.ServerStreamingEchoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.testing.ServerStreamingEchoResponse.deserializeBinary
);


/**
 * @param {!proto.testing.ServerStreamingEchoRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.testing.ServerStreamingEchoResponse>}
 *     The XHR Node Readable Stream
 */
proto.testing.EchoServiceClient.prototype.serverStreamingEchoAbort =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/testing.EchoService/ServerStreamingEchoAbort',
      request,
      metadata || {},
      methodDescriptor_EchoService_ServerStreamingEchoAbort);
};


/**
 * @param {!proto.testing.ServerStreamingEchoRequest} request The request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.testing.ServerStreamingEchoResponse>}
 *     The XHR Node Readable Stream
 */
proto.testing.EchoServicePromiseClient.prototype.serverStreamingEchoAbort =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/testing.EchoService/ServerStreamingEchoAbort',
      request,
      metadata || {},
      methodDescriptor_EchoService_ServerStreamingEchoAbort);
};


module.exports = proto.testing;

