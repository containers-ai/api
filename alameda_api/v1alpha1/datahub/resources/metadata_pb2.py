# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/resources/metadata.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/resources/metadata.proto',
  package='containersai.alameda.v1alpha1.datahub.resources',
  syntax='proto3',
  serialized_options=b'ZCgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/resources',
  create_key=_descriptor._internal_create_key,
  serialized_pb=b'\n5alameda_api/v1alpha1/datahub/resources/metadata.proto\x12/containersai.alameda.v1alpha1.datahub.resources\"c\n\nObjectMeta\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x11\n\tnamespace\x18\x02 \x01(\t\x12\x11\n\tnode_name\x18\x03 \x01(\t\x12\x14\n\x0c\x63luster_name\x18\x04 \x01(\t\x12\x0b\n\x03uid\x18\x05 \x01(\t\"\xa7\x01\n\x0eOwnerReference\x12P\n\x0bobject_meta\x18\x01 \x01(\x0b\x32;.containersai.alameda.v1alpha1.datahub.resources.ObjectMeta\x12\x43\n\x04kind\x18\x02 \x01(\x0e\x32\x35.containersai.alameda.v1alpha1.datahub.resources.Kind*d\n\x04Kind\x12\x12\n\x0eKIND_UNDEFINED\x10\x00\x12\x0e\n\nDEPLOYMENT\x10\x01\x12\x14\n\x10\x44\x45PLOYMENTCONFIG\x10\x02\x12\x0f\n\x0bSTATEFULSET\x10\x03\x12\x11\n\rALAMEDASCALER\x10\x04*M\n\x0bScalingTool\x12\x1a\n\x16SCALING_TOOL_UNDEFINED\x10\x00\x12\x08\n\x04NONE\x10\x01\x12\x07\n\x03VPA\x10\x02\x12\x07\n\x03HPA\x10\x03\x12\x06\n\x02\x43\x41\x10\x04\x42\x45ZCgithub.com/containers-ai/api/alameda_api/v1alpha1/datahub/resourcesb\x06proto3'
)

_KIND = _descriptor.EnumDescriptor(
  name='Kind',
  full_name='containersai.alameda.v1alpha1.datahub.resources.Kind',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='KIND_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DEPLOYMENT', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='DEPLOYMENTCONFIG', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='STATEFULSET', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='ALAMEDASCALER', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=377,
  serialized_end=477,
)
_sym_db.RegisterEnumDescriptor(_KIND)

Kind = enum_type_wrapper.EnumTypeWrapper(_KIND)
_SCALINGTOOL = _descriptor.EnumDescriptor(
  name='ScalingTool',
  full_name='containersai.alameda.v1alpha1.datahub.resources.ScalingTool',
  filename=None,
  file=DESCRIPTOR,
  create_key=_descriptor._internal_create_key,
  values=[
    _descriptor.EnumValueDescriptor(
      name='SCALING_TOOL_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='NONE', index=1, number=1,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='VPA', index=2, number=2,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='HPA', index=3, number=3,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
    _descriptor.EnumValueDescriptor(
      name='CA', index=4, number=4,
      serialized_options=None,
      type=None,
      create_key=_descriptor._internal_create_key),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=479,
  serialized_end=556,
)
_sym_db.RegisterEnumDescriptor(_SCALINGTOOL)

ScalingTool = enum_type_wrapper.EnumTypeWrapper(_SCALINGTOOL)
KIND_UNDEFINED = 0
DEPLOYMENT = 1
DEPLOYMENTCONFIG = 2
STATEFULSET = 3
ALAMEDASCALER = 4
SCALING_TOOL_UNDEFINED = 0
NONE = 1
VPA = 2
HPA = 3
CA = 4



_OBJECTMETA = _descriptor.Descriptor(
  name='ObjectMeta',
  full_name='containersai.alameda.v1alpha1.datahub.resources.ObjectMeta',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.alameda.v1alpha1.datahub.resources.ObjectMeta.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='namespace', full_name='containersai.alameda.v1alpha1.datahub.resources.ObjectMeta.namespace', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='node_name', full_name='containersai.alameda.v1alpha1.datahub.resources.ObjectMeta.node_name', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='cluster_name', full_name='containersai.alameda.v1alpha1.datahub.resources.ObjectMeta.cluster_name', index=3,
      number=4, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='uid', full_name='containersai.alameda.v1alpha1.datahub.resources.ObjectMeta.uid', index=4,
      number=5, type=9, cpp_type=9, label=1,
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
  serialized_start=106,
  serialized_end=205,
)


_OWNERREFERENCE = _descriptor.Descriptor(
  name='OwnerReference',
  full_name='containersai.alameda.v1alpha1.datahub.resources.OwnerReference',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  create_key=_descriptor._internal_create_key,
  fields=[
    _descriptor.FieldDescriptor(
      name='object_meta', full_name='containersai.alameda.v1alpha1.datahub.resources.OwnerReference.object_meta', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR,  create_key=_descriptor._internal_create_key),
    _descriptor.FieldDescriptor(
      name='kind', full_name='containersai.alameda.v1alpha1.datahub.resources.OwnerReference.kind', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=208,
  serialized_end=375,
)

_OWNERREFERENCE.fields_by_name['object_meta'].message_type = _OBJECTMETA
_OWNERREFERENCE.fields_by_name['kind'].enum_type = _KIND
DESCRIPTOR.message_types_by_name['ObjectMeta'] = _OBJECTMETA
DESCRIPTOR.message_types_by_name['OwnerReference'] = _OWNERREFERENCE
DESCRIPTOR.enum_types_by_name['Kind'] = _KIND
DESCRIPTOR.enum_types_by_name['ScalingTool'] = _SCALINGTOOL
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ObjectMeta = _reflection.GeneratedProtocolMessageType('ObjectMeta', (_message.Message,), {
  'DESCRIPTOR' : _OBJECTMETA,
  '__module__' : 'alameda_api.v1alpha1.datahub.resources.metadata_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.resources.ObjectMeta)
  })
_sym_db.RegisterMessage(ObjectMeta)

OwnerReference = _reflection.GeneratedProtocolMessageType('OwnerReference', (_message.Message,), {
  'DESCRIPTOR' : _OWNERREFERENCE,
  '__module__' : 'alameda_api.v1alpha1.datahub.resources.metadata_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.resources.OwnerReference)
  })
_sym_db.RegisterMessage(OwnerReference)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
