# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/resource.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from alameda_api.v1alpha1.datahub import predict_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_predict__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from alameda_api.v1alpha1.datahub import metadata_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_metadata__pb2
from alameda_api.v1alpha1.datahub import metric_pb2 as alameda__api_dot_v1alpha1_dot_datahub_dot_metric__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/resource.proto',
  package='containers_ai.alameda.v1alpha1.datahub',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n+alameda_api/v1alpha1/datahub/resource.proto\x12&containers_ai.alameda.v1alpha1.datahub\x1a*alameda_api/v1alpha1/datahub/predict.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a+alameda_api/v1alpha1/datahub/metadata.proto\x1a)alameda_api/v1alpha1/datahub/metric.proto\"\xb3\x01\n\tContainer\x12\x0c\n\x04name\x18\x01 \x01(\t\x12J\n\x0elimit_resource\x18\x02 \x03(\x0b\x32\x32.containers_ai.alameda.v1alpha1.datahub.MetricData\x12L\n\x10request_resource\x18\x03 \x03(\x0b\x32\x32.containers_ai.alameda.v1alpha1.datahub.MetricData\"\x98\x04\n\x03Pod\x12O\n\x0fnamespaced_name\x18\x01 \x01(\x0b\x32\x36.containers_ai.alameda.v1alpha1.datahub.NamespacedName\x12\x15\n\rresource_link\x18\x02 \x01(\t\x12\x45\n\ncontainers\x18\x03 \x03(\x0b\x32\x31.containers_ai.alameda.v1alpha1.datahub.Container\x12\x12\n\nis_alameda\x18\x04 \x01(\x08\x12N\n\x0e\x61lameda_scaler\x18\x05 \x01(\x0b\x32\x36.containers_ai.alameda.v1alpha1.datahub.NamespacedName\x12\x11\n\tnode_name\x18\x06 \x01(\t\x12.\n\nstart_time\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12L\n\x06policy\x18\x08 \x01(\x0e\x32<.containers_ai.alameda.v1alpha1.datahub.RecommendationPolicy\x12M\n\x0etop_controller\x18\t \x01(\x0b\x32\x35.containers_ai.alameda.v1alpha1.datahub.TopController\x12\x1e\n\x16used_recommendation_id\x18\n \x01(\t\"X\n\x08\x43\x61pacity\x12\x11\n\tcpu_cores\x18\x01 \x01(\x03\x12\x14\n\x0cmemory_bytes\x18\x02 \x01(\x03\x12#\n\x1bnetwotk_megabits_per_second\x18\x03 \x01(\x03\"\x88\x01\n\x04Node\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x42\n\x08\x63\x61pacity\x18\x02 \x01(\x0b\x32\x30.containers_ai.alameda.v1alpha1.datahub.Capacity\x12.\n\nstart_time\x18\x03 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"\x9c\x01\n\rTopController\x12O\n\x0fnamespaced_name\x18\x01 \x01(\x0b\x32\x36.containers_ai.alameda.v1alpha1.datahub.NamespacedName\x12:\n\x04kind\x18\x02 \x01(\x0e\x32,.containers_ai.alameda.v1alpha1.datahub.Kind*H\n\x04Kind\x12\x07\n\x03POD\x10\x00\x12\x0e\n\nDEPLOYMENT\x10\x01\x12\x14\n\x10\x44\x45PLOYMENTCONFIG\x10\x02\x12\x11\n\rALAMEDASCALER\x10\x03\x62\x06proto3')
  ,
  dependencies=[alameda__api_dot_v1alpha1_dot_datahub_dot_predict__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_metadata__pb2.DESCRIPTOR,alameda__api_dot_v1alpha1_dot_datahub_dot_metric__pb2.DESCRIPTOR,])

_KIND = _descriptor.EnumDescriptor(
  name='Kind',
  full_name='containers_ai.alameda.v1alpha1.datahub.Kind',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='POD', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DEPLOYMENT', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DEPLOYMENTCONFIG', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ALAMEDASCALER', index=3, number=3,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1361,
  serialized_end=1433,
)
_sym_db.RegisterEnumDescriptor(_KIND)

Kind = enum_type_wrapper.EnumTypeWrapper(_KIND)
POD = 0
DEPLOYMENT = 1
DEPLOYMENTCONFIG = 2
ALAMEDASCALER = 3



_CONTAINER = _descriptor.Descriptor(
  name='Container',
  full_name='containers_ai.alameda.v1alpha1.datahub.Container',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containers_ai.alameda.v1alpha1.datahub.Container.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit_resource', full_name='containers_ai.alameda.v1alpha1.datahub.Container.limit_resource', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='request_resource', full_name='containers_ai.alameda.v1alpha1.datahub.Container.request_resource', index=2,
      number=3, type=11, cpp_type=10, label=3,
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
  serialized_start=253,
  serialized_end=432,
)


_POD = _descriptor.Descriptor(
  name='Pod',
  full_name='containers_ai.alameda.v1alpha1.datahub.Pod',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespaced_name', full_name='containers_ai.alameda.v1alpha1.datahub.Pod.namespaced_name', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='resource_link', full_name='containers_ai.alameda.v1alpha1.datahub.Pod.resource_link', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='containers', full_name='containers_ai.alameda.v1alpha1.datahub.Pod.containers', index=2,
      number=3, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='is_alameda', full_name='containers_ai.alameda.v1alpha1.datahub.Pod.is_alameda', index=3,
      number=4, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='alameda_scaler', full_name='containers_ai.alameda.v1alpha1.datahub.Pod.alameda_scaler', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='node_name', full_name='containers_ai.alameda.v1alpha1.datahub.Pod.node_name', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containers_ai.alameda.v1alpha1.datahub.Pod.start_time', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='policy', full_name='containers_ai.alameda.v1alpha1.datahub.Pod.policy', index=7,
      number=8, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='top_controller', full_name='containers_ai.alameda.v1alpha1.datahub.Pod.top_controller', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='used_recommendation_id', full_name='containers_ai.alameda.v1alpha1.datahub.Pod.used_recommendation_id', index=9,
      number=10, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
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
  serialized_start=435,
  serialized_end=971,
)


_CAPACITY = _descriptor.Descriptor(
  name='Capacity',
  full_name='containers_ai.alameda.v1alpha1.datahub.Capacity',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='cpu_cores', full_name='containers_ai.alameda.v1alpha1.datahub.Capacity.cpu_cores', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='memory_bytes', full_name='containers_ai.alameda.v1alpha1.datahub.Capacity.memory_bytes', index=1,
      number=2, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='netwotk_megabits_per_second', full_name='containers_ai.alameda.v1alpha1.datahub.Capacity.netwotk_megabits_per_second', index=2,
      number=3, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=973,
  serialized_end=1061,
)


_NODE = _descriptor.Descriptor(
  name='Node',
  full_name='containers_ai.alameda.v1alpha1.datahub.Node',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containers_ai.alameda.v1alpha1.datahub.Node.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='capacity', full_name='containers_ai.alameda.v1alpha1.datahub.Node.capacity', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containers_ai.alameda.v1alpha1.datahub.Node.start_time', index=2,
      number=3, type=11, cpp_type=10, label=1,
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
  serialized_start=1064,
  serialized_end=1200,
)


_TOPCONTROLLER = _descriptor.Descriptor(
  name='TopController',
  full_name='containers_ai.alameda.v1alpha1.datahub.TopController',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespaced_name', full_name='containers_ai.alameda.v1alpha1.datahub.TopController.namespaced_name', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='kind', full_name='containers_ai.alameda.v1alpha1.datahub.TopController.kind', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
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
  serialized_start=1203,
  serialized_end=1359,
)

_CONTAINER.fields_by_name['limit_resource'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_metric__pb2._METRICDATA
_CONTAINER.fields_by_name['request_resource'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_metric__pb2._METRICDATA
_POD.fields_by_name['namespaced_name'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_metadata__pb2._NAMESPACEDNAME
_POD.fields_by_name['containers'].message_type = _CONTAINER
_POD.fields_by_name['alameda_scaler'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_metadata__pb2._NAMESPACEDNAME
_POD.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_POD.fields_by_name['policy'].enum_type = alameda__api_dot_v1alpha1_dot_datahub_dot_predict__pb2._RECOMMENDATIONPOLICY
_POD.fields_by_name['top_controller'].message_type = _TOPCONTROLLER
_NODE.fields_by_name['capacity'].message_type = _CAPACITY
_NODE.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TOPCONTROLLER.fields_by_name['namespaced_name'].message_type = alameda__api_dot_v1alpha1_dot_datahub_dot_metadata__pb2._NAMESPACEDNAME
_TOPCONTROLLER.fields_by_name['kind'].enum_type = _KIND
DESCRIPTOR.message_types_by_name['Container'] = _CONTAINER
DESCRIPTOR.message_types_by_name['Pod'] = _POD
DESCRIPTOR.message_types_by_name['Capacity'] = _CAPACITY
DESCRIPTOR.message_types_by_name['Node'] = _NODE
DESCRIPTOR.message_types_by_name['TopController'] = _TOPCONTROLLER
DESCRIPTOR.enum_types_by_name['Kind'] = _KIND
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Container = _reflection.GeneratedProtocolMessageType('Container', (_message.Message,), dict(
  DESCRIPTOR = _CONTAINER,
  __module__ = 'alameda_api.v1alpha1.datahub.resource_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.Container)
  ))
_sym_db.RegisterMessage(Container)

Pod = _reflection.GeneratedProtocolMessageType('Pod', (_message.Message,), dict(
  DESCRIPTOR = _POD,
  __module__ = 'alameda_api.v1alpha1.datahub.resource_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.Pod)
  ))
_sym_db.RegisterMessage(Pod)

Capacity = _reflection.GeneratedProtocolMessageType('Capacity', (_message.Message,), dict(
  DESCRIPTOR = _CAPACITY,
  __module__ = 'alameda_api.v1alpha1.datahub.resource_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.Capacity)
  ))
_sym_db.RegisterMessage(Capacity)

Node = _reflection.GeneratedProtocolMessageType('Node', (_message.Message,), dict(
  DESCRIPTOR = _NODE,
  __module__ = 'alameda_api.v1alpha1.datahub.resource_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.Node)
  ))
_sym_db.RegisterMessage(Node)

TopController = _reflection.GeneratedProtocolMessageType('TopController', (_message.Message,), dict(
  DESCRIPTOR = _TOPCONTROLLER,
  __module__ = 'alameda_api.v1alpha1.datahub.resource_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.TopController)
  ))
_sym_db.RegisterMessage(TopController)


# @@protoc_insertion_point(module_scope)
