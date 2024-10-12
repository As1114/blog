import { Outlet } from "react-router-dom";
import "./index.css";
import { Nav } from "@/components/nav/nav";

export function Web() {
  return (
    <div className="web">
      <Nav></Nav>
      <main>
        <Outlet></Outlet>
      </main>
    </div>
  );
}
