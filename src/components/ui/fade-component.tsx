"use client";

import { useMemo } from "react";
import { motion, Variants } from "framer-motion";

type FadeComponentProps = {
  className?: string;
  direction?: "up" | "down" | "left" | "right";
  framerProps?: Variants;
  children: React.ReactNode;
};

export function FadeComponent({
  direction = "up",
  className,
  framerProps,
  children,
}: FadeComponentProps) {
  // 提取默认方向的偏移值，避免在useMemo中进行
  const directionOffset = direction === "up" || direction === "right" ? 50 : -50;
  const axis = direction === "up" || direction === "down" ? "y" : "x";

  // 优化动画配置，减少多次解构和属性覆盖
  const FADE_ANIMATION_VARIANTS = useMemo(() => {
    const defaultVariants: Variants = {
      hidden: { opacity: 0, [axis]: directionOffset },
      show: { opacity: 1, [axis]: 0, transition: { type: "spring" } },
    };

    return {
      ...defaultVariants,
      ...framerProps,
      hidden: { ...defaultVariants.hidden, ...(framerProps?.hidden || {}) },
      show: { ...defaultVariants.show, ...(framerProps?.show || {}) },
    };
  }, [directionOffset, axis, framerProps]);

  return (
    <motion.div
      initial="hidden"
      animate="show"
      viewport={{ once: true }}
      variants={FADE_ANIMATION_VARIANTS}
      className={className}
    >
      {children}
    </motion.div>
  );
}
