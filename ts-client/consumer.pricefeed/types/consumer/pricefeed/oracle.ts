/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";

export const protobufPackage = "consumer.pricefeed";

export interface RequestInterval {
  symbol: string;
  oracleScriptId: number;
  blockInterval: number;
}

function createBaseRequestInterval(): RequestInterval {
  return { symbol: "", oracleScriptId: 0, blockInterval: 0 };
}

export const RequestInterval = {
  encode(message: RequestInterval, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    if (message.oracleScriptId !== 0) {
      writer.uint32(16).uint64(message.oracleScriptId);
    }
    if (message.blockInterval !== 0) {
      writer.uint32(24).uint64(message.blockInterval);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): RequestInterval {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseRequestInterval();
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
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): RequestInterval {
    return {
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      oracleScriptId: isSet(object.oracleScriptId) ? Number(object.oracleScriptId) : 0,
      blockInterval: isSet(object.blockInterval) ? Number(object.blockInterval) : 0,
    };
  },

  toJSON(message: RequestInterval): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.oracleScriptId !== undefined && (obj.oracleScriptId = Math.round(message.oracleScriptId));
    message.blockInterval !== undefined && (obj.blockInterval = Math.round(message.blockInterval));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<RequestInterval>, I>>(object: I): RequestInterval {
    const message = createBaseRequestInterval();
    message.symbol = object.symbol ?? "";
    message.oracleScriptId = object.oracleScriptId ?? 0;
    message.blockInterval = object.blockInterval ?? 0;
    return message;
  },
};

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
