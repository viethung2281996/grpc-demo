/**
 * @fileoverview gRPC-Web generated client stub for compress.service.v1
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.compress = {};
proto.compress.service = {};
proto.compress.service.v1 = require('./user_tracking_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.compress.service.v1.CompressServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

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
proto.compress.service.v1.CompressServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

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
 *   !proto.compress.service.v1.RequestCompress,
 *   !proto.compress.service.v1.ResponseCompress>}
 */
const methodDescriptor_CompressService_Compress = new grpc.web.MethodDescriptor(
  '/compress.service.v1.CompressService/Compress',
  grpc.web.MethodType.UNARY,
  proto.compress.service.v1.RequestCompress,
  proto.compress.service.v1.ResponseCompress,
  /**
   * @param {!proto.compress.service.v1.RequestCompress} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.compress.service.v1.ResponseCompress.deserializeBinary
);


/**
 * @param {!proto.compress.service.v1.RequestCompress} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.compress.service.v1.ResponseCompress)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.compress.service.v1.ResponseCompress>|undefined}
 *     The XHR Node Readable Stream
 */
proto.compress.service.v1.CompressServiceClient.prototype.compress =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/compress.service.v1.CompressService/Compress',
      request,
      metadata || {},
      methodDescriptor_CompressService_Compress,
      callback);
};


/**
 * @param {!proto.compress.service.v1.RequestCompress} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.compress.service.v1.ResponseCompress>}
 *     Promise that resolves to the response
 */
proto.compress.service.v1.CompressServicePromiseClient.prototype.compress =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/compress.service.v1.CompressService/Compress',
      request,
      metadata || {},
      methodDescriptor_CompressService_Compress);
};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?grpc.web.ClientOptions} options
 * @constructor
 * @struct
 * @final
 */
proto.compress.service.v1.DeviceServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

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
proto.compress.service.v1.DeviceServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options.format = 'binary';

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
 *   !proto.compress.service.v1.DeviceInfo,
 *   !proto.compress.service.v1.Device>}
 */
const methodDescriptor_DeviceService_RegisterDevice = new grpc.web.MethodDescriptor(
  '/compress.service.v1.DeviceService/RegisterDevice',
  grpc.web.MethodType.UNARY,
  proto.compress.service.v1.DeviceInfo,
  proto.compress.service.v1.Device,
  /**
   * @param {!proto.compress.service.v1.DeviceInfo} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.compress.service.v1.Device.deserializeBinary
);


/**
 * @param {!proto.compress.service.v1.DeviceInfo} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.RpcError, ?proto.compress.service.v1.Device)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.compress.service.v1.Device>|undefined}
 *     The XHR Node Readable Stream
 */
proto.compress.service.v1.DeviceServiceClient.prototype.registerDevice =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/compress.service.v1.DeviceService/RegisterDevice',
      request,
      metadata || {},
      methodDescriptor_DeviceService_RegisterDevice,
      callback);
};


/**
 * @param {!proto.compress.service.v1.DeviceInfo} request The
 *     request proto
 * @param {?Object<string, string>=} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.compress.service.v1.Device>}
 *     Promise that resolves to the response
 */
proto.compress.service.v1.DeviceServicePromiseClient.prototype.registerDevice =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/compress.service.v1.DeviceService/RegisterDevice',
      request,
      metadata || {},
      methodDescriptor_DeviceService_RegisterDevice);
};


module.exports = proto.compress.service.v1;

