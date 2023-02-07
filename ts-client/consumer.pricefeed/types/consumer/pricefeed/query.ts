/* eslint-disable */
import _m0 from "protobufjs/minimal";
import { Price, SymbolRequest } from "./oracle";
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

export interface QuerySymbolRequest {
  symbol: string;
}

export interface QuerySymbolRequestResponse {
  symbolRequest: SymbolRequest | undefined;
}

export interface QuerySymbolRequests {
}

export interface QuerySymbolRequestsResponse {
  symbolRequests: SymbolRequest[];
}

export interface QueryPrice {
  symbol: string;
}

export interface QueryPriceResponse {
  price: Price | undefined;
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

function createBaseQuerySymbolRequest(): QuerySymbolRequest {
  return { symbol: "" };
}

export const QuerySymbolRequest = {
  encode(message: QuerySymbolRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.symbol !== "") {
      writer.uint32(10).string(message.symbol);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerySymbolRequest {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerySymbolRequest();
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

  fromJSON(object: any): QuerySymbolRequest {
    return { symbol: isSet(object.symbol) ? String(object.symbol) : "" };
  },

  toJSON(message: QuerySymbolRequest): unknown {
    const obj: any = {};
    message.symbol !== undefined && (obj.symbol = message.symbol);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QuerySymbolRequest>, I>>(object: I): QuerySymbolRequest {
    const message = createBaseQuerySymbolRequest();
    message.symbol = object.symbol ?? "";
    return message;
  },
};

function createBaseQuerySymbolRequestResponse(): QuerySymbolRequestResponse {
  return { symbolRequest: undefined };
}

export const QuerySymbolRequestResponse = {
  encode(message: QuerySymbolRequestResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.symbolRequest !== undefined) {
      SymbolRequest.encode(message.symbolRequest, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerySymbolRequestResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerySymbolRequestResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbolRequest = SymbolRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuerySymbolRequestResponse {
    return { symbolRequest: isSet(object.symbolRequest) ? SymbolRequest.fromJSON(object.symbolRequest) : undefined };
  },

  toJSON(message: QuerySymbolRequestResponse): unknown {
    const obj: any = {};
    message.symbolRequest !== undefined
      && (obj.symbolRequest = message.symbolRequest ? SymbolRequest.toJSON(message.symbolRequest) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QuerySymbolRequestResponse>, I>>(object: I): QuerySymbolRequestResponse {
    const message = createBaseQuerySymbolRequestResponse();
    message.symbolRequest = (object.symbolRequest !== undefined && object.symbolRequest !== null)
      ? SymbolRequest.fromPartial(object.symbolRequest)
      : undefined;
    return message;
  },
};

function createBaseQuerySymbolRequests(): QuerySymbolRequests {
  return {};
}

export const QuerySymbolRequests = {
  encode(_: QuerySymbolRequests, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerySymbolRequests {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerySymbolRequests();
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

  fromJSON(_: any): QuerySymbolRequests {
    return {};
  },

  toJSON(_: QuerySymbolRequests): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QuerySymbolRequests>, I>>(_: I): QuerySymbolRequests {
    const message = createBaseQuerySymbolRequests();
    return message;
  },
};

function createBaseQuerySymbolRequestsResponse(): QuerySymbolRequestsResponse {
  return { symbolRequests: [] };
}

export const QuerySymbolRequestsResponse = {
  encode(message: QuerySymbolRequestsResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    for (const v of message.symbolRequests) {
      SymbolRequest.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): QuerySymbolRequestsResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseQuerySymbolRequestsResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.symbolRequests.push(SymbolRequest.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QuerySymbolRequestsResponse {
    return {
      symbolRequests: Array.isArray(object?.symbolRequests)
        ? object.symbolRequests.map((e: any) => SymbolRequest.fromJSON(e))
        : [],
    };
  },

  toJSON(message: QuerySymbolRequestsResponse): unknown {
    const obj: any = {};
    if (message.symbolRequests) {
      obj.symbolRequests = message.symbolRequests.map((e) => e ? SymbolRequest.toJSON(e) : undefined);
    } else {
      obj.symbolRequests = [];
    }
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QuerySymbolRequestsResponse>, I>>(object: I): QuerySymbolRequestsResponse {
    const message = createBaseQuerySymbolRequestsResponse();
    message.symbolRequests = object.symbolRequests?.map((e) => SymbolRequest.fromPartial(e)) || [];
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
  return { price: undefined };
}

export const QueryPriceResponse = {
  encode(message: QueryPriceResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.price !== undefined) {
      Price.encode(message.price, writer.uint32(10).fork()).ldelim();
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
          message.price = Price.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryPriceResponse {
    return { price: isSet(object.price) ? Price.fromJSON(object.price) : undefined };
  },

  toJSON(message: QueryPriceResponse): unknown {
    const obj: any = {};
    message.price !== undefined && (obj.price = message.price ? Price.toJSON(message.price) : undefined);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<QueryPriceResponse>, I>>(object: I): QueryPriceResponse {
    const message = createBaseQueryPriceResponse();
    message.price = (object.price !== undefined && object.price !== null) ? Price.fromPartial(object.price) : undefined;
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  SymbolRequest(request: QuerySymbolRequest): Promise<QuerySymbolRequestResponse>;
  SymbolRequests(request: QuerySymbolRequests): Promise<QuerySymbolRequestsResponse>;
  Price(request: QueryPrice): Promise<QueryPriceResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.Params = this.Params.bind(this);
    this.SymbolRequest = this.SymbolRequest.bind(this);
    this.SymbolRequests = this.SymbolRequests.bind(this);
    this.Price = this.Price.bind(this);
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("consumer.pricefeed.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new _m0.Reader(data)));
  }

  SymbolRequest(request: QuerySymbolRequest): Promise<QuerySymbolRequestResponse> {
    const data = QuerySymbolRequest.encode(request).finish();
    const promise = this.rpc.request("consumer.pricefeed.Query", "SymbolRequest", data);
    return promise.then((data) => QuerySymbolRequestResponse.decode(new _m0.Reader(data)));
  }

  SymbolRequests(request: QuerySymbolRequests): Promise<QuerySymbolRequestsResponse> {
    const data = QuerySymbolRequests.encode(request).finish();
    const promise = this.rpc.request("consumer.pricefeed.Query", "SymbolRequests", data);
    return promise.then((data) => QuerySymbolRequestsResponse.decode(new _m0.Reader(data)));
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

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
