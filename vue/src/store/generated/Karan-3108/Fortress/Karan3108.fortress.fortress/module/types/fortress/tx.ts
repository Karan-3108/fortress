/* eslint-disable */
import { Reader, util, configure, Writer } from "protobufjs/minimal";
import * as Long from "long";

export const protobufPackage = "Karan3108.fortress.fortress";

export interface MsgRequestFortress {
  creator: string;
  amount: string;
  fee: string;
  collateral: string;
  deadline: string;
}

export interface MsgRequestFortressResponse {}

export interface MsgApproveFortress {
  creator: string;
  id: number;
}

export interface MsgApproveFortressResponse {}

export interface MsgRepayFortress {
  creator: string;
  id: number;
}

export interface MsgRepayFortressResponse {}

export interface MsgLiquidateFortress {
  creator: string;
  id: number;
}

export interface MsgLiquidateFortressResponse {}

export interface MsgCancelFortress {
  creator: string;
  id: number;
}

export interface MsgCancelFortressResponse {}

const baseMsgRequestFortress: object = {
  creator: "",
  amount: "",
  fee: "",
  collateral: "",
  deadline: "",
};

export const MsgRequestFortress = {
  encode(
    message: MsgRequestFortress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.amount !== "") {
      writer.uint32(18).string(message.amount);
    }
    if (message.fee !== "") {
      writer.uint32(26).string(message.fee);
    }
    if (message.collateral !== "") {
      writer.uint32(34).string(message.collateral);
    }
    if (message.deadline !== "") {
      writer.uint32(42).string(message.deadline);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRequestFortress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRequestFortress } as MsgRequestFortress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.amount = reader.string();
          break;
        case 3:
          message.fee = reader.string();
          break;
        case 4:
          message.collateral = reader.string();
          break;
        case 5:
          message.deadline = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRequestFortress {
    const message = { ...baseMsgRequestFortress } as MsgRequestFortress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = String(object.amount);
    } else {
      message.amount = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = String(object.fee);
    } else {
      message.fee = "";
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = String(object.collateral);
    } else {
      message.collateral = "";
    }
    if (object.deadline !== undefined && object.deadline !== null) {
      message.deadline = String(object.deadline);
    } else {
      message.deadline = "";
    }
    return message;
  },

  toJSON(message: MsgRequestFortress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.amount !== undefined && (obj.amount = message.amount);
    message.fee !== undefined && (obj.fee = message.fee);
    message.collateral !== undefined && (obj.collateral = message.collateral);
    message.deadline !== undefined && (obj.deadline = message.deadline);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRequestFortress>): MsgRequestFortress {
    const message = { ...baseMsgRequestFortress } as MsgRequestFortress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.amount !== undefined && object.amount !== null) {
      message.amount = object.amount;
    } else {
      message.amount = "";
    }
    if (object.fee !== undefined && object.fee !== null) {
      message.fee = object.fee;
    } else {
      message.fee = "";
    }
    if (object.collateral !== undefined && object.collateral !== null) {
      message.collateral = object.collateral;
    } else {
      message.collateral = "";
    }
    if (object.deadline !== undefined && object.deadline !== null) {
      message.deadline = object.deadline;
    } else {
      message.deadline = "";
    }
    return message;
  },
};

const baseMsgRequestFortressResponse: object = {};

export const MsgRequestFortressResponse = {
  encode(
    _: MsgRequestFortressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRequestFortressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRequestFortressResponse,
    } as MsgRequestFortressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRequestFortressResponse {
    const message = {
      ...baseMsgRequestFortressResponse,
    } as MsgRequestFortressResponse;
    return message;
  },

  toJSON(_: MsgRequestFortressResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRequestFortressResponse>
  ): MsgRequestFortressResponse {
    const message = {
      ...baseMsgRequestFortressResponse,
    } as MsgRequestFortressResponse;
    return message;
  },
};

const baseMsgApproveFortress: object = { creator: "", id: 0 };

