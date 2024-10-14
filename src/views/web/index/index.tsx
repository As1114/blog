import "./index.css";
import { Outlet } from "react-router-dom";
import { Nav } from "@/components/nav/nav";
import BackgroundMedia from "@/components/ui/backgroundmedia";
import { useWebContext, WebProvider } from "@/store/web_context";
const GlobalPlayButton = () => {
  const { togglePlay, isPlaying } = useWebContext();

  return (
    <button
      onClick={togglePlay}
      className="fixed bottom-4 left-4 z-50 px-4 py-2 bg-gray-900 text-white hover:bg-gray-700"
    >
      {isPlaying ? "Pause Background" : "Play Background"}
    </button>
  );
};
export function Web() {
  return (
    <WebProvider>
      <BackgroundMedia
        type="video"
        variant="none"
        src="/public/media/185096-874643413.mp4"
      ></BackgroundMedia>
      <GlobalPlayButton></GlobalPlayButton>
      <div className="web">
        <Nav></Nav>
        <main>
          <Outlet></Outlet>
        </main>
      </div>
    </WebProvider>
  );
}
