import { Order } from "../store/states/orderState";

export const OrdersResponseMock: Order[] = [
  {
    id: "1",
    order: "Order 1",
    store: "Americanino (Stand Alone)",
    ip: "192.168.0.1",
    capacity: "500 GB",
    memory: "16 GB",
    status: "active"
  },
  {
    id: "2",
    order: "Order 2",
    store: "Americanino (Stand Alone)",
    ip: "192.168.0.2",
    capacity: "1 TB",
    memory: "32 GB",
    status: "inactive"
  }
];
