/**
 * @fileoverview gRPC-Web generated client stub for containersai.datahub.keycodes
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var datahub_keycodes_keycodes_pb = require('../../datahub/keycodes/keycodes_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')

var google_rpc_status_pb = require('../../google/rpc/status_pb.js')
const proto = {};
proto.containersai = {};
proto.containersai.datahub = {};
proto.containersai.datahub.keycodes = require('./services_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.containersai.datahub.keycodes.KeycodesServiceClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.containersai.datahub.keycodes.KeycodesServicePromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

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
 *   !proto.containersai.datahub.keycodes.AddKeycodeRequest,
 *   !proto.containersai.datahub.keycodes.AddKeycodeResponse>}
 */
const methodDescriptor_KeycodesService_AddKeycode = new grpc.web.MethodDescriptor(
  '/containersai.datahub.keycodes.KeycodesService/AddKeycode',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.keycodes.AddKeycodeRequest,
  proto.containersai.datahub.keycodes.AddKeycodeResponse,
  /**
   * @param {!proto.containersai.datahub.keycodes.AddKeycodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.keycodes.AddKeycodeResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.keycodes.AddKeycodeRequest,
 *   !proto.containersai.datahub.keycodes.AddKeycodeResponse>}
 */
const methodInfo_KeycodesService_AddKeycode = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.keycodes.AddKeycodeResponse,
  /**
   * @param {!proto.containersai.datahub.keycodes.AddKeycodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.keycodes.AddKeycodeResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.keycodes.AddKeycodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.keycodes.AddKeycodeResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.keycodes.AddKeycodeResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.keycodes.KeycodesServiceClient.prototype.addKeycode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.keycodes.KeycodesService/AddKeycode',
      request,
      metadata || {},
      methodDescriptor_KeycodesService_AddKeycode,
      callback);
};


/**
 * @param {!proto.containersai.datahub.keycodes.AddKeycodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.keycodes.AddKeycodeResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.keycodes.KeycodesServicePromiseClient.prototype.addKeycode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.keycodes.KeycodesService/AddKeycode',
      request,
      metadata || {},
      methodDescriptor_KeycodesService_AddKeycode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.keycodes.ListKeycodesRequest,
 *   !proto.containersai.datahub.keycodes.ListKeycodesResponse>}
 */
const methodDescriptor_KeycodesService_ListKeycodes = new grpc.web.MethodDescriptor(
  '/containersai.datahub.keycodes.KeycodesService/ListKeycodes',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.keycodes.ListKeycodesRequest,
  proto.containersai.datahub.keycodes.ListKeycodesResponse,
  /**
   * @param {!proto.containersai.datahub.keycodes.ListKeycodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.keycodes.ListKeycodesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.keycodes.ListKeycodesRequest,
 *   !proto.containersai.datahub.keycodes.ListKeycodesResponse>}
 */
const methodInfo_KeycodesService_ListKeycodes = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.keycodes.ListKeycodesResponse,
  /**
   * @param {!proto.containersai.datahub.keycodes.ListKeycodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.keycodes.ListKeycodesResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.keycodes.ListKeycodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.keycodes.ListKeycodesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.keycodes.ListKeycodesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.keycodes.KeycodesServiceClient.prototype.listKeycodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.keycodes.KeycodesService/ListKeycodes',
      request,
      metadata || {},
      methodDescriptor_KeycodesService_ListKeycodes,
      callback);
};


/**
 * @param {!proto.containersai.datahub.keycodes.ListKeycodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.keycodes.ListKeycodesResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.keycodes.KeycodesServicePromiseClient.prototype.listKeycodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.keycodes.KeycodesService/ListKeycodes',
      request,
      metadata || {},
      methodDescriptor_KeycodesService_ListKeycodes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.keycodes.DeleteKeycodeRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_KeycodesService_DeleteKeycode = new grpc.web.MethodDescriptor(
  '/containersai.datahub.keycodes.KeycodesService/DeleteKeycode',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.keycodes.DeleteKeycodeRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.keycodes.DeleteKeycodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.keycodes.DeleteKeycodeRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_KeycodesService_DeleteKeycode = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.keycodes.DeleteKeycodeRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.keycodes.DeleteKeycodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.keycodes.KeycodesServiceClient.prototype.deleteKeycode =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.keycodes.KeycodesService/DeleteKeycode',
      request,
      metadata || {},
      methodDescriptor_KeycodesService_DeleteKeycode,
      callback);
};


/**
 * @param {!proto.containersai.datahub.keycodes.DeleteKeycodeRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.keycodes.KeycodesServicePromiseClient.prototype.deleteKeycode =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.keycodes.KeycodesService/DeleteKeycode',
      request,
      metadata || {},
      methodDescriptor_KeycodesService_DeleteKeycode);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.containersai.datahub.keycodes.GenerateRegistrationDataResponse>}
 */
const methodDescriptor_KeycodesService_GenerateRegistrationData = new grpc.web.MethodDescriptor(
  '/containersai.datahub.keycodes.KeycodesService/GenerateRegistrationData',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.containersai.datahub.keycodes.GenerateRegistrationDataResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.keycodes.GenerateRegistrationDataResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.containersai.datahub.keycodes.GenerateRegistrationDataResponse>}
 */
const methodInfo_KeycodesService_GenerateRegistrationData = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.keycodes.GenerateRegistrationDataResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.keycodes.GenerateRegistrationDataResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.keycodes.GenerateRegistrationDataResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.keycodes.GenerateRegistrationDataResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.keycodes.KeycodesServiceClient.prototype.generateRegistrationData =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.keycodes.KeycodesService/GenerateRegistrationData',
      request,
      metadata || {},
      methodDescriptor_KeycodesService_GenerateRegistrationData,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.keycodes.GenerateRegistrationDataResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.keycodes.KeycodesServicePromiseClient.prototype.generateRegistrationData =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.keycodes.KeycodesService/GenerateRegistrationData',
      request,
      metadata || {},
      methodDescriptor_KeycodesService_GenerateRegistrationData);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.keycodes.ActivateRegistrationDataRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_KeycodesService_ActivateRegistrationData = new grpc.web.MethodDescriptor(
  '/containersai.datahub.keycodes.KeycodesService/ActivateRegistrationData',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.keycodes.ActivateRegistrationDataRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.keycodes.ActivateRegistrationDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.keycodes.ActivateRegistrationDataRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_KeycodesService_ActivateRegistrationData = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.keycodes.ActivateRegistrationDataRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.keycodes.ActivateRegistrationDataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.keycodes.KeycodesServiceClient.prototype.activateRegistrationData =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.keycodes.KeycodesService/ActivateRegistrationData',
      request,
      metadata || {},
      methodDescriptor_KeycodesService_ActivateRegistrationData,
      callback);
};


/**
 * @param {!proto.containersai.datahub.keycodes.ActivateRegistrationDataRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.keycodes.KeycodesServicePromiseClient.prototype.activateRegistrationData =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.keycodes.KeycodesService/ActivateRegistrationData',
      request,
      metadata || {},
      methodDescriptor_KeycodesService_ActivateRegistrationData);
};


module.exports = proto.containersai.datahub.keycodes;

