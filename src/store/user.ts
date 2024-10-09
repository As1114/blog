import { createSlice } from "@reduxjs/toolkit";
const userSlice = createSlice({
  name: "user",
  initialState: {
    id: 0,
    nick_name: "",
    avatar: "",
    role: 0,
    token: "",
  },
  reducers: {},
});

export const isLogin = (state: any) => state.user.role !== 0;

export const {} = userSlice.actions;
export default userSlice.reducer;
