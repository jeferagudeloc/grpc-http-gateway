import { combineReducers } from "@reduxjs/toolkit";
import orderSlice from "../slices/order";
import userSlice from "../slices/user";
import grpcGatewayApi from "../api/grpcGatewayApi";

const rootReducers = combineReducers({
  order: orderSlice,
  user: userSlice,
  [grpcGatewayApi.reducerPath]: grpcGatewayApi.reducer,
});

export default rootReducers;
