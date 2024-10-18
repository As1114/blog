﻿import { Admin } from "@/views/admin";
import { Web } from "@/views/web/index";
import { Home } from "@/views/web/home/home";
import Notfound from "@/views/web/not_found/not_found";
import { createBrowserRouter } from "react-router-dom";
import { Timeline } from "@/views/web/timeline/timeline";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Web />,
    children: [
      {
        path: "",
        element: <Home />,
      },
      {
        path: "timeline",
        element: <Timeline />,
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
