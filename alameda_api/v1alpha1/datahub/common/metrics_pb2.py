# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/common/metrics.proto

from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/common/metrics.proto',
  package='containersai.alameda.v1alpha1.datahub.common',
  syntax='proto3',
  serialized_options=b'Z@github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common',
  serialized_pb=b'\n1alameda_api/v1alpha1/datahub/common/metrics.proto\x12,containersai.alameda.v1alpha1.datahub.common\x1a\x1fgoogle/protobuf/timestamp.proto\"s\n\x06Sample\x12(\n\x04time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x11\n\tnum_value\x18\x03 \x01(\t\"\xb4\x01\n\nMetricData\x12M\n\x0bmetric_type\x18\x01 \x01(\x0e\x32\x38.containersai.alameda.v1alpha1.datahub.common.MetricType\x12\x42\n\x04\x64\x61ta\x18\x02 \x03(\x0b\x32\x34.containersai.alameda.v1alpha1.datahub.common.Sample\x12\x13\n\x0bgranularity\x18\x03 \x01(\x03*\xc3\x05\n\nMetricType\x12\x1a\n\x16METRICS_TYPE_UNDEFINED\x10\x00\x12\x15\n\x11\x43PU_SECONDS_TOTAL\x10\x01\x12\x18\n\x14\x43PU_MILLICORES_TOTAL\x10\x02\x12\x18\n\x14\x43PU_MILLICORES_AVAIL\x10\x03\x12\x18\n\x14\x43PU_MILLICORES_USAGE\x10\x04\x12\x1c\n\x18\x43PU_MILLICORES_USAGE_PCT\x10\x05\x12\x1e\n\x1a\x43PU_MILLICORES_ALLOCATABLE\x10\x06\x12\x16\n\x12MEMORY_BYTES_TOTAL\x10\x07\x12\x16\n\x12MEMORY_BYTES_AVAIL\x10\x08\x12\x16\n\x12MEMORY_BYTES_USAGE\x10\t\x12\x1a\n\x16MEMORY_BYTES_USAGE_PCT\x10\n\x12\x1c\n\x18MEMORY_BYTES_ALLOCATABLE\x10\x0b\x12\x12\n\x0e\x46S_BYTES_TOTAL\x10\x0c\x12\x12\n\x0e\x46S_BYTES_AVAIL\x10\r\x12\x12\n\x0e\x46S_BYTES_USAGE\x10\x0e\x12\x16\n\x12\x46S_BYTES_USAGE_PCT\x10\x0f\x12\x17\n\x13HTTP_REQUESTS_COUNT\x10\x10\x12\x17\n\x13HTTP_REQUESTS_TOTAL\x10\x11\x12\x17\n\x13HTTP_RESPONSE_COUNT\x10\x12\x12\x17\n\x13HTTP_RESPONSE_TOTAL\x10\x13\x12\x19\n\x15\x44ISK_IO_SECONDS_TOTAL\x10\x14\x12\x17\n\x13\x44ISK_IO_UTILIZATION\x10\x15\x12\x12\n\x0eRESTARTS_TOTAL\x10\x16\x12\x15\n\x11POWER_USAGE_WATTS\x10\x17\x12\x17\n\x13TEMPERATURE_CELSIUS\x10\x18\x12\x0e\n\nDUTY_CYCLE\x10\x19\x12\x12\n\x0e\x43URRENT_OFFSET\x10\x1a\x12\x07\n\x03LAG\x10\x1b\x12\x0b\n\x07LATENCY\x10\x1c\x12\n\n\x06NUMBER\x10\x1d*@\n\x0cResourceName\x12\x1b\n\x17RESOURCE_NAME_UNDEFINED\x10\x00\x12\x07\n\x03\x43PU\x10\x01\x12\n\n\x06MEMORY\x10\x02\x42\x42Z@github.com/containers-ai/api/alameda_api/v1alpha1/datahub/commonb\x06proto3'
  ,
  dependencies=[google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])

_METRICTYPE = _descriptor.EnumDescriptor(
  name='MetricType',
  full_name='containersai.alameda.v1alpha1.datahub.common.MetricType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='METRICS_TYPE_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CPU_SECONDS_TOTAL', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CPU_MILLICORES_TOTAL', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CPU_MILLICORES_AVAIL', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CPU_MILLICORES_USAGE', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CPU_MILLICORES_USAGE_PCT', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CPU_MILLICORES_ALLOCATABLE', index=6, number=6,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MEMORY_BYTES_TOTAL', index=7, number=7,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MEMORY_BYTES_AVAIL', index=8, number=8,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MEMORY_BYTES_USAGE', index=9, number=9,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MEMORY_BYTES_USAGE_PCT', index=10, number=10,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MEMORY_BYTES_ALLOCATABLE', index=11, number=11,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FS_BYTES_TOTAL', index=12, number=12,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FS_BYTES_AVAIL', index=13, number=13,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FS_BYTES_USAGE', index=14, number=14,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='FS_BYTES_USAGE_PCT', index=15, number=15,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTTP_REQUESTS_COUNT', index=16, number=16,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTTP_REQUESTS_TOTAL', index=17, number=17,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTTP_RESPONSE_COUNT', index=18, number=18,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTTP_RESPONSE_TOTAL', index=19, number=19,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DISK_IO_SECONDS_TOTAL', index=20, number=20,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DISK_IO_UTILIZATION', index=21, number=21,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='RESTARTS_TOTAL', index=22, number=22,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='POWER_USAGE_WATTS', index=23, number=23,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='TEMPERATURE_CELSIUS', index=24, number=24,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DUTY_CYCLE', index=25, number=25,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CURRENT_OFFSET', index=26, number=26,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='LAG', index=27, number=27,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='LATENCY', index=28, number=28,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NUMBER', index=29, number=29,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=433,
  serialized_end=1140,
)
_sym_db.RegisterEnumDescriptor(_METRICTYPE)

