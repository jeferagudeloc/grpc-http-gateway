import { OrdersResponseMock } from "../../__mocks__/ordersMock";

export interface Order {
  id: string;
  orderType: string;
  store: string;
  address: string;
  creationDate: string;
  status: string;
}

export interface OrdersState {
  data: Order[];
}

export const OrdersStateDefault: OrdersState = {
  data: OrdersResponseMock,
};
