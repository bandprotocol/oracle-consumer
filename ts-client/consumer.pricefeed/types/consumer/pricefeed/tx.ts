/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "consumer.pricefeed";

export interface MsgCreateRequestInterval {
  symbol: string;
  oracleScriptId: number;
  blockInterval: number;
  sender: string;
}

export interface MsgCreateRequestIntervalResponse {
}

function createBaseMsgCreateRequestInterval(): MsgCreateRequestInterval {
  return { symbol: "", oracleScriptId: 0, blockInterval: 0, sender: "" };
}

export const MsgCreateRequestInterval = {
  encode(message: MsgCreateRequestInterval, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    if (message.oracleScriptId !== 0) {
      writer.uint32(16).uint64(message.oracleScriptId);
    }
    if (message.blockInterval !== 0) {
      writer.uint32(24).uint64(message.blockInterval);
    }
    if (message.sender !== "") {
      writer.uint32(34).string(message.sender);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateRequestInterval {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateRequestInterval();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbol = reader.string();
          break;
        case 2:
          message.oracleScriptId = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.blockInterval = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.sender = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateRequestInterval {
    return {
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      oracleScriptId: isSet(object.oracleScriptId) ? Number(object.oracleScriptId) : 0,
      blockInterval: isSet(object.blockInterval) ? Number(object.blockInterval) : 0,
      sender: isSet(object.sender) ? String(object.sender) : "",
    };
  },

  toJSON(message: MsgCreateRequestInterval): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.oracleScriptId !== undefined && (obj.oracleScriptId = Math.round(message.oracleScriptId));
    message.blockInterval !== undefined && (obj.blockInterval = Math.round(message.blockInterval));
    message.sender !== undefined && (obj.sender = message.sender);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateRequestInterval>, I>>(object: I): MsgCreateRequestInterval {
    const message = createBaseMsgCreateRequestInterval();
    message.symbol = object.symbol ?? "";
    message.oracleScriptId = object.oracleScriptId ?? 0;
    message.blockInterval = object.blockInterval ?? 0;
    message.sender = object.sender ?? "";
    return message;
  },
};

function createBaseMsgCreateRequestIntervalResponse(): MsgCreateRequestIntervalResponse {
  return {};
}

export const MsgCreateRequestIntervalResponse = {
  encode(_: MsgCreateRequestIntervalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgCreateRequestIntervalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgCreateRequestIntervalResponse();
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

  fromJSON(_: any): MsgCreateRequestIntervalResponse {
    return {};
  },

  toJSON(_: MsgCreateRequestIntervalResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgCreateRequestIntervalResponse>, I>>(
    _: I,
  ): MsgCreateRequestIntervalResponse {
    const message = createBaseMsgCreateRequestIntervalResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  CreateRequestInterval(request: MsgCreateRequestInterval): Promise<MsgCreateRequestIntervalResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.CreateRequestInterval = this.CreateRequestInterval.bind(this);
  }
  CreateRequestInterval(request: MsgCreateRequestInterval): Promise<MsgCreateRequestIntervalResponse> {
    const data = MsgCreateRequestInterval.encode(request).finish();
    const promise = this.rpc.request("consumer.pricefeed.Msg", "CreateRequestInterval", data);
    return promise.then((data) => MsgCreateRequestIntervalResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

declare var self: any | undefined;
declare var window: any | undefined;
declare var global: any | undefined;
var globalThis: any = (() => {
  if (typeof globalThis !== "undefined") {
    return globalThis;
  }
  if (typeof self !== "undefined") {
    return self;
  }
  if (typeof window !== "undefined") {
    return window;
  }
  if (typeof global !== "undefined") {
    return global;
  }
  throw "Unable to locate global object";
})();

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error("Value is larger than Number.MAX_SAFE_INTEGER");
  }
  return long.toNumber();
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any;
  _m0.configure();
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
