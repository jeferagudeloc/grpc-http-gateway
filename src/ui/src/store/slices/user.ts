import { createSlice } from "@reduxjs/toolkit";
import { UserStateDefault } from "../states/userState";

const userSlice = createSlice({
    name: 'user',
    initialState: UserStateDefault,
    reducers: {}
})


export default userSlice.reducer;