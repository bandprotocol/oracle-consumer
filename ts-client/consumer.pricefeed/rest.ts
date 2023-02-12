/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

/**
 * Params defines the parameters for the module.
 */
export interface PricefeedParams {
  /** @format uint64 */
  ask_count?: string;

  /** @format uint64 */
  min_count?: string;

  /** @format uint64 */
  min_ds_count?: string;

  /** @format uint64 */
  prepare_gas_base?: string;

  /** @format uint64 */
  prepare_gas_each?: string;

  /** @format uint64 */
  execute_gas_base?: string;

  /** @format uint64 */
  execute_gas_each?: string;
  source_channel?: string;
  fee_limit?: V1Beta1Coin[];
}

export interface PricefeedPrice {
  symbol?: string;

  /** @format uint64 */
  price?: string;

  /** @format int64 */
  resolve_time?: string;
}

/**
 * QueryParamsResponse is response type for the Query/Params RPC method.
 */
export interface PricefeedQueryParamsResponse {
  /** params holds all the parameters of this module. */
  params?: PricefeedParams;
}

export interface PricefeedQueryPriceResponse {
  price?: PricefeedPrice;
}

export interface PricefeedQuerySymbolRequestResponse {
  symbol_request?: PricefeedSymbolRequest;
}

export interface PricefeedQuerySymbolRequestsResponse {
  symbol_requests?: PricefeedSymbolRequest[];
}

export interface PricefeedSymbolRequest {
  symbol?: string;

  /** @format uint64 */
  oracle_script_id?: string;

  /** @format uint64 */
  block_interval?: string;
}

export interface ProtobufAny {
  "@type"?: string;
}

export interface RpcStatus {
  /** @format int32 */
  code?: number;
  message?: string;
  details?: ProtobufAny[];
}

/**
* Coin defines a token with a denomination and an amount.

NOTE: The amount field is an Int which implements the custom method
signatures required by gogoproto.
*/
export interface V1Beta1Coin {
  denom?: string;
  amount?: string;
}

import axios, { AxiosInstance, AxiosRequestConfig, AxiosResponse, ResponseType } from "axios";

export type QueryParamsType = Record<string | number, any>;

export interface FullRequestParams extends Omit<AxiosRequestConfig, "data" | "params" | "url" | "responseType"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseType;
  /** request body */
  body?: unknown;
}

export type RequestParams = Omit<FullRequestParams, "body" | "method" | "query" | "path">;

export interface ApiConfig<SecurityDataType = unknown> extends Omit<AxiosRequestConfig, "data" | "cancelToken"> {
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<AxiosRequestConfig | void> | AxiosRequestConfig | void;
  secure?: boolean;
  format?: ResponseType;
}

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public instance: AxiosInstance;
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private secure?: boolean;
  private format?: ResponseType;

  constructor({ securityWorker, secure, format, ...axiosConfig }: ApiConfig<SecurityDataType> = {}) {
    this.instance = axios.create({ ...axiosConfig, baseURL: axiosConfig.baseURL || "" });
    this.secure = secure;
    this.format = format;
    this.securityWorker = securityWorker;
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  private mergeRequestParams(params1: AxiosRequestConfig, params2?: AxiosRequestConfig): AxiosRequestConfig {
    return {
      ...this.instance.defaults,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.instance.defaults.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  private createFormData(input: Record<string, unknown>): FormData {
    return Object.keys(input || {}).reduce((formData, key) => {
      const property = input[key];
      formData.append(
        key,
        property instanceof Blob
          ? property
          : typeof property === "object" && property !== null
          ? JSON.stringify(property)
          : `${property}`,
      );
      return formData;
    }, new FormData());
  }

  public request = async <T = any, _E = any>({
    secure,
    path,
    type,
    query,
    format,
    body,
    ...params
  }: FullRequestParams): Promise<AxiosResponse<T>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const responseFormat = (format && this.format) || void 0;

    if (type === ContentType.FormData && body && body !== null && typeof body === "object") {
      requestParams.headers.common = { Accept: "*/*" };
      requestParams.headers.post = {};
      requestParams.headers.put = {};

      body = this.createFormData(body as Record<string, unknown>);
    }

    return this.instance.request({
      ...requestParams,
      headers: {
        ...(type && type !== ContentType.FormData ? { "Content-Type": type } : {}),
        ...(requestParams.headers || {}),
      },
      params: query,
      responseType: responseFormat,
      data: body,
      url: path,
    });
  };
}

/**
 * @title consumer/pricefeed/genesis.proto
 * @version version not set
 */
export class Api<SecurityDataType extends unknown> extends HttpClient<SecurityDataType> {
  /**
   * No description
   *
   * @tags Query
   * @name QueryParams
   * @summary Parameters queries the parameters of the module.
   * @request GET:/consumer/pricefeed/params
   */
  queryParams = (params: RequestParams = {}) =>
    this.request<PricefeedQueryParamsResponse, RpcStatus>({
      path: `/consumer/pricefeed/params`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QueryPrice
   * @request GET:/consumer/pricefeed/price/{symbol}
   */
  queryPrice = (symbol: string, params: RequestParams = {}) =>
    this.request<PricefeedQueryPriceResponse, RpcStatus>({
      path: `/consumer/pricefeed/price/${symbol}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySymbolRequest
   * @request GET:/consumer/pricefeed/symbol/{symbol}
   */
  querySymbolRequest = (symbol: string, params: RequestParams = {}) =>
    this.request<PricefeedQuerySymbolRequestResponse, RpcStatus>({
      path: `/consumer/pricefeed/symbol/${symbol}`,
      method: "GET",
      format: "json",
      ...params,
    });

  /**
   * No description
   *
   * @tags Query
   * @name QuerySymbolRequests
   * @request GET:/consumer/pricefeed/symbols
   */
  querySymbolRequests = (params: RequestParams = {}) =>
    this.request<PricefeedQuerySymbolRequestsResponse, RpcStatus>({
      path: `/consumer/pricefeed/symbols`,
      method: "GET",
      format: "json",
      ...params,
    });
}