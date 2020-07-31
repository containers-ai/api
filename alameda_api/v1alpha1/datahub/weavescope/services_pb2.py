# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/weavescope/services.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/weavescope/services.proto',
  package='containersai.alameda.v1alpha1.datahub.weavescope',
  syntax='proto3',
  serialized_options=b'ZDgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/weavescope',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n6alameda_api/v1alpha1/datahub/weavescope/services.proto\x12\x30\x63ontainersai.alameda.v1alpha1.datahub.weavescope\x1a\x17google/rpc/status.proto\"-\n\x1aListWeaveScopeHostsRequest\x12\x0f\n\x07host_id\x18\x01 \x01(\t\"+\n\x19ListWeaveScopePodsRequest\x12\x0e\n\x06pod_id\x18\x01 \x01(\t\"7\n\x1fListWeaveScopeContainersRequest\x12\x14\n\x0c\x63ontainer_id\x18\x01 \x01(\t\"I\n\x12WeaveScopeResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12\x0f\n\x07rawdata\x18\x02 \x01(\tBFZDgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/weavescopeb\x06proto3'
  ,
  dependencies=[google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_LISTWEAVESCOPEHOSTSREQUEST = _descriptor.Descriptor(
  name='ListWeaveScopeHostsRequest',
  full_name='containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='host_id', full_name='containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest.host_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=133,
  serialized_end=178,
)


_LISTWEAVESCOPEPODSREQUEST = _descriptor.Descriptor(
  name='ListWeaveScopePodsRequest',
  full_name='containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='pod_id', full_name='containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest.pod_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=180,
  serialized_end=223,
)


_LISTWEAVESCOPECONTAINERSREQUEST = _descriptor.Descriptor(
  name='ListWeaveScopeContainersRequest',
  full_name='containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='container_id', full_name='containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest.container_id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=225,
  serialized_end=280,
)


_WEAVESCOPERESPONSE = _descriptor.Descriptor(
  name='WeaveScopeResponse',
  full_name='containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='rawdata', full_name='containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse.rawdata', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
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
  serialized_start=282,
  serialized_end=355,
)

_WEAVESCOPERESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
DESCRIPTOR.message_types_by_name['ListWeaveScopeHostsRequest'] = _LISTWEAVESCOPEHOSTSREQUEST
DESCRIPTOR.message_types_by_name['ListWeaveScopePodsRequest'] = _LISTWEAVESCOPEPODSREQUEST
DESCRIPTOR.message_types_by_name['ListWeaveScopeContainersRequest'] = _LISTWEAVESCOPECONTAINERSREQUEST
DESCRIPTOR.message_types_by_name['WeaveScopeResponse'] = _WEAVESCOPERESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ListWeaveScopeHostsRequest = _reflection.GeneratedProtocolMessageType('ListWeaveScopeHostsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTWEAVESCOPEHOSTSREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.weavescope.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeHostsRequest)
  })
_sym_db.RegisterMessage(ListWeaveScopeHostsRequest)

ListWeaveScopePodsRequest = _reflection.GeneratedProtocolMessageType('ListWeaveScopePodsRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTWEAVESCOPEPODSREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.weavescope.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopePodsRequest)
  })
_sym_db.RegisterMessage(ListWeaveScopePodsRequest)

ListWeaveScopeContainersRequest = _reflection.GeneratedProtocolMessageType('ListWeaveScopeContainersRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTWEAVESCOPECONTAINERSREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.weavescope.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.weavescope.ListWeaveScopeContainersRequest)
  })
_sym_db.RegisterMessage(ListWeaveScopeContainersRequest)

WeaveScopeResponse = _reflection.GeneratedProtocolMessageType('WeaveScopeResponse', (_message.Message,), {
  'DESCRIPTOR' : _WEAVESCOPERESPONSE,
  '__module__' : 'alameda_api.v1alpha1.datahub.weavescope.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.weavescope.WeaveScopeResponse)
  })
_sym_db.RegisterMessage(WeaveScopeResponse)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
