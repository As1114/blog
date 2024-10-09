import { Admin } from "@/views/admin";
import { Web } from "@/views/web";
import { Home } from "@/views/web/home";
import Notfound from "@/views/web/not_found";
import { createBrowserRouter } from "react-router-dom";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Web />,
    children: [
      {
        path: "",
        element: <Home />,
      },
    ],
  },
  {
    path: "/admin",
    element: <Admin />,
  },
  {
    path: "*",
    element: <Notfound />,
  },
]);

export default router;
