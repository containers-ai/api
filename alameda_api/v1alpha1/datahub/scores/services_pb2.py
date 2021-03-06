# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/scores/services.proto

from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.common import queries_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_queries__pb2
from alameda_api.v1alpha1.datahub.scores import scores_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_scores_dot_scores__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/scores/services.proto',
  package='containersai.alameda.v1alpha1.datahub.scores',
  syntax='proto3',
  serialized_options=b'Z@github.com/containers-ai/api/alameda_api/v1alpha1/datahub/scores',
  serialized_pb=b'\n2alameda_api/v1alpha1/datahub/scores/services.proto\x12,containersai.alameda.v1alpha1.datahub.scores\x1a\x31\x61lameda_api/v1alpha1/datahub/common/queries.proto\x1a\x30\x61lameda_api/v1alpha1/datahub/scores/scores.proto\x1a\x17google/rpc/status.proto\"}\n$ListSimulatedSchedulingScoresRequest\x12U\n\x0fquery_condition\x18\x01 \x01(\x0b\x32<.containersai.alameda.v1alpha1.datahub.common.QueryCondition\"\xa3\x01\n%ListSimulatedSchedulingScoresResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12V\n\x06scores\x18\x02 \x03(\x0b\x32\x46.containersai.alameda.v1alpha1.datahub.scores.SimulatedSchedulingScore\"\x80\x01\n&CreateSimulatedSchedulingScoresRequest\x12V\n\x06scores\x18\x01 \x03(\x0b\x32\x46.containersai.alameda.v1alpha1.datahub.scores.SimulatedSchedulingScoreBBZ@github.com/containers-ai/api/alameda_api/v1alpha1/datahub/scoresb\x06proto3'
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_queries__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_scores_dot_scores__pb2.DESCRIPTOR,google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_LISTSIMULATEDSCHEDULINGSCORESREQUEST = _descriptor.Descriptor(
  name='ListSimulatedSchedulingScoresRequest',
  full_name='containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='query_condition', full_name='containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest.query_condition', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=226,
  serialized_end=351,
)


_LISTSIMULATEDSCHEDULINGSCORESRESPONSE = _descriptor.Descriptor(
  name='ListSimulatedSchedulingScoresResponse',
  full_name='containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='scores', full_name='containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse.scores', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=354,
  serialized_end=517,
)


_CREATESIMULATEDSCHEDULINGSCORESREQUEST = _descriptor.Descriptor(
  name='CreateSimulatedSchedulingScoresRequest',
  full_name='containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='scores', full_name='containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest.scores', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=520,
  serialized_end=648,
)

_LISTSIMULATEDSCHEDULINGSCORESREQUEST.fields_by_name['query_condition'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_common_dot_queries__pb2._QUERYCONDITION
_LISTSIMULATEDSCHEDULINGSCORESRESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_LISTSIMULATEDSCHEDULINGSCORESRESPONSE.fields_by_name['scores'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_scores_dot_scores__pb2._SIMULATEDSCHEDULINGSCORE
_CREATESIMULATEDSCHEDULINGSCORESREQUEST.fields_by_name['scores'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_scores_dot_scores__pb2._SIMULATEDSCHEDULINGSCORE
DESCRIPTOR.message_types_by_name['ListSimulatedSchedulingScoresRequest'] = _LISTSIMULATEDSCHEDULINGSCORESREQUEST
DESCRIPTOR.message_types_by_name['ListSimulatedSchedulingScoresResponse'] = _LISTSIMULATEDSCHEDULINGSCORESRESPONSE
DESCRIPTOR.message_types_by_name['CreateSimulatedSchedulingScoresRequest'] = _CREATESIMULATEDSCHEDULINGSCORESREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListSimulatedSchedulingScoresRequest = _reflection.GeneratedProtocolMessageType('ListSimulatedSchedulingScoresRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTSIMULATEDSCHEDULINGSCORESREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.scores.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresRequest)
  })
_sym_db.RegisterMessage(ListSimulatedSchedulingScoresRequest)

ListSimulatedSchedulingScoresResponse = _reflection.GeneratedProtocolMessageType('ListSimulatedSchedulingScoresResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTSIMULATEDSCHEDULINGSCORESRESPONSE,
  '__module__' : 'alameda_api.v1alpha1.datahub.scores.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.scores.ListSimulatedSchedulingScoresResponse)
  })
_sym_db.RegisterMessage(ListSimulatedSchedulingScoresResponse)

CreateSimulatedSchedulingScoresRequest = _reflection.GeneratedProtocolMessageType('CreateSimulatedSchedulingScoresRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATESIMULATEDSCHEDULINGSCORESREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.scores.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.scores.CreateSimulatedSchedulingScoresRequest)
  })
_sym_db.RegisterMessage(CreateSimulatedSchedulingScoresRequest)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
