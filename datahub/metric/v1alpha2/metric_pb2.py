# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: datahub/metric/v1alpha2/metric.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from datahub.resource.metadata.v1alpha2 import metadata_pb2 as datahub_dot_resource_dot_metadata_dot_v1alpha2_dot_metadata__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='datahub/metric/v1alpha2/metric.proto',
  package='containersai.datahub.metric.v1alpha2',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n$datahub/metric/v1alpha2/metric.proto\x12$containersai.datahub.metric.v1alpha2\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x31\x64\x61tahub/resource/metadata/v1alpha2/metadata.proto\"\xe0\x01\n\x0f\x43ontainerMetric\x12\x0c\n\x04name\x18\x01 \x01(\t\x12Z\n\x0bmetric_data\x18\x02 \x03(\x0b\x32\x45.containersai.datahub.metric.v1alpha2.ContainerMetric.MetricDataEntry\x1a\x63\n\x0fMetricDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12?\n\x05value\x18\x02 \x01(\x0b\x32\x30.containersai.datahub.metric.v1alpha2.MetricData:\x02\x38\x01\"\xb7\x01\n\tPodMetric\x12X\n\x0fnamespaced_name\x18\x01 \x01(\x0b\x32?.containersai.datahub.resource.metadata.v1alpha2.NamespacedName\x12P\n\x11\x63ontainer_metrics\x18\x02 \x03(\x0b\x32\x35.containersai.datahub.metric.v1alpha2.ContainerMetric\"\xd6\x01\n\nNodeMetric\x12\x0c\n\x04name\x18\x01 \x01(\t\x12U\n\x0bmetric_data\x18\x02 \x03(\x0b\x32@.containersai.datahub.metric.v1alpha2.NodeMetric.MetricDataEntry\x1a\x63\n\x0fMetricDataEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12?\n\x05value\x18\x02 \x01(\x0b\x32\x30.containersai.datahub.metric.v1alpha2.MetricData:\x02\x38\x01\"E\n\x06Sample\x12(\n\x04time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x11\n\tnum_value\x18\x02 \x01(\t\"\x92\x01\n\tTimeRange\x12.\n\nstart_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\'\n\x04step\x18\x03 \x01(\x0b\x32\x19.google.protobuf.Duration\"H\n\nMetricData\x12:\n\x04\x64\x61ta\x18\x01 \x03(\x0b\x32,.containersai.datahub.metric.v1alpha2.Sample*U\n\nMetricType\x12\r\n\tUNDEFINED\x10\x00\x12 \n\x1c\x43PU_USAGE_SECONDS_PERCENTAGE\x10\x01\x12\x16\n\x12MEMORY_USAGE_BYTES\x10\x02\x62\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,datahub_dot_resource_dot_metadata_dot_v1alpha2_dot_metadata__pb2.DESCRIPTOR,])

_METRICTYPE = _descriptor.EnumDescriptor(
  name='MetricType',
  full_name='containersai.datahub.metric.v1alpha2.MetricType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CPU_USAGE_SECONDS_PERCENTAGE', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MEMORY_USAGE_BYTES', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1118,
  serialized_end=1203,
)
_sym_db.RegisterEnumDescriptor(_METRICTYPE)

MetricType = enum_type_wrapper.EnumTypeWrapper(_METRICTYPE)
UNDEFINED = 0
CPU_USAGE_SECONDS_PERCENTAGE = 1
MEMORY_USAGE_BYTES = 2



_CONTAINERMETRIC_METRICDATAENTRY = _descriptor.Descriptor(
  name='MetricDataEntry',
  full_name='containersai.datahub.metric.v1alpha2.ContainerMetric.MetricDataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containersai.datahub.metric.v1alpha2.ContainerMetric.MetricDataEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containersai.datahub.metric.v1alpha2.ContainerMetric.MetricDataEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=320,
  serialized_end=419,
)

