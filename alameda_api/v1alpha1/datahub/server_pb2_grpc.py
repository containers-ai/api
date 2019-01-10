# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from alameda_api.v1alpha1.datahub import server_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2
from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


class DatahubServiceStub(object):
  """*
  Service for providing data stored in the backend
  """

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.ListPodMetrics = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/ListPodMetrics',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodMetricsRequest.SerializeToString,
        response_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodMetricsResponse.FromString,
        )
    self.ListNodeMetrics = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/ListNodeMetrics',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListNodeMetricsRequest.SerializeToString,
        response_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListNodeMetricsResponse.FromString,
        )
    self.ListAlamedaPods = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/ListAlamedaPods',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListAlamedaPodsRequest.SerializeToString,
        response_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodsResponse.FromString,
        )
    self.ListAlamedaNodes = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/ListAlamedaNodes',
        request_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
        response_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListNodesResponse.FromString,
        )
    self.ListPodPredictions = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/ListPodPredictions',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodPredictionsRequest.SerializeToString,
        response_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodPredictionsResponse.FromString,
        )
    self.ListNodePredictions = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/ListNodePredictions',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListNodePredictionsRequest.SerializeToString,
        response_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListNodePredictionsResponse.FromString,
        )
    self.ListPodRecommendations = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/ListPodRecommendations',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodRecommendationsRequest.SerializeToString,
        response_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodRecommendationsResponse.FromString,
        )
    self.ListPodsByNodeName = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/ListPodsByNodeName',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodsByNodeNamesRequest.SerializeToString,
        response_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodsResponse.FromString,
        )
    self.ListSimulatedSchedulingScores = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/ListSimulatedSchedulingScores',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListSimulatedSchedulingScoresRequest.SerializeToString,
        response_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListSimulatedSchedulingScoresResponse.FromString,
        )
    self.CreatePods = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/CreatePods',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreatePodsRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.CreateAlamedaNodes = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/CreateAlamedaNodes',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreateAlamedaNodesRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.CreatePodPredictions = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/CreatePodPredictions',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreatePodPredictionsRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.CreateNodePredictions = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/CreateNodePredictions',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreateNodePredictionsRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.CreatePodRecommendations = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/CreatePodRecommendations',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreatePodRecommendationsRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.CreateSimulatedSchedulingScores = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/CreateSimulatedSchedulingScores',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreateSimulatedSchedulingScoresRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.DeletePods = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/DeletePods',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.DeletePodsRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )
    self.DeleteAlamedaNodes = channel.unary_unary(
        '/containers_ai.alameda.v1alpha1.datahub.DatahubService/DeleteAlamedaNodes',
        request_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.DeleteAlamedaNodesRequest.SerializeToString,
        response_deserializer=google_dot_rpc_dot_status__pb2.Status.FromString,
        )


