import { HTTP_GATEWAY_GRPC_API, API_URL } from "../../constants/api";
import {
  BaseQueryFn,
  createApi,
  FetchArgs,
  fetchBaseQuery,
  FetchBaseQueryError,
  retry,
} from "@reduxjs/toolkit/query/react";
import {
  ORDER_URL,
  USER_URL,
} from "../../constants/urls";
import {
  Order,
} from "../states/orderState";

import {
  User,
} from "../states/userState";

const baseQuery = fetchBaseQuery({
  baseUrl: `${API_URL}`,
  prepareHeaders: (headers: Headers) => {
    headers.set("Content-Type", "application/json");
    headers.set('Access-Control-Allow-Origin', '*');
    return headers;
  },
  mode: 'no-cors',
});

const confBaseQuery: BaseQueryFn<
  string | FetchArgs,
  unknown,
  FetchBaseQueryError
> = async (args, api, extraOptions) => {
  let result = await baseQuery(args, api, extraOptions);
  return result;
};

const grpcGatewayApi = createApi({
  reducerPath: HTTP_GATEWAY_GRPC_API,
  baseQuery: retry(confBaseQuery, {
    maxRetries: 1,
  }),
  
  tagTypes: ["order", "user"],
  endpoints: (build) => ({
    getOrdersGrpc: build.query<Order, void>({
      query: () => `/grpc/${ORDER_URL}`,
      providesTags: ["order"],
      extraOptions: {
        maxRetries: 0,
      },   
    }),
    getUsersGrpc: build.query<User, void>({
      query: () => `/grpc/${USER_URL}`,
      providesTags: ["user"],
      extraOptions: {
        maxRetries: 0,
      },   
    }),
    getOrdersHttp: build.query<Order, void>({
      query: () => `/http/${ORDER_URL}`,
      providesTags: ["order"],
      extraOptions: {
        maxRetries: 0,
      },   
    }),
    getUsersHttp: build.query<User, void>({
      query: () => `/http/${USER_URL}`,
      providesTags: ["user"],
      extraOptions: {
        maxRetries: 0,
      },   
    }),
  }),
});

export const {
  useGetOrdersGrpcQuery,
  useGetUsersGrpcQuery,
  useGetOrdersHttpQuery,
  useGetUsersHttpQuery,
} = grpcGatewayApi;

export default grpcGatewayApi;