_CONTAINERMETRIC = _descriptor.Descriptor(
  name='ContainerMetric',
  full_name='containersai.datahub.metric.v1alpha2.ContainerMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.datahub.metric.v1alpha2.ContainerMetric.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metric_data', full_name='containersai.datahub.metric.v1alpha2.ContainerMetric.metric_data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_CONTAINERMETRIC_METRICDATAENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=195,
  serialized_end=419,
)


_PODMETRIC = _descriptor.Descriptor(
  name='PodMetric',
  full_name='containersai.datahub.metric.v1alpha2.PodMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='namespaced_name', full_name='containersai.datahub.metric.v1alpha2.PodMetric.namespaced_name', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='container_metrics', full_name='containersai.datahub.metric.v1alpha2.PodMetric.container_metrics', index=1,
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
  serialized_start=422,
  serialized_end=605,
)


_NODEMETRIC_METRICDATAENTRY = _descriptor.Descriptor(
  name='MetricDataEntry',
  full_name='containersai.datahub.metric.v1alpha2.NodeMetric.MetricDataEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containersai.datahub.metric.v1alpha2.NodeMetric.MetricDataEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containersai.datahub.metric.v1alpha2.NodeMetric.MetricDataEntry.value', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=320,
  serialized_end=419,
)

_NODEMETRIC = _descriptor.Descriptor(
  name='NodeMetric',
  full_name='containersai.datahub.metric.v1alpha2.NodeMetric',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='name', full_name='containersai.datahub.metric.v1alpha2.NodeMetric.name', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='metric_data', full_name='containersai.datahub.metric.v1alpha2.NodeMetric.metric_data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_NODEMETRIC_METRICDATAENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=608,
  serialized_end=822,
)


_SAMPLE = _descriptor.Descriptor(
  name='Sample',
  full_name='containersai.datahub.metric.v1alpha2.Sample',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='containersai.datahub.metric.v1alpha2.Sample.time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='num_value', full_name='containersai.datahub.metric.v1alpha2.Sample.num_value', index=1,
      number=2, type=9, cpp_type=9, label=1,
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
  serialized_start=824,
  serialized_end=893,
)


_TIMERANGE = _descriptor.Descriptor(
  name='TimeRange',
  full_name='containersai.datahub.metric.v1alpha2.TimeRange',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containersai.datahub.metric.v1alpha2.TimeRange.start_time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='containersai.datahub.metric.v1alpha2.TimeRange.end_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='step', full_name='containersai.datahub.metric.v1alpha2.TimeRange.step', index=2,
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
  serialized_start=896,
  serialized_end=1042,
)


_METRICDATA = _descriptor.Descriptor(
  name='MetricData',
  full_name='containersai.datahub.metric.v1alpha2.MetricData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='data', full_name='containersai.datahub.metric.v1alpha2.MetricData.data', index=0,
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
  serialized_start=1044,
  serialized_end=1116,
)