class DatahubServiceServicer(object):
  """*
  Service for providing data stored in the backend
  """

  def ListPodMetrics(self, request, context):
    """/ Used to list pod metric data
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListNodeMetrics(self, request, context):
    """/ Used to list node metric data
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListAlamedaPods(self, request, context):
    """Used to list pods need to be predicted
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListAlamedaNodes(self, request, context):
    """Used to list nodes need to be predicted
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListPodPredictions(self, request, context):
    """/ Used to list pod predictions
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListNodePredictions(self, request, context):
    """/ Used to list node predictions
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListPodRecommendations(self, request, context):
    """/ Used to list pod recommenations
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListPodsByNodeName(self, request, context):
    """/ Used to list pods by a node name
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def ListSimulatedSchedulingScores(self, request, context):
    """/ Used to list system scores
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreatePods(self, request, context):
    """/ Used to add pods that need to be predicted
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateAlamedaNodes(self, request, context):
    """/ Used to add nodes that need to be predicted
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreatePodPredictions(self, request, context):
    """/ Used to create predictions of pods
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateNodePredictions(self, request, context):
    """/ Used to create predictions of nodes
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreatePodRecommendations(self, request, context):
    """/ Used to create recommendations of pods
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def CreateSimulatedSchedulingScores(self, request, context):
    """/ Used to create scores of system 
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeletePods(self, request, context):
    """/ Used to delete info of pods
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def DeleteAlamedaNodes(self, request, context):
    """/ Used to stop generating predictions for the nodes
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_DatahubServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'ListPodMetrics': grpc.unary_unary_rpc_method_handler(
          servicer.ListPodMetrics,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodMetricsRequest.FromString,
          response_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodMetricsResponse.SerializeToString,
      ),
      'ListNodeMetrics': grpc.unary_unary_rpc_method_handler(
          servicer.ListNodeMetrics,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListNodeMetricsRequest.FromString,
          response_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListNodeMetricsResponse.SerializeToString,
      ),
      'ListAlamedaPods': grpc.unary_unary_rpc_method_handler(
          servicer.ListAlamedaPods,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListAlamedaPodsRequest.FromString,
          response_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodsResponse.SerializeToString,
      ),
      'ListAlamedaNodes': grpc.unary_unary_rpc_method_handler(
          servicer.ListAlamedaNodes,
          request_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
          response_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListNodesResponse.SerializeToString,
      ),
      'ListPodPredictions': grpc.unary_unary_rpc_method_handler(
          servicer.ListPodPredictions,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodPredictionsRequest.FromString,
          response_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodPredictionsResponse.SerializeToString,
      ),
      'ListNodePredictions': grpc.unary_unary_rpc_method_handler(
          servicer.ListNodePredictions,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListNodePredictionsRequest.FromString,
          response_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListNodePredictionsResponse.SerializeToString,
      ),
      'ListPodRecommendations': grpc.unary_unary_rpc_method_handler(
          servicer.ListPodRecommendations,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodRecommendationsRequest.FromString,
          response_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodRecommendationsResponse.SerializeToString,
      ),
      'ListPodsByNodeName': grpc.unary_unary_rpc_method_handler(
          servicer.ListPodsByNodeName,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodsByNodeNamesRequest.FromString,
          response_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListPodsResponse.SerializeToString,
      ),
      'ListSimulatedSchedulingScores': grpc.unary_unary_rpc_method_handler(
          servicer.ListSimulatedSchedulingScores,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListSimulatedSchedulingScoresRequest.FromString,
          response_serializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.ListSimulatedSchedulingScoresResponse.SerializeToString,
      ),
      'CreatePods': grpc.unary_unary_rpc_method_handler(
          servicer.CreatePods,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreatePodsRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'CreateAlamedaNodes': grpc.unary_unary_rpc_method_handler(
          servicer.CreateAlamedaNodes,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreateAlamedaNodesRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'CreatePodPredictions': grpc.unary_unary_rpc_method_handler(
          servicer.CreatePodPredictions,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreatePodPredictionsRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'CreateNodePredictions': grpc.unary_unary_rpc_method_handler(
          servicer.CreateNodePredictions,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreateNodePredictionsRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'CreatePodRecommendations': grpc.unary_unary_rpc_method_handler(
          servicer.CreatePodRecommendations,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreatePodRecommendationsRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'CreateSimulatedSchedulingScores': grpc.unary_unary_rpc_method_handler(
          servicer.CreateSimulatedSchedulingScores,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.CreateSimulatedSchedulingScoresRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'DeletePods': grpc.unary_unary_rpc_method_handler(
          servicer.DeletePods,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.DeletePodsRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
      'DeleteAlamedaNodes': grpc.unary_unary_rpc_method_handler(
          servicer.DeleteAlamedaNodes,
          request_deserializer=alameda__api_dot_v1alpha1_dot_datahub_dot_server__pb2.DeleteAlamedaNodesRequest.FromString,
          response_serializer=google_dot_rpc_dot_status__pb2.Status.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'containers_ai.alameda.v1alpha1.datahub.DatahubService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
