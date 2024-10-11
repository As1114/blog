import { Outlet } from "react-router-dom";
import "./index.css";
import { Nav } from "@/components/nav/nav";

export function Web() {
  return (
    <div className="web">
      <div className="web_nav">
        <Nav></Nav>
      </div>
      <main>
        <Outlet></Outlet>
      </main>
    </div>
  );
}
