import React, {
  createContext,
  useContext,
  useState,
  ReactNode,
  useRef,
} from "react";

interface MediaContextType {
  isPlaying: boolean;
  togglePlay: () => void;
  mediaRef: React.RefObject<HTMLVideoElement>;
}

const MediaContext = createContext<MediaContextType | undefined>(undefined);

export const MediaProvider = ({ children }: { children: ReactNode }) => {
  const mediaRef = useRef<HTMLVideoElement | null>(null);
  const [isPlaying, setIsPlaying] = useState(false);

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
    <MediaContext.Provider value={{ isPlaying, togglePlay, mediaRef }}>
      {children}
    </MediaContext.Provider>
  );
};

export const useMediaContext = () => {
  const context = useContext(MediaContext);
  if (!context) {
    throw new Error("useMediaContext must be used within a MediaProvider");
  }
  return context;
};
