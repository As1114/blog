import { createSlice } from "@reduxjs/toolkit";
const webSlice = createSlice({
  name: "web",
  initialState: {
    user: { id: 0, nick_name: "", avatar: "", role: 0, token: "" },
  },
  reducers: {},
});

export const {} = webSlice.actions;
export default webSlice.reducer;
