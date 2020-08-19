/**
 * @fileoverview gRPC-Web generated client stub for containersai.datahub.v1alpha2
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_rpc_status_pb = require('../../google/rpc/status_pb.js')

var datahub_resource_metadata_v1alpha2_metadata_pb = require('../../datahub/resource/metadata/v1alpha2/metadata_pb.js')

var datahub_resource_v1alpha2_resource_pb = require('../../datahub/resource/v1alpha2/resource_pb.js')

var datahub_prediction_v1alpha2_prediction_pb = require('../../datahub/prediction/v1alpha2/prediction_pb.js')

var datahub_metric_v1alpha2_metric_pb = require('../../datahub/metric/v1alpha2/metric_pb.js')

var datahub_recommendation_v1alpha2_recommendation_pb = require('../../datahub/recommendation/v1alpha2/recommendation_pb.js')

var datahub_score_v1alpha2_score_pb = require('../../datahub/score/v1alpha2/score_pb.js')
const proto = {};
proto.containersai = {};
proto.containersai.datahub = {};
proto.containersai.datahub.v1alpha2 = require('./datahub_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient =
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
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient =
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
 *   !proto.containersai.datahub.v1alpha2.ListPodMetricsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListPodMetricsResponse>}
 */
const methodDescriptor_DatahubService_ListPodMetrics = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/ListPodMetrics',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.ListPodMetricsRequest,
  proto.containersai.datahub.v1alpha2.ListPodMetricsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListPodMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListPodMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.v1alpha2.ListPodMetricsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListPodMetricsResponse>}
 */
const methodInfo_DatahubService_ListPodMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.v1alpha2.ListPodMetricsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListPodMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListPodMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListPodMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.v1alpha2.ListPodMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.v1alpha2.ListPodMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.listPodMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListPodMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodMetrics,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListPodMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.v1alpha2.ListPodMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.listPodMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListPodMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.ListNodeMetricsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListNodeMetricsResponse>}
 */
const methodDescriptor_DatahubService_ListNodeMetrics = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/ListNodeMetrics',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.ListNodeMetricsRequest,
  proto.containersai.datahub.v1alpha2.ListNodeMetricsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListNodeMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListNodeMetricsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.v1alpha2.ListNodeMetricsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListNodeMetricsResponse>}
 */
const methodInfo_DatahubService_ListNodeMetrics = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.v1alpha2.ListNodeMetricsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListNodeMetricsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListNodeMetricsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListNodeMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.v1alpha2.ListNodeMetricsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.v1alpha2.ListNodeMetricsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.listNodeMetrics =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListNodeMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodeMetrics,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListNodeMetricsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.v1alpha2.ListNodeMetricsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.listNodeMetrics =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListNodeMetrics',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodeMetrics);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.ListPodsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListPodsResponse>}
 */
const methodDescriptor_DatahubService_ListPods = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/ListPods',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.ListPodsRequest,
  proto.containersai.datahub.v1alpha2.ListPodsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListPodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListPodsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.v1alpha2.ListPodsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListPodsResponse>}
 */
const methodInfo_DatahubService_ListPods = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.v1alpha2.ListPodsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListPodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListPodsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListPodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.v1alpha2.ListPodsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.v1alpha2.ListPodsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.listPods =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListPods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPods,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListPodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.v1alpha2.ListPodsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.listPods =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListPods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPods);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.ListNodesRequest,
 *   !proto.containersai.datahub.v1alpha2.ListNodesResponse>}
 */
const methodDescriptor_DatahubService_ListNodes = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/ListNodes',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.ListNodesRequest,
  proto.containersai.datahub.v1alpha2.ListNodesResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListNodesResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.v1alpha2.ListNodesRequest,
 *   !proto.containersai.datahub.v1alpha2.ListNodesResponse>}
 */
const methodInfo_DatahubService_ListNodes = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.v1alpha2.ListNodesResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListNodesResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.v1alpha2.ListNodesResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.v1alpha2.ListNodesResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.listNodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodes,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.v1alpha2.ListNodesResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.listNodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.ListPodPredictionsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListPodPredictionsResponse>}
 */
