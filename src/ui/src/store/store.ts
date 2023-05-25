import { configureStore, Store } from "@reduxjs/toolkit";
import { RootState } from "./types/rootState";
import { STORE_NAME } from "../constants/environment";
import rootReducers from "./reducers/rootReducers";
import grpcGatewayApi from "./api/grpcGatewayApi";

const store: Store<RootState> = configureStore({
  reducer: rootReducers,
  devTools: {
    name: STORE_NAME,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(grpcGatewayApi.middleware),
});

export const getStoreWithState = (preloadedState?: RootState) =>
  configureStore({ reducer: rootReducers, preloadedState });

export default store;
