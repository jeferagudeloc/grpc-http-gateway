import grpcGatewayApi from "../store/api/grpcGatewayApi";


export const useOrders = () => {
  const { orders: ordersByGrpc } = grpcGatewayApi.useGetOrdersGrpcQuery(undefined, {
    selectFromResult: ({ data }) => ({
        orders: data,
    }),
  });

  const { orders: ordersByHttp } = grpcGatewayApi.useGetOrdersHttpQuery(undefined, {
    selectFromResult: ({ data }) => ({
        orders: data,
    }),
  });

  return {
    ordersByGrpc,
    ordersByHttp,
  };
};
