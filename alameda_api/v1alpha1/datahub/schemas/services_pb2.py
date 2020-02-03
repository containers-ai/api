# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/schemas/services.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub.schemas import schemas_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_schemas__pb2
from alameda_api.v1alpha1.datahub.schemas import types_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2
from google.rpc import status_pb2 as google_dot_rpc_dot_status__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/schemas/services.proto',
  package='containersai.alameda.v1alpha1.datahub.schemas',
  syntax='proto3',
  serialized_options=_b('ZAgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemas'),
  serialized_pb=_b('\n3alameda_api/v1alpha1/datahub/schemas/services.proto\x12-containersai.alameda.v1alpha1.datahub.schemas\x1a\x32\x61lameda_api/v1alpha1/datahub/schemas/schemas.proto\x1a\x30\x61lameda_api/v1alpha1/datahub/schemas/types.proto\x1a\x17google/rpc/status.proto\"^\n\x14\x43reateSchemasRequest\x12\x46\n\x07schemas\x18\x01 \x03(\x0b\x32\x35.containersai.alameda.v1alpha1.datahub.schemas.Schema\"e\n\x12ListSchemasRequest\x12O\n\x0cschema_metas\x18\x01 \x01(\x0b\x32\x39.containersai.alameda.v1alpha1.datahub.schemas.SchemaMeta\"\x81\x01\n\x13ListSchemasResponse\x12\"\n\x06status\x18\x01 \x01(\x0b\x32\x12.google.rpc.Status\x12\x46\n\x07schemas\x18\x02 \x03(\x0b\x32\x35.containersai.alameda.v1alpha1.datahub.schemas.Schema\"g\n\x14\x44\x65leteSchemasRequest\x12O\n\x0cschema_metas\x18\x01 \x01(\x0b\x32\x39.containersai.alameda.v1alpha1.datahub.schemas.SchemaMetaBCZAgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/schemasb\x06proto3')
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_schemas__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2.DESCRIPTOR,google_dot_rpc_dot_status__pb2.DESCRIPTOR,])




_CREATESCHEMASREQUEST = _descriptor.Descriptor(
  name='CreateSchemasRequest',
  full_name='containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='schemas', full_name='containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest.schemas', index=0,
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
  serialized_start=229,
  serialized_end=323,
)


_LISTSCHEMASREQUEST = _descriptor.Descriptor(
  name='ListSchemasRequest',
  full_name='containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_metas', full_name='containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest.schema_metas', index=0,
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
  serialized_start=325,
  serialized_end=426,
)


_LISTSCHEMASRESPONSE = _descriptor.Descriptor(
  name='ListSchemasResponse',
  full_name='containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='status', full_name='containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse.status', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='schemas', full_name='containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse.schemas', index=1,
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
  serialized_start=429,
  serialized_end=558,
)


_DELETESCHEMASREQUEST = _descriptor.Descriptor(
  name='DeleteSchemasRequest',
  full_name='containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='schema_metas', full_name='containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest.schema_metas', index=0,
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
  serialized_start=560,
  serialized_end=663,
)

_CREATESCHEMASREQUEST.fields_by_name['schemas'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_schemas__pb2._SCHEMA
_LISTSCHEMASREQUEST.fields_by_name['schema_metas'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
_LISTSCHEMASRESPONSE.fields_by_name['status'].message_type = google_dot_rpc_dot_status__pb2._STATUS
_LISTSCHEMASRESPONSE.fields_by_name['schemas'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_schemas__pb2._SCHEMA
_DELETESCHEMASREQUEST.fields_by_name['schema_metas'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_schemas_dot_types__pb2._SCHEMAMETA
DESCRIPTOR.message_types_by_name['CreateSchemasRequest'] = _CREATESCHEMASREQUEST
DESCRIPTOR.message_types_by_name['ListSchemasRequest'] = _LISTSCHEMASREQUEST
DESCRIPTOR.message_types_by_name['ListSchemasResponse'] = _LISTSCHEMASRESPONSE
DESCRIPTOR.message_types_by_name['DeleteSchemasRequest'] = _DELETESCHEMASREQUEST
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

CreateSchemasRequest = _reflection.GeneratedProtocolMessageType('CreateSchemasRequest', (_message.Message,), {
  'DESCRIPTOR' : _CREATESCHEMASREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.schemas.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.schemas.CreateSchemasRequest)
  })
_sym_db.RegisterMessage(CreateSchemasRequest)

ListSchemasRequest = _reflection.GeneratedProtocolMessageType('ListSchemasRequest', (_message.Message,), {
  'DESCRIPTOR' : _LISTSCHEMASREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.schemas.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.schemas.ListSchemasRequest)
  })
_sym_db.RegisterMessage(ListSchemasRequest)

ListSchemasResponse = _reflection.GeneratedProtocolMessageType('ListSchemasResponse', (_message.Message,), {
  'DESCRIPTOR' : _LISTSCHEMASRESPONSE,
  '__module__' : 'alameda_api.v1alpha1.datahub.schemas.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.schemas.ListSchemasResponse)
  })
_sym_db.RegisterMessage(ListSchemasResponse)

DeleteSchemasRequest = _reflection.GeneratedProtocolMessageType('DeleteSchemasRequest', (_message.Message,), {
  'DESCRIPTOR' : _DELETESCHEMASREQUEST,
  '__module__' : 'alameda_api.v1alpha1.datahub.schemas.services_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.schemas.DeleteSchemasRequest)
  })
_sym_db.RegisterMessage(DeleteSchemasRequest)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
