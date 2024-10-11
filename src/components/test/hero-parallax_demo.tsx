"use client";
import React from "react";
import { HeroParallax } from "../ui/hero-parallax";

export function HeroParallaxDemo() {
  return (
    <div>
      <HeroParallax products={products} />
    </div>
  );
}
export const products = [
  {
    title: "Meam",
    description: "1232456",
    link: "https://gomoonbeam.com",
    thumbnail:
      "https://aceternity.com/images/products/thumbnails/new/invoker.png",
  },
];
