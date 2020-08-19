/**
 * @fileoverview gRPC-Web generated client stub for containers_ai.alameda.v1alpha1.ai_service
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.containers_ai = {};
proto.containers_ai.alameda = {};
proto.containers_ai.alameda.v1alpha1 = {};
proto.containers_ai.alameda.v1alpha1.ai_service = require('./ai_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.containers_ai.alameda.v1alpha1.ai_service.AlamedaAIServiceClient =
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
proto.containers_ai.alameda.v1alpha1.ai_service.AlamedaAIServicePromiseClient =
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
 *   !proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest,
 *   !proto.containers_ai.alameda.v1alpha1.ai_service.RequestResponse>}
 */
const methodDescriptor_AlamedaAIService_CreatePredictionObjects = new grpc.web.MethodDescriptor(
  '/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/CreatePredictionObjects',
  grpc.web.MethodType.UNARY,
  proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest,
  proto.containers_ai.alameda.v1alpha1.ai_service.RequestResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.ai_service.RequestResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest,
 *   !proto.containers_ai.alameda.v1alpha1.ai_service.RequestResponse>}
 */
const methodInfo_AlamedaAIService_CreatePredictionObjects = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containers_ai.alameda.v1alpha1.ai_service.RequestResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.ai_service.RequestResponse.deserializeBinary
);


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containers_ai.alameda.v1alpha1.ai_service.RequestResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containers_ai.alameda.v1alpha1.ai_service.RequestResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containers_ai.alameda.v1alpha1.ai_service.AlamedaAIServiceClient.prototype.createPredictionObjects =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/CreatePredictionObjects',
      request,
      metadata || {},
      methodDescriptor_AlamedaAIService_CreatePredictionObjects,
      callback);
};


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containers_ai.alameda.v1alpha1.ai_service.RequestResponse>}
 *     Promise that resolves to the response
 */
proto.containers_ai.alameda.v1alpha1.ai_service.AlamedaAIServicePromiseClient.prototype.createPredictionObjects =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/CreatePredictionObjects',
      request,
      metadata || {},
      methodDescriptor_AlamedaAIService_CreatePredictionObjects);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodDescriptor_AlamedaAIService_DeletePredictionObjects = new grpc.web.MethodDescriptor(
  '/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/DeletePredictionObjects',
  grpc.web.MethodType.UNARY,
  proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest,
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest,
 *   !proto.google.protobuf.Empty>}
 */
const methodInfo_AlamedaAIService_DeletePredictionObjects = new grpc.web.AbstractClientBase.MethodInfo(
  google_protobuf_empty_pb.Empty,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_protobuf_empty_pb.Empty.deserializeBinary
);


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.protobuf.Empty)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.protobuf.Empty>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containers_ai.alameda.v1alpha1.ai_service.AlamedaAIServiceClient.prototype.deletePredictionObjects =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/DeletePredictionObjects',
      request,
      metadata || {},
      methodDescriptor_AlamedaAIService_DeletePredictionObjects,
      callback);
};


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.protobuf.Empty>}
 *     Promise that resolves to the response
 */
proto.containers_ai.alameda.v1alpha1.ai_service.AlamedaAIServicePromiseClient.prototype.deletePredictionObjects =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/DeletePredictionObjects',
      request,
      metadata || {},
      methodDescriptor_AlamedaAIService_DeletePredictionObjects);
};


module.exports = proto.containers_ai.alameda.v1alpha1.ai_service;

