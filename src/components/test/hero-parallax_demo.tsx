"use client";
import React from "react";
import { HeroParallax } from "../ui/hero-parallax";

export function HeroParallaxDemo() {
  return <HeroParallax products={products} />;
}
export const products = [
  {
    title: "blog",
    description: "react+gin",
    link: "https://github.com/As1114/blog",
    thumbnail: "public/images/image.png",
  },
  {
    title: "blog2",
    description: "react+gin",
    link: "https://github.com/As1114/blog",
    thumbnail: "public/images/image.png",
  },
  {
    title: "blog3",
    description: "react+gin",
    link: "https://github.com/As1114/blog",
    thumbnail: "public/images/image.png",
  },
  {
    title: "blog4",
    description: "react+gin",
    link: "https://github.com/As1114/blog",
    thumbnail: "public/images/image.png",
  },
  {
    title: "blog5",
    description: "react+gin",
    link: "https://github.com/As1114/blog",
    thumbnail: "public/images/image.png",
  },
];
