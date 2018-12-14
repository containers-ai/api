# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/metric.proto

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


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/metric.proto',
  package='containers_ai.alameda.v1alpha1.datahub',
  syntax='proto3',
  serialized_options=None,
  serialized_pb=_b('\n)alameda_api/v1alpha1/datahub/metric.proto\x12&containers_ai.alameda.v1alpha1.datahub\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1egoogle/protobuf/duration.proto\"k\n\x06Sample\x12(\n\x04time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x16\n\x0c\x64ouble_value\x18\x02 \x01(\x01H\x00\x12\x16\n\x0cstring_value\x18\x03 \x01(\tH\x00\x42\x07\n\x05value\"\xd0\x01\n\x0cMetricResult\x12P\n\x06labels\x18\x01 \x03(\x0b\x32@.containers_ai.alameda.v1alpha1.datahub.MetricResult.LabelsEntry\x12?\n\x07samples\x18\x02 \x03(\x0b\x32..containers_ai.alameda.v1alpha1.datahub.Sample\x1a-\n\x0bLabelsEntry\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\r\n\x05value\x18\x02 \x01(\t:\x02\x38\x01\"\x92\x01\n\tTimeRange\x12.\n\nstart_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\'\n\x04step\x18\x03 \x01(\x0b\x32\x19.google.protobuf.Duration\"f\n\rLabelSelector\x12\x0b\n\x03key\x18\x01 \x01(\t\x12\x39\n\x02op\x18\x02 \x01(\x0e\x32-.containers_ai.alameda.v1alpha1.datahub.StrOp\x12\r\n\x05value\x18\x03 \x01(\t*\xb5\x01\n\nMetricType\x12\x18\n\x14METRICTYPE_UNDEFINED\x10\x00\x12,\n(CONTAINER_CPU_USAGE_SECONDS_TOTAL_RATE5M\x10\x01\x12 \n\x1c\x43ONTAINER_MEMORY_USAGE_BYTES\x10\x02\x12 \n\x1cNODE_CPU_USAGE_SECONDS_AVG1M\x10\x03\x12\x1b\n\x17NODE_MEMORY_USAGE_BYTES\x10\x04* \n\x05StrOp\x12\t\n\x05\x45qual\x10\x00\x12\x0c\n\x08NotEqual\x10\x01\x62\x06proto3')
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,])

_METRICTYPE = _descriptor.EnumDescriptor(
  name='MetricType',
  full_name='containers_ai.alameda.v1alpha1.datahub.MetricType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='METRICTYPE_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CONTAINER_CPU_USAGE_SECONDS_TOTAL_RATE5M', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CONTAINER_MEMORY_USAGE_BYTES', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NODE_CPU_USAGE_SECONDS_AVG1M', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NODE_MEMORY_USAGE_BYTES', index=4, number=4,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=724,
  serialized_end=905,
)
_sym_db.RegisterEnumDescriptor(_METRICTYPE)

MetricType = enum_type_wrapper.EnumTypeWrapper(_METRICTYPE)
_STROP = _descriptor.EnumDescriptor(
  name='StrOp',
  full_name='containers_ai.alameda.v1alpha1.datahub.StrOp',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='Equal', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NotEqual', index=1, number=1,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=907,
  serialized_end=939,
)
_sym_db.RegisterEnumDescriptor(_STROP)

StrOp = enum_type_wrapper.EnumTypeWrapper(_STROP)
METRICTYPE_UNDEFINED = 0
CONTAINER_CPU_USAGE_SECONDS_TOTAL_RATE5M = 1
CONTAINER_MEMORY_USAGE_BYTES = 2
NODE_CPU_USAGE_SECONDS_AVG1M = 3
NODE_MEMORY_USAGE_BYTES = 4
Equal = 0
NotEqual = 1