MetricType = enum_type_wrapper.EnumTypeWrapper(_METRICTYPE)
_RESOURCENAME = _descriptor.EnumDescriptor(
  name='ResourceName',
  full_name='containersai.alameda.v1alpha1.datahub.common.ResourceName',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='RESOURCE_NAME_UNDEFINED', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='CPU', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MEMORY', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1142,
  serialized_end=1206,
)
_sym_db.RegisterEnumDescriptor(_RESOURCENAME)

ResourceName = enum_type_wrapper.EnumTypeWrapper(_RESOURCENAME)
METRICS_TYPE_UNDEFINED = 0
CPU_SECONDS_TOTAL = 1
CPU_MILLICORES_TOTAL = 2
CPU_MILLICORES_AVAIL = 3
CPU_MILLICORES_USAGE = 4
CPU_MILLICORES_USAGE_PCT = 5
CPU_MILLICORES_ALLOCATABLE = 6
MEMORY_BYTES_TOTAL = 7
MEMORY_BYTES_AVAIL = 8
MEMORY_BYTES_USAGE = 9
MEMORY_BYTES_USAGE_PCT = 10
MEMORY_BYTES_ALLOCATABLE = 11
FS_BYTES_TOTAL = 12
FS_BYTES_AVAIL = 13
FS_BYTES_USAGE = 14
FS_BYTES_USAGE_PCT = 15
HTTP_REQUESTS_COUNT = 16
HTTP_REQUESTS_TOTAL = 17
HTTP_RESPONSE_COUNT = 18
HTTP_RESPONSE_TOTAL = 19
DISK_IO_SECONDS_TOTAL = 20
DISK_IO_UTILIZATION = 21
RESTARTS_TOTAL = 22
POWER_USAGE_WATTS = 23
TEMPERATURE_CELSIUS = 24
DUTY_CYCLE = 25
CURRENT_OFFSET = 26
LAG = 27
LATENCY = 28
NUMBER = 29
RESOURCE_NAME_UNDEFINED = 0
CPU = 1
MEMORY = 2



_SAMPLE = _descriptor.Descriptor(
  name='Sample',
  full_name='containersai.alameda.v1alpha1.datahub.common.Sample',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time', full_name='containersai.alameda.v1alpha1.datahub.common.Sample.time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='containersai.alameda.v1alpha1.datahub.common.Sample.end_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='num_value', full_name='containersai.alameda.v1alpha1.datahub.common.Sample.num_value', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=b"".decode('utf-8'),
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
  serialized_start=132,
  serialized_end=247,
)


_METRICDATA = _descriptor.Descriptor(
  name='MetricData',
  full_name='containersai.alameda.v1alpha1.datahub.common.MetricData',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='metric_type', full_name='containersai.alameda.v1alpha1.datahub.common.MetricData.metric_type', index=0,
      number=1, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='data', full_name='containersai.alameda.v1alpha1.datahub.common.MetricData.data', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='granularity', full_name='containersai.alameda.v1alpha1.datahub.common.MetricData.granularity', index=2,
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
  serialized_start=250,
  serialized_end=430,
)

_SAMPLE.fields_by_name['time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_SAMPLE.fields_by_name['end_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_METRICDATA.fields_by_name['metric_type'].enum_type = _METRICTYPE
_METRICDATA.fields_by_name['data'].message_type = _SAMPLE
DESCRIPTOR.message_types_by_name['Sample'] = _SAMPLE
DESCRIPTOR.message_types_by_name['MetricData'] = _METRICDATA
DESCRIPTOR.enum_types_by_name['MetricType'] = _METRICTYPE
DESCRIPTOR.enum_types_by_name['ResourceName'] = _RESOURCENAME
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

Sample = _reflection.GeneratedProtocolMessageType('Sample', (_message.Message,), {
  'DESCRIPTOR' : _SAMPLE,
  '__module__' : 'alameda_api.v1alpha1.datahub.common.metrics_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.common.Sample)
  })
_sym_db.RegisterMessage(Sample)

MetricData = _reflection.GeneratedProtocolMessageType('MetricData', (_message.Message,), {
  'DESCRIPTOR' : _METRICDATA,
  '__module__' : 'alameda_api.v1alpha1.datahub.common.metrics_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.common.MetricData)
  })
_sym_db.RegisterMessage(MetricData)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
