"use client";
import React from "react";
import { HeroParallax } from "../ui/hero-parallax";
import { MarqueeDemo } from "./marquee";
import { AnimatedPinDemo } from "./3d_animated_pin";
export function HeroParallaxDemo() {
  return (
    <div>
      <HeroParallax products={products} />
      {/* <MarqueeDemo reviews={products}></MarqueeDemo> */}
      {/* <AnimatedPinDemo
        title={products[0].title}
        description={products[0].description}
        link={products[0].link}
        thumbnail={products[0].thumbnail}
      ></AnimatedPinDemo> */}
    </div>
  );
}
export const products = [
  {
    title: "Moonbeam",
    description:
      "152465达瓦伟大伟大·无法达瓦达瓦放大违法·二等功v色如法国撒饿饭·的歌大无法让违法我i非把我i可否把我iu发表·dawefgswergr官方白色帆布萨尔ui发红包斯恶u回复王府井埃斯珀i俄方祭祀热哦帕金斯哦热狗",
    link: "https://gomoonbeam.com",
    thumbnail: "../../../public/images/image.png",
  },
  {
    title: "Meam",
    description:
      "152465达瓦伟大伟大·无法达瓦达瓦放大违法·二等功v色如法国撒饿饭·的歌大无法让违法我i非把我i可否把我iu发表·dawefgswergr官方白色帆布萨尔ui发红包斯恶u回复王府井埃斯珀i俄方祭祀热哦帕金斯哦热狗",
    link: "https://gomoonbeam.com",
    thumbnail: "../../../public/images/image.png",
  },
  {
    title: "Cursor",
    link: "https://cursor.so",
    description: "152465",
    thumbnail:
      "https://aceternity.com/images/products/thumbnails/new/cursor.png",
  },
  {
    title: "Jack",
    link: "https://cursor.so",
    description: "152465",
    thumbnail:
      "https://aceternity.com/images/products/thumbnails/new/cursor.png",
  },
  {
    title: "Jack",
    link: "https://cursor.so",
    description: "152465",
    thumbnail:
      "https://aceternity.com/images/products/thumbnails/new/cursor.png",
  },
  {
    title: "Jack",
    link: "https://cursor.so",
    description: "152465",
    thumbnail:
      "https://aceternity.com/images/products/thumbnails/new/cursor.png",
  },
];
