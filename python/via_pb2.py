# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: via.proto
"""Generated protocol buffer code."""
from google.protobuf.internal import builder as _builder
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\tvia.proto\x12\x05viapb\"\xdc\x01\n\x08Incident\x12\n\n\x02id\x18\x01 \x01(\t\x12\x0b\n\x03lat\x18\x02 \x01(\t\x12\x0c\n\x04long\x18\x03 \x01(\t\x12\x0c\n\x04time\x18\x04 \x01(\t\x12,\n\x03\x64\x61y\x18\x05 \x01(\x0e\x32\x1f.viapb.Incident.DAY_OF_THE_WEEK\"m\n\x0f\x44\x41Y_OF_THE_WEEK\x12\n\n\x06SUNDAY\x10\x00\x12\n\n\x06MONDAY\x10\x01\x12\x0b\n\x07TUESDAY\x10\x02\x12\r\n\tWEDNESDAY\x10\x03\x12\x0c\n\x08THURSDAY\x10\x04\x12\n\n\x06\x46RIDAY\x10\x05\x12\x0c\n\x08SATURDAY\x10\x06\"=\n\x1bGetDangerProbabilityRequest\x12\x10\n\x08location\x18\x01 \x01(\t\x12\x0c\n\x04time\x18\x02 \x01(\t\".\n\x1cGetDangerProbabilityResponse\x12\x0e\n\x06\x64\x61nger\x18\x01 \x01(\t\"\x14\n\x12GetIncidentRequest\"8\n\x13GetIncidentResponse\x12!\n\x08incident\x18\x01 \x01(\x0b\x32\x0f.viapb.Incident2\xb0\x01\n\x03Via\x12\x61\n\x14GetDangerProbability\x12\".viapb.GetDangerProbabilityRequest\x1a#.viapb.GetDangerProbabilityResponse\"\x00\x12\x46\n\x0bGetIncident\x12\x19.viapb.GetIncidentRequest\x1a\x1a.viapb.GetIncidentResponse\"\x00\x42\nZ\x08./;viapbb\x06proto3')

_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, globals())
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'via_pb2', globals())
if _descriptor._USE_C_DESCRIPTORS == False:

  DESCRIPTOR._options = None
  DESCRIPTOR._serialized_options = b'Z\010./;viapb'
  _INCIDENT._serialized_start=21
  _INCIDENT._serialized_end=241
  _INCIDENT_DAY_OF_THE_WEEK._serialized_start=132
  _INCIDENT_DAY_OF_THE_WEEK._serialized_end=241
  _GETDANGERPROBABILITYREQUEST._serialized_start=243
  _GETDANGERPROBABILITYREQUEST._serialized_end=304
  _GETDANGERPROBABILITYRESPONSE._serialized_start=306
  _GETDANGERPROBABILITYRESPONSE._serialized_end=352
  _GETINCIDENTREQUEST._serialized_start=354
  _GETINCIDENTREQUEST._serialized_end=374
  _GETINCIDENTRESPONSE._serialized_start=376
  _GETINCIDENTRESPONSE._serialized_end=432
  _VIA._serialized_start=435
  _VIA._serialized_end=611
# @@protoc_insertion_point(module_scope)