export const MsgApproveFortress = {
  encode(
    message: MsgApproveFortress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgApproveFortress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgApproveFortress } as MsgApproveFortress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgApproveFortress {
    const message = { ...baseMsgApproveFortress } as MsgApproveFortress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgApproveFortress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgApproveFortress>): MsgApproveFortress {
    const message = { ...baseMsgApproveFortress } as MsgApproveFortress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgApproveFortressResponse: object = {};

export const MsgApproveFortressResponse = {
  encode(
    _: MsgApproveFortressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgApproveFortressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgApproveFortressResponse,
    } as MsgApproveFortressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgApproveFortressResponse {
    const message = {
      ...baseMsgApproveFortressResponse,
    } as MsgApproveFortressResponse;
    return message;
  },

  toJSON(_: MsgApproveFortressResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgApproveFortressResponse>
  ): MsgApproveFortressResponse {
    const message = {
      ...baseMsgApproveFortressResponse,
    } as MsgApproveFortressResponse;
    return message;
  },
};

const baseMsgRepayFortress: object = { creator: "", id: 0 };

export const MsgRepayFortress = {
  encode(message: MsgRepayFortress, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRepayFortress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRepayFortress } as MsgRepayFortress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRepayFortress {
    const message = { ...baseMsgRepayFortress } as MsgRepayFortress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgRepayFortress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRepayFortress>): MsgRepayFortress {
    const message = { ...baseMsgRepayFortress } as MsgRepayFortress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgRepayFortressResponse: object = {};

export const MsgRepayFortressResponse = {
  encode(
    _: MsgRepayFortressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgRepayFortressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgRepayFortressResponse,
    } as MsgRepayFortressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRepayFortressResponse {
    const message = {
      ...baseMsgRepayFortressResponse,
    } as MsgRepayFortressResponse;
    return message;
  },

  toJSON(_: MsgRepayFortressResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgRepayFortressResponse>
  ): MsgRepayFortressResponse {
    const message = {
      ...baseMsgRepayFortressResponse,
    } as MsgRepayFortressResponse;
    return message;
  },
};

const baseMsgLiquidateFortress: object = { creator: "", id: 0 };

export const MsgLiquidateFortress = {
  encode(
    message: MsgLiquidateFortress,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgLiquidateFortress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgLiquidateFortress } as MsgLiquidateFortress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgLiquidateFortress {
    const message = { ...baseMsgLiquidateFortress } as MsgLiquidateFortress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgLiquidateFortress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgLiquidateFortress>): MsgLiquidateFortress {
    const message = { ...baseMsgLiquidateFortress } as MsgLiquidateFortress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgLiquidateFortressResponse: object = {};

export const MsgLiquidateFortressResponse = {
  encode(
    _: MsgLiquidateFortressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgLiquidateFortressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgLiquidateFortressResponse,
    } as MsgLiquidateFortressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgLiquidateFortressResponse {
    const message = {
      ...baseMsgLiquidateFortressResponse,
    } as MsgLiquidateFortressResponse;
    return message;
  },

  toJSON(_: MsgLiquidateFortressResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgLiquidateFortressResponse>
  ): MsgLiquidateFortressResponse {
    const message = {
      ...baseMsgLiquidateFortressResponse,
    } as MsgLiquidateFortressResponse;
    return message;
  },
};

const baseMsgCancelFortress: object = { creator: "", id: 0 };

export const MsgCancelFortress = {
  encode(message: MsgCancelFortress, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.id !== 0) {
      writer.uint32(16).uint64(message.id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCancelFortress {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCancelFortress } as MsgCancelFortress;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.id = longToNumber(reader.uint64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCancelFortress {
    const message = { ...baseMsgCancelFortress } as MsgCancelFortress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = Number(object.id);
    } else {
      message.id = 0;
    }
    return message;
  },

  toJSON(message: MsgCancelFortress): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCancelFortress>): MsgCancelFortress {
    const message = { ...baseMsgCancelFortress } as MsgCancelFortress;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.id !== undefined && object.id !== null) {
      message.id = object.id;
    } else {
      message.id = 0;
    }
    return message;
  },
};

const baseMsgCancelFortressResponse: object = {};

export const MsgCancelFortressResponse = {
  encode(
    _: MsgCancelFortressResponse,
    writer: Writer = Writer.create()
  ): Writer {
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCancelFortressResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCancelFortressResponse,
    } as MsgCancelFortressResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgCancelFortressResponse {
    const message = {
      ...baseMsgCancelFortressResponse,
    } as MsgCancelFortressResponse;
    return message;
  },

  toJSON(_: MsgCancelFortressResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(
    _: DeepPartial<MsgCancelFortressResponse>
  ): MsgCancelFortressResponse {
    const message = {
      ...baseMsgCancelFortressResponse,
    } as MsgCancelFortressResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  RequestFortress(
    request: MsgRequestFortress
  ): Promise<MsgRequestFortressResponse>;
  ApproveFortress(
    request: MsgApproveFortress
  ): Promise<MsgApproveFortressResponse>;
  RepayFortress(request: MsgRepayFortress): Promise<MsgRepayFortressResponse>;
  LiquidateFortress(
    request: MsgLiquidateFortress
  ): Promise<MsgLiquidateFortressResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  CancelFortress(
    request: MsgCancelFortress
  ): Promise<MsgCancelFortressResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  RequestFortress(
    request: MsgRequestFortress
  ): Promise<MsgRequestFortressResponse> {
    const data = MsgRequestFortress.encode(request).finish();
    const promise = this.rpc.request(
      "Karan3108.fortress.fortress.Msg",
      "RequestFortress",
      data
    );
    return promise.then((data) =>
      MsgRequestFortressResponse.decode(new Reader(data))
    );
  }

  ApproveFortress(
    request: MsgApproveFortress
  ): Promise<MsgApproveFortressResponse> {
    const data = MsgApproveFortress.encode(request).finish();
    const promise = this.rpc.request(
      "Karan3108.fortress.fortress.Msg",
      "ApproveFortress",
      data
    );
    return promise.then((data) =>
      MsgApproveFortressResponse.decode(new Reader(data))
    );
  }

  RepayFortress(request: MsgRepayFortress): Promise<MsgRepayFortressResponse> {
    const data = MsgRepayFortress.encode(request).finish();
    const promise = this.rpc.request(
      "Karan3108.fortress.fortress.Msg",
      "RepayFortress",
      data
    );
    return promise.then((data) =>
      MsgRepayFortressResponse.decode(new Reader(data))
    );
  }

  LiquidateFortress(
    request: MsgLiquidateFortress
  ): Promise<MsgLiquidateFortressResponse> {
    const data = MsgLiquidateFortress.encode(request).finish();
    const promise = this.rpc.request(
      "Karan3108.fortress.fortress.Msg",
      "LiquidateFortress",
      data
    );
    return promise.then((data) =>
      MsgLiquidateFortressResponse.decode(new Reader(data))
    );
  }

  CancelFortress(
    request: MsgCancelFortress
  ): Promise<MsgCancelFortressResponse> {
    const data = MsgCancelFortress.encode(request).finish();
    const promise = this.rpc.request(
      "Karan3108.fortress.fortress.Msg",
      "CancelFortress",
      data
    );
    return promise.then((data) =>
      MsgCancelFortressResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") return globalThis;
  if (typeof self !== "undefined") return self;
  if (typeof window !== "undefined") return window;
  if (typeof global !== "undefined") return global;
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (util.Long !== Long) {
  util.Long = Long as any;
  configure();
}
