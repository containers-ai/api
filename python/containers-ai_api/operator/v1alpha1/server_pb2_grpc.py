# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from operator.v1alpha1 import server_pb2 as operator_dot_v1alpha1_dot_server__pb2


class OperatorServiceStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetMetric = channel.unary_unary(
        '/OperatorService/GetMetric',
        request_serializer=operator_dot_v1alpha1_dot_server__pb2.GetMetricRequest.SerializeToString,
        response_deserializer=operator_dot_v1alpha1_dot_server__pb2.GetMetricResponse.FromString,
        )
    self.PostPredictResult = channel.unary_unary(
        '/OperatorService/PostPredictResult',
        request_serializer=operator_dot_v1alpha1_dot_server__pb2.PostPredictResultRequest.SerializeToString,
        response_deserializer=operator_dot_v1alpha1_dot_server__pb2.PostPredictResultResponse.FromString,
        )


class OperatorServiceServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetMetric(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')

  def PostPredictResult(self, request, context):
    # missing associated documentation comment in .proto file
    pass
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_OperatorServiceServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetMetric': grpc.unary_unary_rpc_method_handler(
          servicer.GetMetric,
          request_deserializer=operator_dot_v1alpha1_dot_server__pb2.GetMetricRequest.FromString,
          response_serializer=operator_dot_v1alpha1_dot_server__pb2.GetMetricResponse.SerializeToString,
      ),
      'PostPredictResult': grpc.unary_unary_rpc_method_handler(
          servicer.PostPredictResult,
          request_deserializer=operator_dot_v1alpha1_dot_server__pb2.PostPredictResultRequest.FromString,
          response_serializer=operator_dot_v1alpha1_dot_server__pb2.PostPredictResultResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'OperatorService', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))