_SAMPLE = _descriptor.Descriptor(
  name='Sample',
  full_name='containers_ai.alameda.v1alpha1.datahub.Sample',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='containers_ai.alameda.v1alpha1.datahub.Sample.time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='double_value', full_name='containers_ai.alameda.v1alpha1.datahub.Sample.double_value', index=1,
      number=2, type=1, cpp_type=5, label=1,
      has_default_value=False, default_value=float(0),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='string_value', full_name='containers_ai.alameda.v1alpha1.datahub.Sample.string_value', index=2,
      number=3, type=9, cpp_type=9, label=1,
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
    _descriptor.OneofDescriptor(
      name='value', full_name='containers_ai.alameda.v1alpha1.datahub.Sample.value',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=150,
  serialized_end=257,
)


_METRICRESULT_LABELSENTRY = _descriptor.Descriptor(
  name='LabelsEntry',
  full_name='containers_ai.alameda.v1alpha1.datahub.MetricResult.LabelsEntry',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containers_ai.alameda.v1alpha1.datahub.MetricResult.LabelsEntry.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containers_ai.alameda.v1alpha1.datahub.MetricResult.LabelsEntry.value', index=1,
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
  serialized_options=_b('8\001'),
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=423,
  serialized_end=468,
)

_METRICRESULT = _descriptor.Descriptor(
  name='MetricResult',
  full_name='containers_ai.alameda.v1alpha1.datahub.MetricResult',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='labels', full_name='containers_ai.alameda.v1alpha1.datahub.MetricResult.labels', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='samples', full_name='containers_ai.alameda.v1alpha1.datahub.MetricResult.samples', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[_METRICRESULT_LABELSENTRY, ],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=260,
  serialized_end=468,
)


_TIMERANGE = _descriptor.Descriptor(
  name='TimeRange',
  full_name='containers_ai.alameda.v1alpha1.datahub.TimeRange',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containers_ai.alameda.v1alpha1.datahub.TimeRange.start_time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='containers_ai.alameda.v1alpha1.datahub.TimeRange.end_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='step', full_name='containers_ai.alameda.v1alpha1.datahub.TimeRange.step', index=2,
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
  serialized_start=471,
  serialized_end=617,
)


_LABELSELECTOR = _descriptor.Descriptor(
  name='LabelSelector',
  full_name='containers_ai.alameda.v1alpha1.datahub.LabelSelector',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key', full_name='containers_ai.alameda.v1alpha1.datahub.LabelSelector.key', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='op', full_name='containers_ai.alameda.v1alpha1.datahub.LabelSelector.op', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='value', full_name='containers_ai.alameda.v1alpha1.datahub.LabelSelector.value', index=2,
      number=3, type=9, cpp_type=9, label=1,
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
  serialized_start=619,
  serialized_end=721,
)

_SAMPLE.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_SAMPLE.oneofs_by_name['value'].fields.append(
  _SAMPLE.fields_by_name['double_value'])
_SAMPLE.fields_by_name['double_value'].containing_oneof = _SAMPLE.oneofs_by_name['value']
_SAMPLE.oneofs_by_name['value'].fields.append(
  _SAMPLE.fields_by_name['string_value'])
_SAMPLE.fields_by_name['string_value'].containing_oneof = _SAMPLE.oneofs_by_name['value']
_METRICRESULT_LABELSENTRY.containing_type = _METRICRESULT
_METRICRESULT.fields_by_name['labels'].message_type = _METRICRESULT_LABELSENTRY
_METRICRESULT.fields_by_name['samples'].message_type = _SAMPLE
_TIMERANGE.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE.fields_by_name['end_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE.fields_by_name['step'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_LABELSELECTOR.fields_by_name['op'].enum_type = _STROP
DESCRIPTOR.message_types_by_name['Sample'] = _SAMPLE
DESCRIPTOR.message_types_by_name['MetricResult'] = _METRICRESULT
DESCRIPTOR.message_types_by_name['TimeRange'] = _TIMERANGE
DESCRIPTOR.message_types_by_name['LabelSelector'] = _LABELSELECTOR
DESCRIPTOR.enum_types_by_name['MetricType'] = _METRICTYPE
DESCRIPTOR.enum_types_by_name['StrOp'] = _STROP
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Sample = _reflection.GeneratedProtocolMessageType('Sample', (_message.Message,), dict(
  DESCRIPTOR = _SAMPLE,
  __module__ = 'alameda_api.v1alpha1.datahub.metric_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.Sample)
  ))
_sym_db.RegisterMessage(Sample)

MetricResult = _reflection.GeneratedProtocolMessageType('MetricResult', (_message.Message,), dict(

  LabelsEntry = _reflection.GeneratedProtocolMessageType('LabelsEntry', (_message.Message,), dict(
    DESCRIPTOR = _METRICRESULT_LABELSENTRY,
    __module__ = 'alameda_api.v1alpha1.datahub.metric_pb2'
    # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.MetricResult.LabelsEntry)
    ))
  ,
  DESCRIPTOR = _METRICRESULT,
  __module__ = 'alameda_api.v1alpha1.datahub.metric_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.MetricResult)
  ))
_sym_db.RegisterMessage(MetricResult)
_sym_db.RegisterMessage(MetricResult.LabelsEntry)

TimeRange = _reflection.GeneratedProtocolMessageType('TimeRange', (_message.Message,), dict(
  DESCRIPTOR = _TIMERANGE,
  __module__ = 'alameda_api.v1alpha1.datahub.metric_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.TimeRange)
  ))
_sym_db.RegisterMessage(TimeRange)

LabelSelector = _reflection.GeneratedProtocolMessageType('LabelSelector', (_message.Message,), dict(
  DESCRIPTOR = _LABELSELECTOR,
  __module__ = 'alameda_api.v1alpha1.datahub.metric_pb2'
  # @@protoc_insertion_point(class_scope:containers_ai.alameda.v1alpha1.datahub.LabelSelector)
  ))
_sym_db.RegisterMessage(LabelSelector)


_METRICRESULT_LABELSENTRY._options = None
# @@protoc_insertion_point(module_scope)
