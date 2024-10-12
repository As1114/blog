"use client";

import Confetti from "../ui/confetti";
import confetti from "canvas-confetti";

export function ConfettiDemo() {
  //   const handleClick = () => {
  //     const end = Date.now() + 3 * 1000;
  //     const colors = ["#a786ff", "#fd8bbc", "#eca184", "#f8deb1"];

  //     const frame = () => {
  //       if (Date.now() > end) return;

  //       confetti({
  //         particleCount: 2,
  //         angle: 60,
  //         spread: 55,
  //         startVelocity: 60,
  //         origin: { x: 0, y: 0.5 },
  //         colors: colors,
  //       });
  //       confetti({
  //         particleCount: 2,
  //         angle: 120,
  //         spread: 55,
  //         startVelocity: 60,
  //         origin: { x: 1, y: 0.5 },
  //         colors: colors,
  //       });

  //       requestAnimationFrame(frame);
  //     };

  //     frame();
  //   };
  return (
    <div className="relative flex h-[80px] w-full flex-col items-center justify-center overflow-hidden rounded-lg  bg-background">
      <span className="pointer-events-none whitespace-pre-wrap bg-gradient-to-b from-black to-gray-300/80 bg-clip-text text-center font-semibold leading-none text-transparent">
        NSXZ
      </span>
      {/* <Confetti
        className="absolute left-0 top-0 z-0 size-full"
        onMouseEnter={handleClick}
      /> */}
    </div>
  );
}
