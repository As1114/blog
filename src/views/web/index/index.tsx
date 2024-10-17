import "./index.css";
import { Outlet } from "react-router-dom";
import { Nav } from "@/components/nav/nav";
import BackgroundMedia from "@/components/ui/backgroundmedia";
import { WebProvider } from "@/store/web_context";

export function Web() {

  return (
    <WebProvider>
      <BackgroundMedia
        type="video"
        variant="none"
        src="/public/media/185096-874643413.mp4"
      ></BackgroundMedia>
      <div className="web">
        <Nav></Nav>
        <main>
          <Outlet></Outlet>
        </main>
      </div>
    </WebProvider>
  );
}

