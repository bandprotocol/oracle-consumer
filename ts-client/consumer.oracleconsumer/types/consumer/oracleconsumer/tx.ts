/* eslint-disable */
import Long from 'long'
import _m0 from 'protobufjs/minimal'
import { Coin } from '../../cosmos/base/v1beta1/coin'

export const protobufPackage = 'consumer.pricefeed'

export interface MsgSetRequestInterval {
  osId: number
  calldata: Uint8Array
  blockInterval: number
  askCount: number
  minCount: number
  feeLimit: Coin[]
  prepareGas: number
  executeGas: number
  sender: string
}

export interface MsgSetRequestIntervalResponse {}

function createBaseMsgSetRequestInterval(): MsgSetRequestInterval {
  return {
    osId: 0,
    calldata: new Uint8Array(),
    blockInterval: 0,
    askCount: 0,
    minCount: 0,
    feeLimit: [],
    prepareGas: 0,
    executeGas: 0,
    sender: '',
  }
}

export const MsgSetRequestInterval = {
  encode(message: MsgSetRequestInterval, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.osId !== 0) {
      writer.uint32(8).uint64(message.osId)
    }
    if (message.calldata.length !== 0) {
      writer.uint32(18).bytes(message.calldata)
    }
    if (message.blockInterval !== 0) {
      writer.uint32(24).uint64(message.blockInterval)
    }
    if (message.askCount !== 0) {
      writer.uint32(32).uint64(message.askCount)
    }
    if (message.minCount !== 0) {
      writer.uint32(40).uint64(message.minCount)
    }
    for (const v of message.feeLimit) {
      Coin.encode(v!, writer.uint32(50).fork()).ldelim()
    }
    if (message.prepareGas !== 0) {
      writer.uint32(56).uint64(message.prepareGas)
    }
    if (message.executeGas !== 0) {
      writer.uint32(64).uint64(message.executeGas)
    }
    if (message.sender !== '') {
      writer.uint32(74).string(message.sender)
    }
    return writer
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetRequestInterval {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
    let end = length === undefined ? reader.len : reader.pos + length
    const message = createBaseMsgSetRequestInterval()
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        case 1:
          message.osId = longToNumber(reader.uint64() as Long)
          break
        case 2:
          message.calldata = reader.bytes()
          break
        case 3:
          message.blockInterval = longToNumber(reader.uint64() as Long)
          break
        case 4:
          message.askCount = longToNumber(reader.uint64() as Long)
          break
        case 5:
          message.minCount = longToNumber(reader.uint64() as Long)
          break
        case 6:
          message.feeLimit.push(Coin.decode(reader, reader.uint32()))
          break
        case 7:
          message.prepareGas = longToNumber(reader.uint64() as Long)
          break
        case 8:
          message.executeGas = longToNumber(reader.uint64() as Long)
          break
        case 9:
          message.sender = reader.string()
          break
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(object: any): MsgSetRequestInterval {
    return {
      osId: isSet(object.osId) ? Number(object.osId) : 0,
      calldata: isSet(object.calldata) ? bytesFromBase64(object.calldata) : new Uint8Array(),
      blockInterval: isSet(object.blockInterval) ? Number(object.blockInterval) : 0,
      askCount: isSet(object.askCount) ? Number(object.askCount) : 0,
      minCount: isSet(object.minCount) ? Number(object.minCount) : 0,
      feeLimit: Array.isArray(object?.feeLimit) ? object.feeLimit.map((e: any) => Coin.fromJSON(e)) : [],
      prepareGas: isSet(object.prepareGas) ? Number(object.prepareGas) : 0,
      executeGas: isSet(object.executeGas) ? Number(object.executeGas) : 0,
      sender: isSet(object.sender) ? String(object.sender) : '',
    }
  },

  toJSON(message: MsgSetRequestInterval): unknown {
    const obj: any = {}
    message.osId !== undefined && (obj.osId = Math.round(message.osId))
    message.calldata !== undefined &&
      (obj.calldata = base64FromBytes(message.calldata !== undefined ? message.calldata : new Uint8Array()))
    message.blockInterval !== undefined && (obj.blockInterval = Math.round(message.blockInterval))
    message.askCount !== undefined && (obj.askCount = Math.round(message.askCount))
    message.minCount !== undefined && (obj.minCount = Math.round(message.minCount))
    if (message.feeLimit) {
      obj.feeLimit = message.feeLimit.map((e) => (e ? Coin.toJSON(e) : undefined))
    } else {
      obj.feeLimit = []
    }
    message.prepareGas !== undefined && (obj.prepareGas = Math.round(message.prepareGas))
    message.executeGas !== undefined && (obj.executeGas = Math.round(message.executeGas))
    message.sender !== undefined && (obj.sender = message.sender)
    return obj
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetRequestInterval>, I>>(object: I): MsgSetRequestInterval {
    const message = createBaseMsgSetRequestInterval()
    message.osId = object.osId ?? 0
    message.calldata = object.calldata ?? new Uint8Array()
    message.blockInterval = object.blockInterval ?? 0
    message.askCount = object.askCount ?? 0
    message.minCount = object.minCount ?? 0
    message.feeLimit = object.feeLimit?.map((e) => Coin.fromPartial(e)) || []
    message.prepareGas = object.prepareGas ?? 0
    message.executeGas = object.executeGas ?? 0
    message.sender = object.sender ?? ''
    return message
  },
}

function createBaseMsgSetRequestIntervalResponse(): MsgSetRequestIntervalResponse {
  return {}
}

export const MsgSetRequestIntervalResponse = {
  encode(_: MsgSetRequestIntervalResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSetRequestIntervalResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input)
    let end = length === undefined ? reader.len : reader.pos + length
    const message = createBaseMsgSetRequestIntervalResponse()
    while (reader.pos < end) {
      const tag = reader.uint32()
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7)
          break
      }
    }
    return message
  },

  fromJSON(_: any): MsgSetRequestIntervalResponse {
    return {}
  },

  toJSON(_: MsgSetRequestIntervalResponse): unknown {
    const obj: any = {}
    return obj
  },

  fromPartial<I extends Exact<DeepPartial<MsgSetRequestIntervalResponse>, I>>(_: I): MsgSetRequestIntervalResponse {
    const message = createBaseMsgSetRequestIntervalResponse()
    return message
  },
}

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SetRequestInterval(request: MsgSetRequestInterval): Promise<MsgSetRequestIntervalResponse>
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc
  constructor(rpc: Rpc) {
    this.rpc = rpc
    this.SetRequestInterval = this.SetRequestInterval.bind(this)
  }
  SetRequestInterval(request: MsgSetRequestInterval): Promise<MsgSetRequestIntervalResponse> {
    const data = MsgSetRequestInterval.encode(request).finish()
    const promise = this.rpc.request('consumer.pricefeed.Msg', 'SetRequestInterval', data)
    return promise.then((data) => MsgSetRequestIntervalResponse.decode(new _m0.Reader(data)))
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>
}

