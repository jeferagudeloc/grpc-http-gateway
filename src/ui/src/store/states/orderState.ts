import { OrdersResponseMock } from "../../__mocks__/ordersMock";

export interface Order {
  id: string;
  store: string;
  order: string;
  ip: string;
  capacity: string;
  memory: string;
  status: string;
}

export interface OrdersState {
  data: Order[];
}

export const OrdersStateDefault: OrdersState = {
  data: OrdersResponseMock,
};
