// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// / This file has messages related to the CRUD to Prometheus. This file will be deprated in v0.2.
'use strict';
var grpc = require('grpc');
var alameda_api_v1alpha1_operator_server_pb = require('../../../alameda_api/v1alpha1/operator/server_pb.js');
var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js');
var google_protobuf_duration_pb = require('google-protobuf/google/protobuf/duration_pb.js');
var google_rpc_status_pb = require('../../../google/rpc/status_pb.js');

function serialize_containers_ai_alameda_v1alpha1_operator_CreatePredictResultRequest(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_operator_server_pb.CreatePredictResultRequest)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.operator.CreatePredictResultRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_operator_CreatePredictResultRequest(buffer_arg) {
  return alameda_api_v1alpha1_operator_server_pb.CreatePredictResultRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_operator_CreatePredictResultResponse(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_operator_server_pb.CreatePredictResultResponse)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.operator.CreatePredictResultResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_operator_CreatePredictResultResponse(buffer_arg) {
  return alameda_api_v1alpha1_operator_server_pb.CreatePredictResultResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_operator_GetResourceInfoRequest(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_operator_server_pb.GetResourceInfoRequest)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.operator.GetResourceInfoRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_operator_GetResourceInfoRequest(buffer_arg) {
  return alameda_api_v1alpha1_operator_server_pb.GetResourceInfoRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_operator_GetResourceInfoResponse(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_operator_server_pb.GetResourceInfoResponse)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.operator.GetResourceInfoResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_operator_GetResourceInfoResponse(buffer_arg) {
  return alameda_api_v1alpha1_operator_server_pb.GetResourceInfoResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_operator_GetResourceRecommendationRequest(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_operator_server_pb.GetResourceRecommendationRequest)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_operator_GetResourceRecommendationRequest(buffer_arg) {
  return alameda_api_v1alpha1_operator_server_pb.GetResourceRecommendationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_operator_GetResourceRecommendationResponse(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_operator_server_pb.GetResourceRecommendationResponse)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.operator.GetResourceRecommendationResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_operator_GetResourceRecommendationResponse(buffer_arg) {
  return alameda_api_v1alpha1_operator_server_pb.GetResourceRecommendationResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_operator_ListMetricsRequest(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_operator_server_pb.ListMetricsRequest)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.operator.ListMetricsRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_operator_ListMetricsRequest(buffer_arg) {
  return alameda_api_v1alpha1_operator_server_pb.ListMetricsRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_operator_ListMetricsResponse(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_operator_server_pb.ListMetricsResponse)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.operator.ListMetricsResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_operator_ListMetricsResponse(buffer_arg) {
  return alameda_api_v1alpha1_operator_server_pb.ListMetricsResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_operator_ListMetricsSumRequest(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_operator_server_pb.ListMetricsSumRequest)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.operator.ListMetricsSumRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_operator_ListMetricsSumRequest(buffer_arg) {
  return alameda_api_v1alpha1_operator_server_pb.ListMetricsSumRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_operator_ListMetricsSumResponse(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_operator_server_pb.ListMetricsSumResponse)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.operator.ListMetricsSumResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_operator_ListMetricsSumResponse(buffer_arg) {
  return alameda_api_v1alpha1_operator_server_pb.ListMetricsSumResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// *
// Service for providing data stored in Prometheus
var OperatorServiceService = exports.OperatorServiceService = {
  // / Used to list metric data
listMetrics: {
    path: '/containers_ai.alameda.v1alpha1.operator.OperatorService/ListMetrics',
    requestStream: false,
    responseStream: false,
    requestType: alameda_api_v1alpha1_operator_server_pb.ListMetricsRequest,
    responseType: alameda_api_v1alpha1_operator_server_pb.ListMetricsResponse,
    requestSerialize: serialize_containers_ai_alameda_v1alpha1_operator_ListMetricsRequest,
    requestDeserialize: deserialize_containers_ai_alameda_v1alpha1_operator_ListMetricsRequest,
    responseSerialize: serialize_containers_ai_alameda_v1alpha1_operator_ListMetricsResponse,
    responseDeserialize: deserialize_containers_ai_alameda_v1alpha1_operator_ListMetricsResponse,
  },
  // / Used to list summation of metric data
listMetricsSum: {
    path: '/containers_ai.alameda.v1alpha1.operator.OperatorService/ListMetricsSum',
    requestStream: false,
    responseStream: false,
    requestType: alameda_api_v1alpha1_operator_server_pb.ListMetricsSumRequest,
    responseType: alameda_api_v1alpha1_operator_server_pb.ListMetricsSumResponse,
    requestSerialize: serialize_containers_ai_alameda_v1alpha1_operator_ListMetricsSumRequest,
    requestDeserialize: deserialize_containers_ai_alameda_v1alpha1_operator_ListMetricsSumRequest,
    responseSerialize: serialize_containers_ai_alameda_v1alpha1_operator_ListMetricsSumResponse,
    responseDeserialize: deserialize_containers_ai_alameda_v1alpha1_operator_ListMetricsSumResponse,
  },
  // / Used to requst for creating prediction results
createPredictResult: {
    path: '/containers_ai.alameda.v1alpha1.operator.OperatorService/CreatePredictResult',
    requestStream: false,
    responseStream: false,
    requestType: alameda_api_v1alpha1_operator_server_pb.CreatePredictResultRequest,
    responseType: alameda_api_v1alpha1_operator_server_pb.CreatePredictResultResponse,
    requestSerialize: serialize_containers_ai_alameda_v1alpha1_operator_CreatePredictResultRequest,
    requestDeserialize: deserialize_containers_ai_alameda_v1alpha1_operator_CreatePredictResultRequest,
    responseSerialize: serialize_containers_ai_alameda_v1alpha1_operator_CreatePredictResultResponse,
    responseDeserialize: deserialize_containers_ai_alameda_v1alpha1_operator_CreatePredictResultResponse,
  },
  // / Used to get resource data
getResourceInfo: {
    path: '/containers_ai.alameda.v1alpha1.operator.OperatorService/GetResourceInfo',
    requestStream: false,
    responseStream: false,
    requestType: alameda_api_v1alpha1_operator_server_pb.GetResourceInfoRequest,
    responseType: alameda_api_v1alpha1_operator_server_pb.GetResourceInfoResponse,
    requestSerialize: serialize_containers_ai_alameda_v1alpha1_operator_GetResourceInfoRequest,
    requestDeserialize: deserialize_containers_ai_alameda_v1alpha1_operator_GetResourceInfoRequest,
    responseSerialize: serialize_containers_ai_alameda_v1alpha1_operator_GetResourceInfoResponse,
    responseDeserialize: deserialize_containers_ai_alameda_v1alpha1_operator_GetResourceInfoResponse,
  },
  // / Used to get resource recommendation
getResourceRecommendation: {
    path: '/containers_ai.alameda.v1alpha1.operator.OperatorService/GetResourceRecommendation',
    requestStream: false,
    responseStream: false,
    requestType: alameda_api_v1alpha1_operator_server_pb.GetResourceRecommendationRequest,
    responseType: alameda_api_v1alpha1_operator_server_pb.GetResourceRecommendationResponse,
    requestSerialize: serialize_containers_ai_alameda_v1alpha1_operator_GetResourceRecommendationRequest,
    requestDeserialize: deserialize_containers_ai_alameda_v1alpha1_operator_GetResourceRecommendationRequest,
    responseSerialize: serialize_containers_ai_alameda_v1alpha1_operator_GetResourceRecommendationResponse,
    responseDeserialize: deserialize_containers_ai_alameda_v1alpha1_operator_GetResourceRecommendationResponse,
  },
};

exports.OperatorServiceClient = grpc.makeGenericClientConstructor(OperatorServiceService);
