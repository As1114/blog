import React, {
  createContext,
  useContext,
  useState,
  ReactNode,
  useRef,
} from "react";

interface WebContextType {
  isPlaying: boolean;
  togglePlay: () => void;
  mediaRef: React.RefObject<HTMLVideoElement>;
}

const WebContext = createContext<WebContextType | undefined>(undefined);

export const WebProvider = ({ children }: { children: ReactNode }) => {
  const mediaRef = useRef<HTMLVideoElement | null>(null);
  const [isPlaying, setIsPlaying] = useState(true);

  const togglePlay = () => {
    if (mediaRef.current) {
      if (isPlaying) {
        mediaRef.current.pause();
      } else {
        mediaRef.current.play();
      }
      setIsPlaying(!isPlaying);
    }
  };

  return (
    <WebContext.Provider value={{ isPlaying, togglePlay, mediaRef }}>
      {children}
    </WebContext.Provider>
  );
};

export const useWebContext = () => {
  const context = useContext(WebContext);
  if (!context) {
    throw new Error("useWebContext must be used within a WebProvider");
  }
  return context;
};
