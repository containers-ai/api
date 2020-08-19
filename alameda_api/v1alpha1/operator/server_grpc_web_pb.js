/**
 * @fileoverview gRPC-Web generated client stub for containers_ai.alameda.v1alpha1.operator
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js')

var google_rpc_status_pb = require('../../../google/rpc/status_pb.js')
const proto = {};
proto.containers_ai = {};
proto.containers_ai.alameda = {};
proto.containers_ai.alameda.v1alpha1 = {};
proto.containers_ai.alameda.v1alpha1.operator = require('./server_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServiceClient =
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
proto.containers_ai.alameda.v1alpha1.operator.OperatorServicePromiseClient =
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
 *   !proto.containers_ai.alameda.v1alpha1.operator.ListMetricsRequest,
 *   !proto.containers_ai.alameda.v1alpha1.operator.ListMetricsResponse>}
 */
const methodDescriptor_OperatorService_ListMetrics = new grpc.web.MethodDescriptor(
  '/containers_ai.alameda.v1alpha1.operator.OperatorService/ListMetrics',
  grpc.web.MethodType.UNARY,
  proto.containers_ai.alameda.v1alpha1.operator.ListMetricsRequest,
  proto.containers_ai.alameda.v1alpha1.operator.ListMetricsResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.operator.ListMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containers_ai.alameda.v1alpha1.operator.ListMetricsRequest,
 *   !proto.containers_ai.alameda.v1alpha1.operator.ListMetricsResponse>}
 */
const methodInfo_OperatorService_ListMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containers_ai.alameda.v1alpha1.operator.ListMetricsResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.operator.ListMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containers_ai.alameda.v1alpha1.operator.ListMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServiceClient.prototype.listMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.operator.OperatorService/ListMetrics',
      request,
      metadata || {},
      methodDescriptor_OperatorService_ListMetrics,
      callback);
};


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServicePromiseClient.prototype.listMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.operator.OperatorService/ListMetrics',
      request,
      metadata || {},
      methodDescriptor_OperatorService_ListMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumRequest,
 *   !proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumResponse>}
 */
const methodDescriptor_OperatorService_ListMetricsSum = new grpc.web.MethodDescriptor(
  '/containers_ai.alameda.v1alpha1.operator.OperatorService/ListMetricsSum',
  grpc.web.MethodType.UNARY,
  proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumRequest,
  proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumRequest,
 *   !proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumResponse>}
 */
const methodInfo_OperatorService_ListMetricsSum = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumResponse.deserializeBinary
);


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServiceClient.prototype.listMetricsSum =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.operator.OperatorService/ListMetricsSum',
      request,
      metadata || {},
      methodDescriptor_OperatorService_ListMetricsSum,
      callback);
};


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containers_ai.alameda.v1alpha1.operator.ListMetricsSumResponse>}
 *     Promise that resolves to the response
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServicePromiseClient.prototype.listMetricsSum =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.operator.OperatorService/ListMetricsSum',
      request,
      metadata || {},
      methodDescriptor_OperatorService_ListMetricsSum);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultRequest,
 *   !proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultResponse>}
 */
const methodDescriptor_OperatorService_CreatePredictResult = new grpc.web.MethodDescriptor(
  '/containers_ai.alameda.v1alpha1.operator.OperatorService/CreatePredictResult',
  grpc.web.MethodType.UNARY,
  proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultRequest,
  proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultRequest,
 *   !proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultResponse>}
 */
const methodInfo_OperatorService_CreatePredictResult = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultResponse.deserializeBinary
);


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServiceClient.prototype.createPredictResult =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.operator.OperatorService/CreatePredictResult',
      request,
      metadata || {},
      methodDescriptor_OperatorService_CreatePredictResult,
      callback);
};


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containers_ai.alameda.v1alpha1.operator.CreatePredictResultResponse>}
 *     Promise that resolves to the response
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServicePromiseClient.prototype.createPredictResult =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.operator.OperatorService/CreatePredictResult',
      request,
      metadata || {},
      methodDescriptor_OperatorService_CreatePredictResult);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoRequest,
 *   !proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoResponse>}
 */
const methodDescriptor_OperatorService_GetResourceInfo = new grpc.web.MethodDescriptor(
  '/containers_ai.alameda.v1alpha1.operator.OperatorService/GetResourceInfo',
  grpc.web.MethodType.UNARY,
  proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoRequest,
  proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoRequest,
 *   !proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoResponse>}
 */
const methodInfo_OperatorService_GetResourceInfo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoResponse.deserializeBinary
);


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServiceClient.prototype.getResourceInfo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.operator.OperatorService/GetResourceInfo',
      request,
      metadata || {},
      methodDescriptor_OperatorService_GetResourceInfo,
      callback);
};


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containers_ai.alameda.v1alpha1.operator.GetResourceInfoResponse>}
 *     Promise that resolves to the response
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServicePromiseClient.prototype.getResourceInfo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.operator.OperatorService/GetResourceInfo',
      request,
      metadata || {},
      methodDescriptor_OperatorService_GetResourceInfo);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationRequest,
 *   !proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationResponse>}
 */
const methodDescriptor_OperatorService_GetResourceRecommendation = new grpc.web.MethodDescriptor(
  '/containers_ai.alameda.v1alpha1.operator.OperatorService/GetResourceRecommendation',
  grpc.web.MethodType.UNARY,
  proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationRequest,
  proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationRequest,
 *   !proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationResponse>}
 */
const methodInfo_OperatorService_GetResourceRecommendation = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationResponse,
  /**
   * @param {!proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationResponse.deserializeBinary
);


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServiceClient.prototype.getResourceRecommendation =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.operator.OperatorService/GetResourceRecommendation',
      request,
      metadata || {},
      methodDescriptor_OperatorService_GetResourceRecommendation,
      callback);
};


/**
 * @param {!proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationResponse>}
 *     Promise that resolves to the response
 */
proto.containers_ai.alameda.v1alpha1.operator.OperatorServicePromiseClient.prototype.getResourceRecommendation =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containers_ai.alameda.v1alpha1.operator.OperatorService/GetResourceRecommendation',
      request,
      metadata || {},
      methodDescriptor_OperatorService_GetResourceRecommendation);
};


module.exports = proto.containers_ai.alameda.v1alpha1.operator;

