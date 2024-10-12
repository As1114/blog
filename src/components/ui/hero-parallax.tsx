"use client";
import React from "react";
import {
  motion,
  useScroll,
  useTransform,
  useSpring,
  MotionValue,
} from "framer-motion";
import Image from "next/image";
import Link from "next/link";
import { MarqueeDemo } from "../test/marquee_demo";

export const HeroParallax = ({
  products,
}: {
  products: {
    title: string;
    description?: string;
    link: string;
    thumbnail: string;
  }[];
}) => {
  const ref = React.useRef(null);
  const { scrollYProgress } = useScroll({
    target: ref,
    offset: ["start start", "end start"],
  });

  // 定义弹簧配置
  const springConfig = { stiffness: 300, damping: 30, bounce: 100 };

  // 正面倾斜度
  const rotateX = useSpring(
    useTransform(scrollYProgress, [0, 0.1], [1, 0]),
    springConfig
  );

  // 透明度
  const opacity = useSpring(
    useTransform(scrollYProgress, [0, 0.2], [1, 1]),
    springConfig
  );

  // 倾斜角度
  const rotateZ = useSpring(
    useTransform(scrollYProgress, [0, 0.2], [10, 0]),
    springConfig
  );

  // 背景位置
  const translateY = useSpring(
    useTransform(scrollYProgress, [0, 0.2], [-100, 500]),
    springConfig
  );

  return (
    <div
      ref={ref}
      className="h-[200vh] py-40 overflow-hidden antialiased relative flex flex-col self-auto [perspective:1000px] [transform-style:preserve-3d]"
    >
      <motion.div
        style={{
          rotateX,
          rotateZ,
          translateY,
          opacity,
        }}
        className=""
      >
        <motion.div className="flex flex-row mb-20 space-x-20 justify-center">
          <MarqueeDemo reviews={products}></MarqueeDemo>
        </motion.div>
      </motion.div>
    </div>
  );
};

export const ProductCard = ({
  product,
}: {
  product: {
    title?: string;
    link: string;
    thumbnail: string;
  };
  translate?: MotionValue<number>;
}) => {
  return (
    <motion.div
      whileHover={{
        y: -10,
      }}
      key={product.title}
      className="group/product h-[15rem] w-[20rem] relative flex-shrink-0"
    >
      <Link href={product.link}>
        <Image
          loader={() => product.thumbnail}
          src={product.thumbnail}
          height="600"
          width="600"
          className="object-cover object-left-top absolute h-full w-full inset-0"
          alt={product.link}
          unoptimized
          priority
        ></Image>
      </Link>
    </motion.div>
  );
};
