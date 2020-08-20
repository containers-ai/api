// GENERATED CODE -- DO NOT EDIT!

// Original file comments:
// / This file has messages and services (methods) related to the AI engine. This file will be deprated in v0.2.
'use strict';
var grpc = require('grpc');
var alameda_api_v1alpha1_ai_service_ai_service_pb = require('../../../alameda_api/v1alpha1/ai_service/ai_service_pb.js');
var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js');

function serialize_containers_ai_alameda_v1alpha1_ai_service_PredictionObjectListCreationRequest(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_ai_service_ai_service_pb.PredictionObjectListCreationRequest)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListCreationRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_ai_service_PredictionObjectListCreationRequest(buffer_arg) {
  return alameda_api_v1alpha1_ai_service_ai_service_pb.PredictionObjectListCreationRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_ai_service_PredictionObjectListDeletionRequest(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_ai_service_ai_service_pb.PredictionObjectListDeletionRequest)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.ai_service.PredictionObjectListDeletionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_ai_service_PredictionObjectListDeletionRequest(buffer_arg) {
  return alameda_api_v1alpha1_ai_service_ai_service_pb.PredictionObjectListDeletionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_containers_ai_alameda_v1alpha1_ai_service_RequestResponse(arg) {
  if (!(arg instanceof alameda_api_v1alpha1_ai_service_ai_service_pb.RequestResponse)) {
    throw new Error('Expected argument of type containers_ai.alameda.v1alpha1.ai_service.RequestResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_containers_ai_alameda_v1alpha1_ai_service_RequestResponse(buffer_arg) {
  return alameda_api_v1alpha1_ai_service_ai_service_pb.RequestResponse.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_google_protobuf_Empty(arg) {
  if (!(arg instanceof google_protobuf_empty_pb.Empty)) {
    throw new Error('Expected argument of type google.protobuf.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_google_protobuf_Empty(buffer_arg) {
  return google_protobuf_empty_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}


// *
// Service for creating prediction objects.
var AlamedaAIServiceService = exports.AlamedaAIServiceService = {
  // / Used to create prediction objects
createPredictionObjects: {
    path: '/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/CreatePredictionObjects',
    requestStream: false,
    responseStream: false,
    requestType: alameda_api_v1alpha1_ai_service_ai_service_pb.PredictionObjectListCreationRequest,
    responseType: alameda_api_v1alpha1_ai_service_ai_service_pb.RequestResponse,
    requestSerialize: serialize_containers_ai_alameda_v1alpha1_ai_service_PredictionObjectListCreationRequest,
    requestDeserialize: deserialize_containers_ai_alameda_v1alpha1_ai_service_PredictionObjectListCreationRequest,
    responseSerialize: serialize_containers_ai_alameda_v1alpha1_ai_service_RequestResponse,
    responseDeserialize: deserialize_containers_ai_alameda_v1alpha1_ai_service_RequestResponse,
  },
  // / Used to remove prediction objects
deletePredictionObjects: {
    path: '/containers_ai.alameda.v1alpha1.ai_service.AlamedaAIService/DeletePredictionObjects',
    requestStream: false,
    responseStream: false,
    requestType: alameda_api_v1alpha1_ai_service_ai_service_pb.PredictionObjectListDeletionRequest,
    responseType: google_protobuf_empty_pb.Empty,
    requestSerialize: serialize_containers_ai_alameda_v1alpha1_ai_service_PredictionObjectListDeletionRequest,
    requestDeserialize: deserialize_containers_ai_alameda_v1alpha1_ai_service_PredictionObjectListDeletionRequest,
    responseSerialize: serialize_google_protobuf_Empty,
    responseDeserialize: deserialize_google_protobuf_Empty,
  },
};

exports.AlamedaAIServiceClient = grpc.makeGenericClientConstructor(AlamedaAIServiceService);
