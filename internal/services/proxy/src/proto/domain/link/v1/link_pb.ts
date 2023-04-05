// @generated by protoc-gen-es v1.2.0 with parameter "target=ts"
// @generated from file domain/link/v1/link.proto (package domain.link.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { FieldMask, Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum domain.link.v1.LinkEvent
 */
export enum LinkEvent {
  /**
   * @generated from enum value: LINK_EVENT_UNSPECIFIED = 0;
   */
  UNSPECIFIED = 0,

  /**
   * @generated from enum value: LINK_EVENT_ADD = 1;
   */
  ADD = 1,

  /**
   * @generated from enum value: LINK_EVENT_GET = 2;
   */
  GET = 2,

  /**
   * @generated from enum value: LINK_EVENT_LIST = 3;
   */
  LIST = 3,

  /**
   * @generated from enum value: LINK_EVENT_UPDATE = 4;
   */
  UPDATE = 4,

  /**
   * @generated from enum value: LINK_EVENT_DELETE = 5;
   */
  DELETE = 5,
}
// Retrieve enum metadata with: proto3.getEnumType(LinkEvent)
proto3.util.setEnumType(LinkEvent, "domain.link.v1.LinkEvent", [
  { no: 0, name: "LINK_EVENT_UNSPECIFIED" },
  { no: 1, name: "LINK_EVENT_ADD" },
  { no: 2, name: "LINK_EVENT_GET" },
  { no: 3, name: "LINK_EVENT_LIST" },
  { no: 4, name: "LINK_EVENT_UPDATE" },
  { no: 5, name: "LINK_EVENT_DELETE" },
]);

/**
 * @generated from message domain.link.v1.Link
 */
export class Link extends Message<Link> {
  /**
   * @generated from field: google.protobuf.FieldMask field_mask = 6;
   */
  fieldMask?: FieldMask;

  /**
   * URL
   *
   * @generated from field: string url = 1;
   */
  url = "";

  /**
   * Hash by URL + salt
   *
   * @generated from field: string hash = 2;
   */
  hash = "";

  /**
   * Describe of link
   *
   * @generated from field: string describe = 3;
   */
  describe = "";

  /**
   * Create at
   *
   * @generated from field: google.protobuf.Timestamp created_at = 4;
   */
  createdAt?: Timestamp;

  /**
   * Update at
   *
   * @generated from field: google.protobuf.Timestamp updated_at = 5;
   */
  updatedAt?: Timestamp;

  constructor(data?: PartialMessage<Link>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "domain.link.v1.Link";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 6, name: "field_mask", kind: "message", T: FieldMask },
    { no: 1, name: "url", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "hash", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "describe", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "created_at", kind: "message", T: Timestamp },
    { no: 5, name: "updated_at", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Link {
    return new Link().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Link {
    return new Link().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Link {
    return new Link().fromJsonString(jsonString, options);
  }

  static equals(a: Link | PlainMessage<Link> | undefined, b: Link | PlainMessage<Link> | undefined): boolean {
    return proto3.util.equals(Link, a, b);
  }
}

/**
 * @generated from message domain.link.v1.Links
 */
export class Links extends Message<Links> {
  /**
   * @generated from field: repeated domain.link.v1.Link link = 1;
   */
  link: Link[] = [];

  constructor(data?: PartialMessage<Links>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "domain.link.v1.Links";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "link", kind: "message", T: Link, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Links {
    return new Links().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Links {
    return new Links().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Links {
    return new Links().fromJsonString(jsonString, options);
  }

  static equals(a: Links | PlainMessage<Links> | undefined, b: Links | PlainMessage<Links> | undefined): boolean {
    return proto3.util.equals(Links, a, b);
  }
}

