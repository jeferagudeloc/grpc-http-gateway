import { createSlice } from "@reduxjs/toolkit";
import { OrdersStateDefault } from "../states/orderState";

const ordersSlice = createSlice({
    name: 'order',
    initialState: OrdersStateDefault,
    reducers: {}
})


export default ordersSlice.reducer;