const methodDescriptor_DatahubService_ListPodPredictions = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/ListPodPredictions',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.ListPodPredictionsRequest,
  proto.containersai.datahub.v1alpha2.ListPodPredictionsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListPodPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListPodPredictionsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.v1alpha2.ListPodPredictionsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListPodPredictionsResponse>}
 */
const methodInfo_DatahubService_ListPodPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.v1alpha2.ListPodPredictionsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListPodPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListPodPredictionsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListPodPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.v1alpha2.ListPodPredictionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.v1alpha2.ListPodPredictionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.listPodPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListPodPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodPredictions,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListPodPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.v1alpha2.ListPodPredictionsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.listPodPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListPodPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.ListNodePredictionsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListNodePredictionsResponse>}
 */
const methodDescriptor_DatahubService_ListNodePredictions = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/ListNodePredictions',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.ListNodePredictionsRequest,
  proto.containersai.datahub.v1alpha2.ListNodePredictionsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListNodePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListNodePredictionsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.v1alpha2.ListNodePredictionsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListNodePredictionsResponse>}
 */
const methodInfo_DatahubService_ListNodePredictions = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.v1alpha2.ListNodePredictionsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListNodePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListNodePredictionsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListNodePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.v1alpha2.ListNodePredictionsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.v1alpha2.ListNodePredictionsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.listNodePredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListNodePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodePredictions,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListNodePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.v1alpha2.ListNodePredictionsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.listNodePredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListNodePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListNodePredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.ListPodRecommendationsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListPodRecommendationsResponse>}
 */
const methodDescriptor_DatahubService_ListPodRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/ListPodRecommendations',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.ListPodRecommendationsRequest,
  proto.containersai.datahub.v1alpha2.ListPodRecommendationsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListPodRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListPodRecommendationsResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.v1alpha2.ListPodRecommendationsRequest,
 *   !proto.containersai.datahub.v1alpha2.ListPodRecommendationsResponse>}
 */
const methodInfo_DatahubService_ListPodRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.v1alpha2.ListPodRecommendationsResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListPodRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListPodRecommendationsResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListPodRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.v1alpha2.ListPodRecommendationsResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.v1alpha2.ListPodRecommendationsResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.listPodRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListPodRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListPodRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.v1alpha2.ListPodRecommendationsResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.listPodRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListPodRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListPodRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest,
 *   !proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse>}
 */
const methodDescriptor_DatahubService_ListSimulatedSchedulingScores = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/ListSimulatedSchedulingScores',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest,
  proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest,
 *   !proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse>}
 */
