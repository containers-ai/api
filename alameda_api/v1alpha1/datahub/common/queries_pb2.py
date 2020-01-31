# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: alameda_api/v1alpha1/datahub/common/queries.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from common import types_pb2 as common_dot_types__pb2
from google.protobuf import duration_pb2 as google_dot_protobuf_dot_duration__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='alameda_api/v1alpha1/datahub/common/queries.proto',
  package='containersai.alameda.v1alpha1.datahub.common',
  syntax='proto3',
  serialized_options=_b('Z@github.com/containers-ai/api/alameda_api/v1alpha1/datahub/common'),
  serialized_pb=_b('\n1alameda_api/v1alpha1/datahub/common/queries.proto\x12,containersai.alameda.v1alpha1.datahub.common\x1a\x12\x63ommon/types.proto\x1a\x1egoogle/protobuf/duration.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"\xd9\x02\n\tTimeRange\x12.\n\nstart_time\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12,\n\x08\x65nd_time\x18\x02 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\'\n\x04step\x18\x03 \x01(\x0b\x32\x19.google.protobuf.Duration\x12\x64\n\x11\x61ggregateFunction\x18\x04 \x01(\x0e\x32I.containersai.alameda.v1alpha1.datahub.common.TimeRange.AggregateFunction\x12.\n\napply_time\x18\x05 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\"/\n\x11\x41ggregateFunction\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03MAX\x10\x01\x12\x07\n\x03\x41VG\x10\x02\"j\n\tCondition\x12\x0c\n\x04keys\x18\x01 \x03(\t\x12\x0e\n\x06values\x18\x02 \x03(\t\x12\x11\n\toperators\x18\x03 \x03(\t\x12,\n\x05types\x18\x04 \x03(\x0e\x32\x1d.containersai.common.DataType\"\xdd\x02\n\x0eQueryCondition\x12K\n\ntime_range\x18\x01 \x01(\x0b\x32\x37.containersai.alameda.v1alpha1.datahub.common.TimeRange\x12Q\n\x05order\x18\x02 \x01(\x0e\x32\x42.containersai.alameda.v1alpha1.datahub.common.QueryCondition.Order\x12\x14\n\x0cwhere_clause\x18\x03 \x01(\t\x12P\n\x0fwhere_condition\x18\x04 \x03(\x0b\x32\x37.containersai.alameda.v1alpha1.datahub.common.Condition\x12\x0e\n\x06groups\x18\x05 \x03(\t\x12\r\n\x05limit\x18\x06 \x01(\x04\"$\n\x05Order\x12\x08\n\x04NONE\x10\x00\x12\x07\n\x03\x41SC\x10\x01\x12\x08\n\x04\x44\x45SC\x10\x02\x42\x42Z@github.com/containers-ai/api/alameda_api/v1alpha1/datahub/commonb\x06proto3')
  ,
  dependencies=[common_dot_types__pb2.DESCRIPTOR,google_dot_protobuf_dot_duration__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,])



_TIMERANGE_AGGREGATEFUNCTION = _descriptor.EnumDescriptor(
  name='AggregateFunction',
  full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.AggregateFunction',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='MAX', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='AVG', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=483,
  serialized_end=530,
)
_sym_db.RegisterEnumDescriptor(_TIMERANGE_AGGREGATEFUNCTION)

_QUERYCONDITION_ORDER = _descriptor.EnumDescriptor(
  name='Order',
  full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.Order',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='NONE', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='ASC', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='DESC', index=2, number=2,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=954,
  serialized_end=990,
)
_sym_db.RegisterEnumDescriptor(_QUERYCONDITION_ORDER)


_TIMERANGE = _descriptor.Descriptor(
  name='TimeRange',
  full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='start_time', full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.start_time', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='end_time', full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.end_time', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='step', full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.step', index=2,
      number=3, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='aggregateFunction', full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.aggregateFunction', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='apply_time', full_name='containersai.alameda.v1alpha1.datahub.common.TimeRange.apply_time', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _TIMERANGE_AGGREGATEFUNCTION,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=185,
  serialized_end=530,
)


_CONDITION = _descriptor.Descriptor(
  name='Condition',
  full_name='containersai.alameda.v1alpha1.datahub.common.Condition',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='keys', full_name='containersai.alameda.v1alpha1.datahub.common.Condition.keys', index=0,
      number=1, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='values', full_name='containersai.alameda.v1alpha1.datahub.common.Condition.values', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='operators', full_name='containersai.alameda.v1alpha1.datahub.common.Condition.operators', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='types', full_name='containersai.alameda.v1alpha1.datahub.common.Condition.types', index=3,
      number=4, type=14, cpp_type=8, label=3,
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
  serialized_start=532,
  serialized_end=638,
)


_QUERYCONDITION = _descriptor.Descriptor(
  name='QueryCondition',
  full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='time_range', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.time_range', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='order', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.order', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='where_clause', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.where_clause', index=2,
      number=3, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='where_condition', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.where_condition', index=3,
      number=4, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='groups', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.groups', index=4,
      number=5, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='limit', full_name='containersai.alameda.v1alpha1.datahub.common.QueryCondition.limit', index=5,
      number=6, type=4, cpp_type=4, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
    _QUERYCONDITION_ORDER,
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=641,
  serialized_end=990,
)

_TIMERANGE.fields_by_name['start_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE.fields_by_name['end_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE.fields_by_name['step'].message_type = google_dot_protobuf_dot_duration__pb2._DURATION
_TIMERANGE.fields_by_name['aggregateFunction'].enum_type = _TIMERANGE_AGGREGATEFUNCTION
_TIMERANGE.fields_by_name['apply_time'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TIMERANGE_AGGREGATEFUNCTION.containing_type = _TIMERANGE
_CONDITION.fields_by_name['types'].enum_type = common_dot_types__pb2._DATATYPE
_QUERYCONDITION.fields_by_name['time_range'].message_type = _TIMERANGE
_QUERYCONDITION.fields_by_name['order'].enum_type = _QUERYCONDITION_ORDER
_QUERYCONDITION.fields_by_name['where_condition'].message_type = _CONDITION
_QUERYCONDITION_ORDER.containing_type = _QUERYCONDITION
DESCRIPTOR.message_types_by_name['TimeRange'] = _TIMERANGE
DESCRIPTOR.message_types_by_name['Condition'] = _CONDITION
DESCRIPTOR.message_types_by_name['QueryCondition'] = _QUERYCONDITION
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

TimeRange = _reflection.GeneratedProtocolMessageType('TimeRange', (_message.Message,), {
  'DESCRIPTOR' : _TIMERANGE,
  '__module__' : 'alameda_api.v1alpha1.datahub.common.queries_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.common.TimeRange)
  })
_sym_db.RegisterMessage(TimeRange)

Condition = _reflection.GeneratedProtocolMessageType('Condition', (_message.Message,), {
  'DESCRIPTOR' : _CONDITION,
  '__module__' : 'alameda_api.v1alpha1.datahub.common.queries_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.common.Condition)
  })
_sym_db.RegisterMessage(Condition)

QueryCondition = _reflection.GeneratedProtocolMessageType('QueryCondition', (_message.Message,), {
  'DESCRIPTOR' : _QUERYCONDITION,
  '__module__' : 'alameda_api.v1alpha1.datahub.common.queries_pb2'
  # @@protoc_insertion_point(class_scope:containersai.alameda.v1alpha1.datahub.common.QueryCondition)
  })
_sym_db.RegisterMessage(QueryCondition)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
