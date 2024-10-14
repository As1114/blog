import router from "./router";
import { RouterProvider } from "react-router-dom";

export function App() {
  return (
    <div className="app">
      <RouterProvider router={router} />
    </div>
  );
}