const methodInfo_DatahubService_ListSimulatedSchedulingScores = new grpc.web.AbstractClientBase.MethodInfo(
  proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.listSimulatedSchedulingScores =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListSimulatedSchedulingScores',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListSimulatedSchedulingScores,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.listSimulatedSchedulingScores =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/ListSimulatedSchedulingScores',
      request,
      metadata || {},
      methodDescriptor_DatahubService_ListSimulatedSchedulingScores);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.CreatePodsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreatePods = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/CreatePods',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.CreatePodsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreatePodsRequest} request
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
 *   !proto.containersai.datahub.v1alpha2.CreatePodsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreatePods = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreatePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreatePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.createPods =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreatePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePods,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreatePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.createPods =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreatePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePods);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.CreateNodesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNodes = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/CreateNodes',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.CreateNodesRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreateNodesRequest} request
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
 *   !proto.containersai.datahub.v1alpha2.CreateNodesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNodes = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreateNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreateNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.createNodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreateNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodes,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreateNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.createNodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreateNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.CreatePodPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreatePodPredictions = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/CreatePodPredictions',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.CreatePodPredictionsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreatePodPredictionsRequest} request
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
 *   !proto.containersai.datahub.v1alpha2.CreatePodPredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreatePodPredictions = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreatePodPredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreatePodPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.createPodPredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreatePodPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodPredictions,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreatePodPredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.createPodPredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreatePodPredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodPredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.CreateNodePredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateNodePredictions = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/CreateNodePredictions',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.CreateNodePredictionsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreateNodePredictionsRequest} request
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
 *   !proto.containersai.datahub.v1alpha2.CreateNodePredictionsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateNodePredictions = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreateNodePredictionsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreateNodePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.createNodePredictions =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreateNodePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodePredictions,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreateNodePredictionsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.createNodePredictions =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreateNodePredictions',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateNodePredictions);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.CreatePodRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreatePodRecommendations = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/CreatePodRecommendations',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.CreatePodRecommendationsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreatePodRecommendationsRequest} request
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
 *   !proto.containersai.datahub.v1alpha2.CreatePodRecommendationsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreatePodRecommendations = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreatePodRecommendationsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreatePodRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.createPodRecommendations =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreatePodRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodRecommendations,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreatePodRecommendationsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.createPodRecommendations =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreatePodRecommendations',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreatePodRecommendations);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_CreateSimulatedSchedulingScores = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/CreateSimulatedSchedulingScores',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest} request
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
 *   !proto.containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_CreateSimulatedSchedulingScores = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.createSimulatedSchedulingScores =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreateSimulatedSchedulingScores',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateSimulatedSchedulingScores,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.createSimulatedSchedulingScores =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/CreateSimulatedSchedulingScores',
      request,
      metadata || {},
      methodDescriptor_DatahubService_CreateSimulatedSchedulingScores);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.UpdatePodsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_UpdatePods = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/UpdatePods',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.UpdatePodsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.UpdatePodsRequest} request
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
 *   !proto.containersai.datahub.v1alpha2.UpdatePodsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_UpdatePods = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.UpdatePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.UpdatePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.updatePods =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/UpdatePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_UpdatePods,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.UpdatePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.updatePods =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/UpdatePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_UpdatePods);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.UpdateNodesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_UpdateNodes = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/UpdateNodes',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.UpdateNodesRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.UpdateNodesRequest} request
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
 *   !proto.containersai.datahub.v1alpha2.UpdateNodesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_UpdateNodes = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.UpdateNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.UpdateNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.updateNodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/UpdateNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_UpdateNodes,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.UpdateNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.updateNodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/UpdateNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_UpdateNodes);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.DeletePodsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeletePods = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/DeletePods',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.DeletePodsRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.DeletePodsRequest} request
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
 *   !proto.containersai.datahub.v1alpha2.DeletePodsRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeletePods = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.DeletePodsRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.DeletePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.deletePods =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/DeletePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeletePods,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.DeletePodsRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.deletePods =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/DeletePods',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeletePods);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.containersai.datahub.v1alpha2.DeleteNodesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodDescriptor_DatahubService_DeleteNodes = new grpc.web.MethodDescriptor(
  '/containersai.datahub.v1alpha2.DatahubService/DeleteNodes',
  grpc.web.MethodType.UNARY,
  proto.containersai.datahub.v1alpha2.DeleteNodesRequest,
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.DeleteNodesRequest} request
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
 *   !proto.containersai.datahub.v1alpha2.DeleteNodesRequest,
 *   !proto.google.rpc.Status>}
 */
const methodInfo_DatahubService_DeleteNodes = new grpc.web.AbstractClientBase.MethodInfo(
  google_rpc_status_pb.Status,
  /**
   * @param {!proto.containersai.datahub.v1alpha2.DeleteNodesRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  google_rpc_status_pb.Status.deserializeBinary
);


/**
 * @param {!proto.containersai.datahub.v1alpha2.DeleteNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.google.rpc.Status)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.google.rpc.Status>|undefined}
 *     The XHR Node Readable Stream
 */
proto.containersai.datahub.v1alpha2.DatahubServiceClient.prototype.deleteNodes =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/DeleteNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteNodes,
      callback);
};


/**
 * @param {!proto.containersai.datahub.v1alpha2.DeleteNodesRequest} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.google.rpc.Status>}
 *     Promise that resolves to the response
 */
proto.containersai.datahub.v1alpha2.DatahubServicePromiseClient.prototype.deleteNodes =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/containersai.datahub.v1alpha2.DatahubService/DeleteNodes',
      request,
      metadata || {},
      methodDescriptor_DatahubService_DeleteNodes);
};


module.exports = proto.containersai.datahub.v1alpha2;

