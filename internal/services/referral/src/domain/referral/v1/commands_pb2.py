# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: src/domain/referral/v1/commands.proto
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import field_mask_pb2 as google_dot_protobuf_dot_field__mask__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n%src/domain/referral/v1/commands.proto\x12\x16src.domain.referral.v1\x1a google/protobuf/field_mask.proto\"A\n\x12ReferralAddCommand\x12\x12\n\x04name\x18\x01 \x01(\tR\x04name\x12\x17\n\x07user_id\x18\x02 \x01(\tR\x06userId\"\x8f\x01\n\x15ReferralUpdateCommand\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id\x12\x12\n\x04name\x18\x02 \x01(\tR\x04name\x12\x17\n\x07user_id\x18\x03 \x01(\tR\x06userId\x12\x39\n\nfield_mask\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.FieldMaskR\tfieldMask\"\'\n\x15ReferralDeleteCommand\x12\x0e\n\x02id\x18\x01 \x01(\tR\x02id*\x87\x01\n\x0fReferralCommand\x12 \n\x1cREFERRAL_COMMAND_UNSPECIFIED\x10\x00\x12\x18\n\x14REFERRAL_COMMAND_ADD\x10\x01\x12\x1b\n\x17REFERRAL_COMMAND_UPDATE\x10\x02\x12\x1b\n\x17REFERRAL_COMMAND_DELETE\x10\x03\x42\xa6\x01\n\x1a\x63om.src.domain.referral.v1B\rCommandsProtoP\x01\xa2\x02\x03SDR\xaa\x02\x16Src.Domain.Referral.V1\xca\x02\x16Src\\Domain\\Referral\\V1\xe2\x02\"Src\\Domain\\Referral\\V1\\GPBMetadata\xea\x02\x19Src::Domain::Referral::V1b\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'src.domain.referral.v1.commands_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'\n\032com.src.domain.referral.v1B\rCommandsProtoP\001\242\002\003SDR\252\002\026Src.Domain.Referral.V1\312\002\026Src\\Domain\\Referral\\V1\342\002\"Src\\Domain\\Referral\\V1\\GPBMetadata\352\002\031Src::Domain::Referral::V1'
  _globals['_REFERRALCOMMAND']._serialized_start=354
  _globals['_REFERRALCOMMAND']._serialized_end=489
  _globals['_REFERRALADDCOMMAND']._serialized_start=99
  _globals['_REFERRALADDCOMMAND']._serialized_end=164
  _globals['_REFERRALUPDATECOMMAND']._serialized_start=167
  _globals['_REFERRALUPDATECOMMAND']._serialized_end=310
  _globals['_REFERRALDELETECOMMAND']._serialized_start=312
  _globals['_REFERRALDELETECOMMAND']._serialized_end=351
# @@protoc_insertion_point(module_scope)
