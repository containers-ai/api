# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: ai-service/alameda.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
from google.protobuf import descriptor_pb2
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import empty_pb2 as google_dot_protobuf_dot_empty__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='ai-service/alameda.proto',
  package='',
  syntax='proto3',
  serialized_pb=_b('\n\x18\x61i-service/alameda.proto\x1a\x1bgoogle/protobuf/empty.proto\"c\n\x06Object\x12\x1a\n\x04type\x18\x01 \x01(\x0e\x32\x0c.Object.Type\x12\x0b\n\x03uid\x18\x02 \x01(\t\x12\x11\n\tnamespace\x18\x03 \x01(\t\x12\x0c\n\x04name\x18\x04 \x01(\t\"\x0f\n\x04Type\x12\x07\n\x03POD\x10\x00\"?\n#PredictionObjectListCreationRequest\x12\x18\n\x07objects\x18\x01 \x03(\x0b\x32\x07.Object\"?\n#PredictionObjectListDeletionRequest\x12\x18\n\x07objects\x18\x01 \x03(\x0b\x32\x07.Object\"\"\n\x0fRequestResponse\x12\x0f\n\x07message\x18\x01 \x01(\t2\xc3\x01\n\x11\x41lamendaAIService\x12S\n\x17\x43reatePredictionObjects\x12$.PredictionObjectListCreationRequest\x1a\x10.RequestResponse\"\x00\x12Y\n\x17\x44\x65letePredictionObjects\x12$.PredictionObjectListDeletionRequest\x1a\x16.google.protobuf.Empty\"\x00\x62\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_empty__pb2.DESCRIPTOR,])



_OBJECT_TYPE = _descriptor.EnumDescriptor(
  name='Type',
  full_name='Object.Type',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='POD', index=0, number=0,
      options=None,
      type=None),
  ],
  containing_type=None,
  options=None,
  serialized_start=141,
  serialized_end=156,
)
_sym_db.RegisterEnumDescriptor(_OBJECT_TYPE)