_CONTAINERMETRIC_METRICDATAENTRY.fields_by_name['value'].message_type = _METRICDATA
_CONTAINERMETRIC_METRICDATAENTRY.containing_type = _CONTAINERMETRIC
_CONTAINERMETRIC.fields_by_name['metric_data'].message_type = _CONTAINERMETRIC_METRICDATAENTRY
_PODMETRIC.fields_by_name['namespaced_name'].message_type = datahub_dot_resource_dot_metadata_dot_v1alpha2_dot_metadata__pb2._NAMESPACEDNAME
_PODMETRIC.fields_by_name['container_metrics'].message_type = _CONTAINERMETRIC
_NODEMETRIC_METRICDATAENTRY.fields_by_name['value'].message_type = _METRICDATA
_NODEMETRIC_METRICDATAENTRY.containing_type = _NODEMETRIC
_NODEMETRIC.fields_by_name['metric_data'].message_type = _NODEMETRIC_METRICDATAENTRY
_SAMPLE.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE.fields_by_name['end_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE.fields_by_name['step'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_METRICDATA.fields_by_name['data'].message_type = _SAMPLE
DESCRIPTOR.message_types_by_name['ContainerMetric'] = _CONTAINERMETRIC
DESCRIPTOR.message_types_by_name['PodMetric'] = _PODMETRIC
DESCRIPTOR.message_types_by_name['NodeMetric'] = _NODEMETRIC
DESCRIPTOR.message_types_by_name['Sample'] = _SAMPLE
DESCRIPTOR.message_types_by_name['TimeRange'] = _TIMERANGE
DESCRIPTOR.message_types_by_name['MetricData'] = _METRICDATA
DESCRIPTOR.enum_types_by_name['MetricType'] = _METRICTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

ContainerMetric = _reflection.GeneratedProtocolMessageType('ContainerMetric', (_message.Message,), dict(

  MetricDataEntry = _reflection.GeneratedProtocolMessageType('MetricDataEntry', (_message.Message,), dict(
    DESCRIPTOR = _CONTAINERMETRIC_METRICDATAENTRY,
    __module__ = 'datahub.metric.v1alpha2.metric_pb2'
    # @@protoc_insertion_point(class_scope:containersai.datahub.metric.v1alpha2.ContainerMetric.MetricDataEntry)
    ))
  ,
  DESCRIPTOR = _CONTAINERMETRIC,
  __module__ = 'datahub.metric.v1alpha2.metric_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.metric.v1alpha2.ContainerMetric)
  ))
_sym_db.RegisterMessage(ContainerMetric)
_sym_db.RegisterMessage(ContainerMetric.MetricDataEntry)

PodMetric = _reflection.GeneratedProtocolMessageType('PodMetric', (_message.Message,), dict(
  DESCRIPTOR = _PODMETRIC,
  __module__ = 'datahub.metric.v1alpha2.metric_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.metric.v1alpha2.PodMetric)
  ))
_sym_db.RegisterMessage(PodMetric)

NodeMetric = _reflection.GeneratedProtocolMessageType('NodeMetric', (_message.Message,), dict(

  MetricDataEntry = _reflection.GeneratedProtocolMessageType('MetricDataEntry', (_message.Message,), dict(
    DESCRIPTOR = _NODEMETRIC_METRICDATAENTRY,
    __module__ = 'datahub.metric.v1alpha2.metric_pb2'
    # @@protoc_insertion_point(class_scope:containersai.datahub.metric.v1alpha2.NodeMetric.MetricDataEntry)
    ))
  ,
  DESCRIPTOR = _NODEMETRIC,
  __module__ = 'datahub.metric.v1alpha2.metric_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.metric.v1alpha2.NodeMetric)
  ))
_sym_db.RegisterMessage(NodeMetric)
_sym_db.RegisterMessage(NodeMetric.MetricDataEntry)

Sample = _reflection.GeneratedProtocolMessageType('Sample', (_message.Message,), dict(
  DESCRIPTOR = _SAMPLE,
  __module__ = 'datahub.metric.v1alpha2.metric_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.metric.v1alpha2.Sample)
  ))
_sym_db.RegisterMessage(Sample)

TimeRange = _reflection.GeneratedProtocolMessageType('TimeRange', (_message.Message,), dict(
  DESCRIPTOR = _TIMERANGE,
  __module__ = 'datahub.metric.v1alpha2.metric_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.metric.v1alpha2.TimeRange)
  ))
_sym_db.RegisterMessage(TimeRange)

MetricData = _reflection.GeneratedProtocolMessageType('MetricData', (_message.Message,), dict(
  DESCRIPTOR = _METRICDATA,
  __module__ = 'datahub.metric.v1alpha2.metric_pb2'
  # @@protoc_insertion_point(class_scope:containersai.datahub.metric.v1alpha2.MetricData)
  ))
_sym_db.RegisterMessage(MetricData)


_CONTAINERMETRIC_METRICDATAENTRY._options = None
_NODEMETRIC_METRICDATAENTRY._options = None
# @@protoc_insertion_point(module_scope)
