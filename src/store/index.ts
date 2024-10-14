import { configureStore } from "@reduxjs/toolkit";
import webReducer from "./web_slice";

const store = configureStore({
  reducer: {
    web: webReducer,
  },
});

export default store;