_OBJECT = _descriptor.Descriptor(
  name='Object',
  full_name='Object',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='type', full_name='Object.type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='uid', full_name='Object.uid', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='namespace', full_name='Object.namespace', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
    _descriptor.FieldDescriptor(
      name='name', full_name='Object.name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _OBJECT_TYPE,
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=57,
  serialized_end=156,
)


_PREDICTIONOBJECTLISTCREATIONREQUEST = _descriptor.Descriptor(
  name='PredictionObjectListCreationRequest',
  full_name='PredictionObjectListCreationRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='objects', full_name='PredictionObjectListCreationRequest.objects', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=158,
  serialized_end=221,
)


_PREDICTIONOBJECTLISTDELETIONREQUEST = _descriptor.Descriptor(
  name='PredictionObjectListDeletionRequest',
  full_name='PredictionObjectListDeletionRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='objects', full_name='PredictionObjectListDeletionRequest.objects', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=223,
  serialized_end=286,
)


_REQUESTRESPONSE = _descriptor.Descriptor(
  name='RequestResponse',
  full_name='RequestResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='message', full_name='RequestResponse.message', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      options=None),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=288,
  serialized_end=322,
)

_OBJECT.fields_by_name['type'].enum_type = _OBJECT_TYPE
_OBJECT_TYPE.containing_type = _OBJECT
_PREDICTIONOBJECTLISTCREATIONREQUEST.fields_by_name['objects'].message_type = _OBJECT
_PREDICTIONOBJECTLISTDELETIONREQUEST.fields_by_name['objects'].message_type = _OBJECT
DESCRIPTOR.message_types_by_name['Object'] = _OBJECT
DESCRIPTOR.message_types_by_name['PredictionObjectListCreationRequest'] = _PREDICTIONOBJECTLISTCREATIONREQUEST
DESCRIPTOR.message_types_by_name['PredictionObjectListDeletionRequest'] = _PREDICTIONOBJECTLISTDELETIONREQUEST
DESCRIPTOR.message_types_by_name['RequestResponse'] = _REQUESTRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Object = _reflection.GeneratedProtocolMessageType('Object', (_message.Message,), dict(
  DESCRIPTOR = _OBJECT,
  __module__ = 'ai_service.alameda_pb2'
  # @@protoc_insertion_point(class_scope:Object)
  ))
_sym_db.RegisterMessage(Object)

PredictionObjectListCreationRequest = _reflection.GeneratedProtocolMessageType('PredictionObjectListCreationRequest', (_message.Message,), dict(
  DESCRIPTOR = _PREDICTIONOBJECTLISTCREATIONREQUEST,
  __module__ = 'ai_service.alameda_pb2'
  # @@protoc_insertion_point(class_scope:PredictionObjectListCreationRequest)
  ))
_sym_db.RegisterMessage(PredictionObjectListCreationRequest)

PredictionObjectListDeletionRequest = _reflection.GeneratedProtocolMessageType('PredictionObjectListDeletionRequest', (_message.Message,), dict(
  DESCRIPTOR = _PREDICTIONOBJECTLISTDELETIONREQUEST,
  __module__ = 'ai_service.alameda_pb2'
  # @@protoc_insertion_point(class_scope:PredictionObjectListDeletionRequest)
  ))
_sym_db.RegisterMessage(PredictionObjectListDeletionRequest)

RequestResponse = _reflection.GeneratedProtocolMessageType('RequestResponse', (_message.Message,), dict(
  DESCRIPTOR = _REQUESTRESPONSE,
  __module__ = 'ai_service.alameda_pb2'
  # @@protoc_insertion_point(class_scope:RequestResponse)
  ))
_sym_db.RegisterMessage(RequestResponse)



_ALAMENDAAISERVICE = _descriptor.ServiceDescriptor(
  name='AlamendaAIService',
  full_name='AlamendaAIService',
  file=DESCRIPTOR,
  index=0,
  options=None,
  serialized_start=325,
  serialized_end=520,
  methods=[
  _descriptor.MethodDescriptor(
    name='CreatePredictionObjects',
    full_name='AlamendaAIService.CreatePredictionObjects',
    index=0,
    containing_service=None,
    input_type=_PREDICTIONOBJECTLISTCREATIONREQUEST,
    output_type=_REQUESTRESPONSE,
    options=None,
  ),
  _descriptor.MethodDescriptor(
    name='DeletePredictionObjects',
    full_name='AlamendaAIService.DeletePredictionObjects',
    index=1,
    containing_service=None,
    input_type=_PREDICTIONOBJECTLISTDELETIONREQUEST,
    output_type=google_dot_protobuf_dot_empty__pb2._EMPTY,
    options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_ALAMENDAAISERVICE)

DESCRIPTOR.services_by_name['AlamendaAIService'] = _ALAMENDAAISERVICE

try:
  # THESE ELEMENTS WILL BE DEPRECATED.
  # Please use the generated *_pb2_grpc.py files instead.
  import grpc
  from grpc.beta import implementations as beta_implementations
  from grpc.beta import interfaces as beta_interfaces
  from grpc.framework.common import cardinality
  from grpc.framework.interfaces.face import utilities as face_utilities


  class AlamendaAIServiceStub(object):
    # missing associated documentation comment in .proto file
    pass

    def __init__(self, channel):
      """Constructor.

      Args:
        channel: A grpc.Channel.
      """
      self.CreatePredictionObjects = channel.unary_unary(
          '/AlamendaAIService/CreatePredictionObjects',
          request_serializer=PredictionObjectListCreationRequest.SerializeToString,
          response_deserializer=RequestResponse.FromString,
          )
      self.DeletePredictionObjects = channel.unary_unary(
          '/AlamendaAIService/DeletePredictionObjects',
          request_serializer=PredictionObjectListDeletionRequest.SerializeToString,
          response_deserializer=google_dot_protobuf_dot_empty__pb2.Empty.FromString,
          )


  class AlamendaAIServiceServicer(object):
    # missing associated documentation comment in .proto file
    pass

    def CreatePredictionObjects(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')

    def DeletePredictionObjects(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.set_code(grpc.StatusCode.UNIMPLEMENTED)
      context.set_details('Method not implemented!')
      raise NotImplementedError('Method not implemented!')


  def add_AlamendaAIServiceServicer_to_server(servicer, server):
    rpc_method_handlers = {
        'CreatePredictionObjects': grpc.unary_unary_rpc_method_handler(
            servicer.CreatePredictionObjects,
            request_deserializer=PredictionObjectListCreationRequest.FromString,
            response_serializer=RequestResponse.SerializeToString,
        ),
        'DeletePredictionObjects': grpc.unary_unary_rpc_method_handler(
            servicer.DeletePredictionObjects,
            request_deserializer=PredictionObjectListDeletionRequest.FromString,
            response_serializer=google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
        ),
    }
    generic_handler = grpc.method_handlers_generic_handler(
        'AlamendaAIService', rpc_method_handlers)
    server.add_generic_rpc_handlers((generic_handler,))


  class BetaAlamendaAIServiceServicer(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def CreatePredictionObjects(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)
    def DeletePredictionObjects(self, request, context):
      # missing associated documentation comment in .proto file
      pass
      context.code(beta_interfaces.StatusCode.UNIMPLEMENTED)


  class BetaAlamendaAIServiceStub(object):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This class was generated
    only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0."""
    # missing associated documentation comment in .proto file
    pass
    def CreatePredictionObjects(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    CreatePredictionObjects.future = None
    def DeletePredictionObjects(self, request, timeout, metadata=None, with_call=False, protocol_options=None):
      # missing associated documentation comment in .proto file
      pass
      raise NotImplementedError()
    DeletePredictionObjects.future = None


  def beta_create_AlamendaAIService_server(servicer, pool=None, pool_size=None, default_timeout=None, maximum_timeout=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_deserializers = {
      ('AlamendaAIService', 'CreatePredictionObjects'): PredictionObjectListCreationRequest.FromString,
      ('AlamendaAIService', 'DeletePredictionObjects'): PredictionObjectListDeletionRequest.FromString,
    }
    response_serializers = {
      ('AlamendaAIService', 'CreatePredictionObjects'): RequestResponse.SerializeToString,
      ('AlamendaAIService', 'DeletePredictionObjects'): google_dot_protobuf_dot_empty__pb2.Empty.SerializeToString,
    }
    method_implementations = {
      ('AlamendaAIService', 'CreatePredictionObjects'): face_utilities.unary_unary_inline(servicer.CreatePredictionObjects),
      ('AlamendaAIService', 'DeletePredictionObjects'): face_utilities.unary_unary_inline(servicer.DeletePredictionObjects),
    }
    server_options = beta_implementations.server_options(request_deserializers=request_deserializers, response_serializers=response_serializers, thread_pool=pool, thread_pool_size=pool_size, default_timeout=default_timeout, maximum_timeout=maximum_timeout)
    return beta_implementations.server(method_implementations, options=server_options)


  def beta_create_AlamendaAIService_stub(channel, host=None, metadata_transformer=None, pool=None, pool_size=None):
    """The Beta API is deprecated for 0.15.0 and later.

    It is recommended to use the GA API (classes and functions in this
    file not marked beta) for all further purposes. This function was
    generated only to ease transition from grpcio<0.15.0 to grpcio>=0.15.0"""
    request_serializers = {
      ('AlamendaAIService', 'CreatePredictionObjects'): PredictionObjectListCreationRequest.SerializeToString,
      ('AlamendaAIService', 'DeletePredictionObjects'): PredictionObjectListDeletionRequest.SerializeToString,
    }
    response_deserializers = {
      ('AlamendaAIService', 'CreatePredictionObjects'): RequestResponse.FromString,
      ('AlamendaAIService', 'DeletePredictionObjects'): google_dot_protobuf_dot_empty__pb2.Empty.FromString,
    }
    cardinalities = {
      'CreatePredictionObjects': cardinality.Cardinality.UNARY_UNARY,
      'DeletePredictionObjects': cardinality.Cardinality.UNARY_UNARY,
    }
    stub_options = beta_implementations.stub_options(host=host, metadata_transformer=metadata_transformer, request_serializers=request_serializers, response_deserializers=response_deserializers, thread_pool=pool, thread_pool_size=pool_size)
    return beta_implementations.dynamic_stub(channel, 'AlamendaAIService', cardinalities, options=stub_options)
except ImportError:
  pass
# @@protoc_insertion_point(module_scope)
