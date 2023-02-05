/* eslint-disable */
import Long from "long";
import _m0 from "protobufjs/minimal";
import { Symbols } from "./oracle";
import { Params } from "./params";

export const protobufPackage = "consumer.pricefeed";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {
}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QuerySymbols {
}

export interface QuerySymbolsResponse {
  symbols: Symbols | undefined;
}

export interface QueryPrice {
  symbol: string;
}

export interface QueryPriceResponse {
  symbol: string;
  price: number;
  resolveTime: number;
}

function createBaseQueryParamsRequest(): QueryParamsRequest {
  return {};
}

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsRequest();
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

  fromJSON(_: any): QueryParamsRequest {
    return {};
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsRequest>, I>>(_: I): QueryParamsRequest {
    const message = createBaseQueryParamsRequest();
    return message;
  },
};

function createBaseQueryParamsResponse(): QueryParamsResponse {
  return { params: undefined };
}

export const QueryParamsResponse = {
  encode(message: QueryParamsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryParamsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    return { params: isSet(object.params) ? Params.fromJSON(object.params) : undefined };
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined && (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryParamsResponse>, I>>(object: I): QueryParamsResponse {
    const message = createBaseQueryParamsResponse();
    message.params = (object.params !== undefined && object.params !== null)
      ? Params.fromPartial(object.params)
      : undefined;
    return message;
  },
};

function createBaseQuerySymbols(): QuerySymbols {
  return {};
}

export const QuerySymbols = {
  encode(_: QuerySymbols, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerySymbols {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerySymbols();
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

  fromJSON(_: any): QuerySymbols {
    return {};
  },

  toJSON(_: QuerySymbols): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QuerySymbols>, I>>(_: I): QuerySymbols {
    const message = createBaseQuerySymbols();
    return message;
  },
};

function createBaseQuerySymbolsResponse(): QuerySymbolsResponse {
  return { symbols: undefined };
}

export const QuerySymbolsResponse = {
  encode(message: QuerySymbolsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.symbols !== undefined) {
      Symbols.encode(message.symbols, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerySymbolsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerySymbolsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbols = Symbols.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuerySymbolsResponse {
    return { symbols: isSet(object.symbols) ? Symbols.fromJSON(object.symbols) : undefined };
  },

  toJSON(message: QuerySymbolsResponse): unknown {
    const obj: any = {};
    message.symbols !== undefined && (obj.symbols = message.symbols ? Symbols.toJSON(message.symbols) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QuerySymbolsResponse>, I>>(object: I): QuerySymbolsResponse {
    const message = createBaseQuerySymbolsResponse();
    message.symbols = (object.symbols !== undefined && object.symbols !== null)
      ? Symbols.fromPartial(object.symbols)
      : undefined;
    return message;
  },
};

function createBaseQueryPrice(): QueryPrice {
  return { symbol: "" };
}

export const QueryPrice = {
  encode(message: QueryPrice, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryPrice {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryPrice();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbol = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPrice {
    return { symbol: isSet(object.symbol) ? String(object.symbol) : "" };
  },

  toJSON(message: QueryPrice): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryPrice>, I>>(object: I): QueryPrice {
    const message = createBaseQueryPrice();
    message.symbol = object.symbol ?? "";
    return message;
  },
};

function createBaseQueryPriceResponse(): QueryPriceResponse {
  return { symbol: "", price: 0, resolveTime: 0 };
}

export const QueryPriceResponse = {
  encode(message: QueryPriceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    if (message.price !== 0) {
      writer.uint32(16).uint64(message.price);
    }
    if (message.resolveTime !== 0) {
      writer.uint32(24).int64(message.resolveTime);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QueryPriceResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQueryPriceResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbol = reader.string();
          break;
        case 2:
          message.price = longToNumber(reader.uint64() as Long);
          break;
        case 3:
          message.resolveTime = longToNumber(reader.int64() as Long);
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPriceResponse {
    return {
      symbol: isSet(object.symbol) ? String(object.symbol) : "",
      price: isSet(object.price) ? Number(object.price) : 0,
      resolveTime: isSet(object.resolveTime) ? Number(object.resolveTime) : 0,
    };
  },

  toJSON(message: QueryPriceResponse): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    message.price !== undefined && (obj.price = Math.round(message.price));
    message.resolveTime !== undefined && (obj.resolveTime = Math.round(message.resolveTime));
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryPriceResponse>, I>>(object: I): QueryPriceResponse {
    const message = createBaseQueryPriceResponse();
    message.symbol = object.symbol ?? "";
    message.price = object.price ?? 0;
    message.resolveTime = object.resolveTime ?? 0;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  Symbols(request: QuerySymbols): Promise<QuerySymbolsResponse>;
  Price(request: QueryPrice): Promise<QueryPriceResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.Symbols = this.Symbols.bind(this);
    this.Price = this.Price.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("consumer.pricefeed.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  Symbols(request: QuerySymbols): Promise<QuerySymbolsResponse> {
    const data = QuerySymbols.encode(request).finish();
    const promise = this.rpc.request("consumer.pricefeed.Query", "Symbols", data);
    return promise.then((data) => QuerySymbolsResponse.decode(new _m0.Reader(data)));
  }

  Price(request: QueryPrice): Promise<QueryPriceResponse> {
    const data = QueryPrice.encode(request).finish();
    const promise = this.rpc.request("consumer.pricefeed.Query", "Price", data);
    return promise.then((data) => QueryPriceResponse.decode(new _m0.Reader(data)));
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
