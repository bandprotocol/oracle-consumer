/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Coin } from "../../cosmos/base/v1beta1/coin";

export const protobufPackage = "consumer.pricefeed";

/** Params defines the parameters for the module. */
export interface Params {
  askCount: number;
  minCount: number;
  minDsCount: number;
  prepareGasA: number;
  prepareGasB: number;
  executeGasA: number;
  executeGasB: number;
  sourceChannel: number;
  feeLimit: Coin[];
}

function createBaseParams(): Params {
  return {
    askCount: 0,
    minCount: 0,
    minDsCount: 0,
    prepareGasA: 0,
    prepareGasB: 0,
    executeGasA: 0,
    executeGasB: 0,
    sourceChannel: 0,
    feeLimit: [],
  };
}

export const Params = {
  encode(message: Params, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.askCount !== 0) {
      writer.uint32(8).uint64(message.askCount);
    }
    if (message.minCount !== 0) {
      writer.uint32(16).uint64(message.minCount);
    }
    if (message.minDsCount !== 0) {
      writer.uint32(24).uint64(message.minDsCount);
    }
    if (message.prepareGasA !== 0) {
      writer.uint32(32).uint64(message.prepareGasA);
    }
    if (message.prepareGasB !== 0) {
      writer.uint32(40).uint64(message.prepareGasB);
    }
    if (message.executeGasA !== 0) {
      writer.uint32(48).uint64(message.executeGasA);
    }
    if (message.executeGasB !== 0) {
      writer.uint32(56).uint64(message.executeGasB);
    }
    if (message.sourceChannel !== 0) {
      writer.uint32(64).uint64(message.sourceChannel);
    }
    for (const v of message.feeLimit) {
      Coin.encode(v!, writer.uint32(74).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Params {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseParams();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.askCount = longToNumber(reader.uint64() as Long);
          break;
        case 2:
          message.minCount = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.minDsCount = longToNumber(reader.uint64() as Long);
          break;
        case 4:
          message.prepareGasA = longToNumber(reader.uint64() as Long);
          break;
        case 5:
          message.prepareGasB = longToNumber(reader.uint64() as Long);
          break;
        case 6:
          message.executeGasA = longToNumber(reader.uint64() as Long);
          break;
        case 7:
          message.executeGasB = longToNumber(reader.uint64() as Long);
          break;
        case 8:
          message.sourceChannel = longToNumber(reader.uint64() as Long);
          break;
        case 9:
          message.feeLimit.push(Coin.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Params {
    return {
      askCount: isSet(object.askCount) ? Number(object.askCount) : 0,
      minCount: isSet(object.minCount) ? Number(object.minCount) : 0,
      minDsCount: isSet(object.minDsCount) ? Number(object.minDsCount) : 0,
      prepareGasA: isSet(object.prepareGasA) ? Number(object.prepareGasA) : 0,
      prepareGasB: isSet(object.prepareGasB) ? Number(object.prepareGasB) : 0,
      executeGasA: isSet(object.executeGasA) ? Number(object.executeGasA) : 0,
      executeGasB: isSet(object.executeGasB) ? Number(object.executeGasB) : 0,
      sourceChannel: isSet(object.sourceChannel) ? Number(object.sourceChannel) : 0,
      feeLimit: Array.isArray(object?.feeLimit) ? object.feeLimit.map((e: any) => Coin.fromJSON(e)) : [],
    };
  },

  toJSON(message: Params): unknown {
    const obj: any = {};
    message.askCount !== undefined && (obj.askCount = Math.round(message.askCount));
    message.minCount !== undefined && (obj.minCount = Math.round(message.minCount));
    message.minDsCount !== undefined && (obj.minDsCount = Math.round(message.minDsCount));
    message.prepareGasA !== undefined && (obj.prepareGasA = Math.round(message.prepareGasA));
    message.prepareGasB !== undefined && (obj.prepareGasB = Math.round(message.prepareGasB));
    message.executeGasA !== undefined && (obj.executeGasA = Math.round(message.executeGasA));
    message.executeGasB !== undefined && (obj.executeGasB = Math.round(message.executeGasB));
    message.sourceChannel !== undefined && (obj.sourceChannel = Math.round(message.sourceChannel));
    if (message.feeLimit) {
      obj.feeLimit = message.feeLimit.map((e) => e ? Coin.toJSON(e) : undefined);
    } else {
      obj.feeLimit = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Params>, I>>(object: I): Params {
    const message = createBaseParams();
    message.askCount = object.askCount ?? 0;
    message.minCount = object.minCount ?? 0;
    message.minDsCount = object.minDsCount ?? 0;
    message.prepareGasA = object.prepareGasA ?? 0;
    message.prepareGasB = object.prepareGasB ?? 0;
    message.executeGasA = object.executeGasA ?? 0;
    message.executeGasB = object.executeGasB ?? 0;
    message.sourceChannel = object.sourceChannel ?? 0;
    message.feeLimit = object.feeLimit?.map((e) => Coin.fromPartial(e)) || [];
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
