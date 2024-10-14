"use client";
import React from "react";
import { cva } from "class-variance-authority";
import { cn } from "@/lib/utils";
import { useWebContext } from "@/store/web_context";

// 定义类型
type OverlayVariant = "none" | "light" | "dark";
type MediaType = "image" | "video";

// 样式变体
const backgroundVariants = cva(
  "relative h-screen max-h-[1000px] w-full min-h-[500px] lg:min-h-[600px]",
  {
    variants: {
      overlay: {
        none: "",
        light:
          "before:absolute before:inset-0 before:bg-white before:opacity-30",
        dark: "before:absolute before:inset-0 before:bg-black before:opacity-30",
      },
      type: {
        image: "",
        video: "z-10",
      },
    },
    defaultVariants: {
      overlay: "none",
      type: "image",
    },
  }
);

// 组件属性接口
interface BackgroundMediaProps {
  variant?: OverlayVariant;
  type?: MediaType;
  src: string;
  alt?: string;
}

// BackgroundMedia 组件
const BackgroundMedia: React.FC<BackgroundMediaProps> = ({
  variant = "light",
  type = "image",
  src,
  alt = "",
}) => {
  const { mediaRef, isPlaying } = useWebContext();

  // 渲染媒体内容
  const renderMedia = () => {
    const commonProps = {
      className:
        "absolute inset-0 h-full w-full object-cover transition-opacity duration-300",
    };

    if (type === "video") {
      return (
        <video
          ref={mediaRef}
          aria-hidden="true"
          muted
          autoPlay
          playsInline
          loop
          {...commonProps}
        >
          <source src={src} type="video/mp4" />
          Your browser does not support the video tag.
        </video>
      );
    } else {
      return (
        <img
          src={src}
          alt={alt}
          loading="eager"
          className={`${commonProps.className} rounded-br-[88px]`}
        />
      );
    }
  };

  return (
    <div
      className={cn(
        backgroundVariants({ overlay: variant, type }),
        "overflow-hidden fixed inset-0 z-0"
      )}
    >
      {renderMedia()}
      {type === "video" && (
        <div className="absolute bottom-4 right-4 z-50">
          <span className="px-4 py-2 bg-gray-900 text-white hover:bg-gray-700">
            {isPlaying ? "Playing" : "Paused"}
          </span>
        </div>
      )}
    </div>
  );
};

export default BackgroundMedia;
