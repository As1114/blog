import router from "./router";
import { RouterProvider } from "react-router-dom";
import { MediaProvider, useMediaContext } from "./store/bgmediacontext";
import BackgroundMedia from "./components/ui/backgroundmedia";
const GlobalPlayButton = () => {
  const { togglePlay, isPlaying } = useMediaContext();

  return (
    <button
      onClick={togglePlay}
      className="fixed bottom-4 left-4 z-50 px-4 py-2 bg-gray-900 text-white hover:bg-gray-700"
    >
      {isPlaying ? "Pause Background" : "Play Background"}
    </button>
  );
};
export function App() {
  return (
    // <MediaProvider>
    //   <BackgroundMedia
    //     type="video"
    //     variant="none"
    //     src="/public/media/185096-874643413.mp4"
    //   />
    //   <GlobalPlayButton />
      <div className="app">
        <RouterProvider router={router} />
      </div>
    // </MediaProvider>
  );
}