declare var self: any | undefined
declare var window: any | undefined
declare var global: any | undefined
var globalThis: any = (() => {
  if (typeof globalThis !== 'undefined') {
    return globalThis
  }
  if (typeof self !== 'undefined') {
    return self
  }
  if (typeof window !== 'undefined') {
    return window
  }
  if (typeof global !== 'undefined') {
    return global
  }
  throw 'Unable to locate global object'
})()

function bytesFromBase64(b64: string): Uint8Array {
  if (globalThis.Buffer) {
    return Uint8Array.from(globalThis.Buffer.from(b64, 'base64'))
  } else {
    const bin = globalThis.atob(b64)
    const arr = new Uint8Array(bin.length)
    for (let i = 0; i < bin.length; ++i) {
      arr[i] = bin.charCodeAt(i)
    }
    return arr
  }
}

function base64FromBytes(arr: Uint8Array): string {
  if (globalThis.Buffer) {
    return globalThis.Buffer.from(arr).toString('base64')
  } else {
    const bin: string[] = []
    arr.forEach((byte) => {
      bin.push(String.fromCharCode(byte))
    })
    return globalThis.btoa(bin.join(''))
  }
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined

export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>

type KeysOfUnion<T> = T extends T ? keyof T : never
export type Exact<P, I extends P> = P extends Builtin
  ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never }

function longToNumber(long: Long): number {
  if (long.gt(Number.MAX_SAFE_INTEGER)) {
    throw new globalThis.Error('Value is larger than Number.MAX_SAFE_INTEGER')
  }
  return long.toNumber()
}

if (_m0.util.Long !== Long) {
  _m0.util.Long = Long as any
  _m0.configure()
}

function isSet(value: any): boolean {
  return value !== null && value !== undefined
}
