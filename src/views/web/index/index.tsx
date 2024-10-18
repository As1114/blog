import "./index.css";
import { Outlet } from "react-router-dom";
import { Nav } from "@/components/nav/nav";
import { WebProvider } from "@/store/web_context";

export function Web() {

  return (
    <WebProvider>
      <div className="web">
        <Nav></Nav>
        <main>
          <Outlet></Outlet>
        </main>
      </div>
    </WebProvider>
  );
}

