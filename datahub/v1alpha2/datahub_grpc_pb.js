// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// / This file has messages related to Kubernetes metadata
'use strict';
var grpc = require('grpc');
var datahub_v1alpha2_datahub_pb = require('../../datahub/v1alpha2/datahub_pb.js');
var google_rpc_status_pb = require('../../google/rpc/status_pb.js');
var datahub_resource_metadata_v1alpha2_metadata_pb = require('../../datahub/resource/metadata/v1alpha2/metadata_pb.js');
var datahub_resource_v1alpha2_resource_pb = require('../../datahub/resource/v1alpha2/resource_pb.js');
var datahub_prediction_v1alpha2_prediction_pb = require('../../datahub/prediction/v1alpha2/prediction_pb.js');
var datahub_metric_v1alpha2_metric_pb = require('../../datahub/metric/v1alpha2/metric_pb.js');
var datahub_recommendation_v1alpha2_recommendation_pb = require('../../datahub/recommendation/v1alpha2/recommendation_pb.js');
var datahub_score_v1alpha2_score_pb = require('../../datahub/score/v1alpha2/score_pb.js');

function serialize_containersai_datahub_v1alpha2_CreateNodePredictionsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.CreateNodePredictionsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.CreateNodePredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_CreateNodePredictionsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.CreateNodePredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_CreateNodesRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.CreateNodesRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.CreateNodesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_CreateNodesRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.CreateNodesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_CreatePodPredictionsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.CreatePodPredictionsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.CreatePodPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_CreatePodPredictionsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.CreatePodPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_CreatePodRecommendationsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.CreatePodRecommendationsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.CreatePodRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_CreatePodRecommendationsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.CreatePodRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_CreatePodsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.CreatePodsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.CreatePodsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_CreatePodsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.CreatePodsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_CreateSimulatedSchedulingScoresRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.CreateSimulatedSchedulingScoresRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.CreateSimulatedSchedulingScoresRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_CreateSimulatedSchedulingScoresRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.CreateSimulatedSchedulingScoresRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_DeleteNodesRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.DeleteNodesRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.DeleteNodesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_DeleteNodesRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.DeleteNodesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_DeletePodsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.DeletePodsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.DeletePodsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_DeletePodsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.DeletePodsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListNodeMetricsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListNodeMetricsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListNodeMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListNodeMetricsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListNodeMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListNodeMetricsResponse(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListNodeMetricsResponse)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListNodeMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListNodeMetricsResponse(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListNodeMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListNodePredictionsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListNodePredictionsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListNodePredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListNodePredictionsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListNodePredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListNodePredictionsResponse(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListNodePredictionsResponse)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListNodePredictionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListNodePredictionsResponse(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListNodePredictionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListNodesRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListNodesRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListNodesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListNodesRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListNodesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListNodesResponse(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListNodesResponse)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListNodesResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListNodesResponse(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListNodesResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListPodMetricsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListPodMetricsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListPodMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListPodMetricsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListPodMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListPodMetricsResponse(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListPodMetricsResponse)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListPodMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListPodMetricsResponse(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListPodMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListPodPredictionsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListPodPredictionsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListPodPredictionsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListPodPredictionsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListPodPredictionsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListPodPredictionsResponse(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListPodPredictionsResponse)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListPodPredictionsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListPodPredictionsResponse(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListPodPredictionsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListPodRecommendationsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListPodRecommendationsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListPodRecommendationsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListPodRecommendationsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListPodRecommendationsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListPodRecommendationsResponse(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListPodRecommendationsResponse)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListPodRecommendationsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListPodRecommendationsResponse(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListPodRecommendationsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListPodsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListPodsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListPodsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListPodsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListPodsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListPodsResponse(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListPodsResponse)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListPodsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListPodsResponse(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListPodsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListSimulatedSchedulingScoresRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListSimulatedSchedulingScoresRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListSimulatedSchedulingScoresRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListSimulatedSchedulingScoresRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_ListSimulatedSchedulingScoresResponse(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.ListSimulatedSchedulingScoresResponse)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.ListSimulatedSchedulingScoresResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_ListSimulatedSchedulingScoresResponse(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.ListSimulatedSchedulingScoresResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_UpdateNodesRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.UpdateNodesRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.UpdateNodesRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_UpdateNodesRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.UpdateNodesRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containersai_datahub_v1alpha2_UpdatePodsRequest(arg) {
  if (!(arg instanceof datahub_v1alpha2_datahub_pb.UpdatePodsRequest)) {
    throw new Error('Expected argument of type containersai.datahub.v1alpha2.UpdatePodsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containersai_datahub_v1alpha2_UpdatePodsRequest(buffer_arg) {
  return datahub_v1alpha2_datahub_pb.UpdatePodsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_google_rpc_Status(arg) {
  if (!(arg instanceof google_rpc_status_pb.Status)) {
    throw new Error('Expected argument of type google.rpc.Status');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_rpc_Status(buffer_arg) {
  return google_rpc_status_pb.Status.deserializeBinary(new Uint8Array(buffer_arg));
}


// *
// Service for providing data stored in the backend
var DatahubServiceService = exports.DatahubServiceService = {
  // / Used to list pod metric data
listPodMetrics: {
    path: '/containersai.datahub.v1alpha2.DatahubService/ListPodMetrics',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.ListPodMetricsRequest,
    responseType: datahub_v1alpha2_datahub_pb.ListPodMetricsResponse,
    requestSerialize: serialize_containersai_datahub_v1alpha2_ListPodMetricsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_ListPodMetricsRequest,
    responseSerialize: serialize_containersai_datahub_v1alpha2_ListPodMetricsResponse,
    responseDeserialize: deserialize_containersai_datahub_v1alpha2_ListPodMetricsResponse,
  },
  // / Used to list node metric data
listNodeMetrics: {
    path: '/containersai.datahub.v1alpha2.DatahubService/ListNodeMetrics',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.ListNodeMetricsRequest,
    responseType: datahub_v1alpha2_datahub_pb.ListNodeMetricsResponse,
    requestSerialize: serialize_containersai_datahub_v1alpha2_ListNodeMetricsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_ListNodeMetricsRequest,
    responseSerialize: serialize_containersai_datahub_v1alpha2_ListNodeMetricsResponse,
    responseDeserialize: deserialize_containersai_datahub_v1alpha2_ListNodeMetricsResponse,
  },
  // Used to list pods
listPods: {
    path: '/containersai.datahub.v1alpha2.DatahubService/ListPods',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.ListPodsRequest,
    responseType: datahub_v1alpha2_datahub_pb.ListPodsResponse,
    requestSerialize: serialize_containersai_datahub_v1alpha2_ListPodsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_ListPodsRequest,
    responseSerialize: serialize_containersai_datahub_v1alpha2_ListPodsResponse,
    responseDeserialize: deserialize_containersai_datahub_v1alpha2_ListPodsResponse,
  },
  // Used to list nodes
listNodes: {
    path: '/containersai.datahub.v1alpha2.DatahubService/ListNodes',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.ListNodesRequest,
    responseType: datahub_v1alpha2_datahub_pb.ListNodesResponse,
    requestSerialize: serialize_containersai_datahub_v1alpha2_ListNodesRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_ListNodesRequest,
    responseSerialize: serialize_containersai_datahub_v1alpha2_ListNodesResponse,
    responseDeserialize: deserialize_containersai_datahub_v1alpha2_ListNodesResponse,
  },
  // / Used to list pod predictions
listPodPredictions: {
    path: '/containersai.datahub.v1alpha2.DatahubService/ListPodPredictions',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.ListPodPredictionsRequest,
    responseType: datahub_v1alpha2_datahub_pb.ListPodPredictionsResponse,
    requestSerialize: serialize_containersai_datahub_v1alpha2_ListPodPredictionsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_ListPodPredictionsRequest,
    responseSerialize: serialize_containersai_datahub_v1alpha2_ListPodPredictionsResponse,
    responseDeserialize: deserialize_containersai_datahub_v1alpha2_ListPodPredictionsResponse,
  },
  // / Used to list node predictions
listNodePredictions: {
    path: '/containersai.datahub.v1alpha2.DatahubService/ListNodePredictions',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.ListNodePredictionsRequest,
    responseType: datahub_v1alpha2_datahub_pb.ListNodePredictionsResponse,
    requestSerialize: serialize_containersai_datahub_v1alpha2_ListNodePredictionsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_ListNodePredictionsRequest,
    responseSerialize: serialize_containersai_datahub_v1alpha2_ListNodePredictionsResponse,
    responseDeserialize: deserialize_containersai_datahub_v1alpha2_ListNodePredictionsResponse,
  },
  // / Used to list pod recommenations
listPodRecommendations: {
    path: '/containersai.datahub.v1alpha2.DatahubService/ListPodRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.ListPodRecommendationsRequest,
    responseType: datahub_v1alpha2_datahub_pb.ListPodRecommendationsResponse,
    requestSerialize: serialize_containersai_datahub_v1alpha2_ListPodRecommendationsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_ListPodRecommendationsRequest,
    responseSerialize: serialize_containersai_datahub_v1alpha2_ListPodRecommendationsResponse,
    responseDeserialize: deserialize_containersai_datahub_v1alpha2_ListPodRecommendationsResponse,
  },
  // / Used to list system scores
listSimulatedSchedulingScores: {
    path: '/containersai.datahub.v1alpha2.DatahubService/ListSimulatedSchedulingScores',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.ListSimulatedSchedulingScoresRequest,
    responseType: datahub_v1alpha2_datahub_pb.ListSimulatedSchedulingScoresResponse,
    requestSerialize: serialize_containersai_datahub_v1alpha2_ListSimulatedSchedulingScoresRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_ListSimulatedSchedulingScoresRequest,
    responseSerialize: serialize_containersai_datahub_v1alpha2_ListSimulatedSchedulingScoresResponse,
    responseDeserialize: deserialize_containersai_datahub_v1alpha2_ListSimulatedSchedulingScoresResponse,
  },
  // / Used to add pods that need to be predicted
createPods: {
    path: '/containersai.datahub.v1alpha2.DatahubService/CreatePods',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.CreatePodsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_v1alpha2_CreatePodsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_CreatePodsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // / Used to add nodes
createNodes: {
    path: '/containersai.datahub.v1alpha2.DatahubService/CreateNodes',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.CreateNodesRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_v1alpha2_CreateNodesRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_CreateNodesRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // / Used to create predictions of pods
createPodPredictions: {
    path: '/containersai.datahub.v1alpha2.DatahubService/CreatePodPredictions',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.CreatePodPredictionsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_v1alpha2_CreatePodPredictionsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_CreatePodPredictionsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // / Used to create predictions of nodes
createNodePredictions: {
    path: '/containersai.datahub.v1alpha2.DatahubService/CreateNodePredictions',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.CreateNodePredictionsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_v1alpha2_CreateNodePredictionsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_CreateNodePredictionsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // / Used to create recommendations of pods
createPodRecommendations: {
    path: '/containersai.datahub.v1alpha2.DatahubService/CreatePodRecommendations',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.CreatePodRecommendationsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_v1alpha2_CreatePodRecommendationsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_CreatePodRecommendationsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // / Used to create scores of system 
createSimulatedSchedulingScores: {
    path: '/containersai.datahub.v1alpha2.DatahubService/CreateSimulatedSchedulingScores',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.CreateSimulatedSchedulingScoresRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_v1alpha2_CreateSimulatedSchedulingScoresRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_CreateSimulatedSchedulingScoresRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // / Used to update info of pods
updatePods: {
    path: '/containersai.datahub.v1alpha2.DatahubService/UpdatePods',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.UpdatePodsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_v1alpha2_UpdatePodsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_UpdatePodsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // / Used to update info of nodes
updateNodes: {
    path: '/containersai.datahub.v1alpha2.DatahubService/UpdateNodes',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.UpdateNodesRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_v1alpha2_UpdateNodesRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_UpdateNodesRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // / Used to delete info of pods
deletePods: {
    path: '/containersai.datahub.v1alpha2.DatahubService/DeletePods',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.DeletePodsRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_v1alpha2_DeletePodsRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_DeletePodsRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
  // / Used to delete info of nodes
deleteNodes: {
    path: '/containersai.datahub.v1alpha2.DatahubService/DeleteNodes',
    requestStream: false,
    responseStream: false,
    requestType: datahub_v1alpha2_datahub_pb.DeleteNodesRequest,
    responseType: google_rpc_status_pb.Status,
    requestSerialize: serialize_containersai_datahub_v1alpha2_DeleteNodesRequest,
    requestDeserialize: deserialize_containersai_datahub_v1alpha2_DeleteNodesRequest,
    responseSerialize: serialize_google_rpc_Status,
    responseDeserialize: deserialize_google_rpc_Status,
  },
};

exports.DatahubServiceClient = grpc.makeGenericClientConstructor(DatahubServiceService);
