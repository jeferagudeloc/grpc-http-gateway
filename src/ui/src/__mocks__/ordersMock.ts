import { Order } from "../store/states/orderState";

export const OrdersResponseMock: Order[] = [
  {
    id: "1",
    orderType: "Order 1",
    store: "Americanino (Stand Alone)",
    address: "192.168.0.1",
    creationDate: "500 GB",
    status: "16 GB",
  },
  {
    id: "2",
    orderType: "Order 2",
    store: "Americanino (Stand Alone)",
    address: "192.168.0.2",
    creationDate: "1 TB",
    status: "32 GB",
  }
];
