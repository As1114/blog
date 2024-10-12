import { Outlet } from "react-router-dom";
import "./index.css";
import { Nav } from "@/components/nav/nav";
import { FadeComponent } from "@/components/ui/fade-component";

export function Web() {
  return (
    <div className="web">
      <FadeComponent
        framerProps={{
          show: { transition: { delay: 0.6 } },
        }}
        direction="right"
      >
        <Nav></Nav>
      </FadeComponent>
      <main>
        <Outlet></Outlet>
      </main>
    </div>
  );
}
