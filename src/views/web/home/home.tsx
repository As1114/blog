﻿import "./home.css";
import { OrbitingCirclesDemo } from "@/components/test/orbiting_circles_demo";
import { HeroParallaxDemo } from "@/components/test/hero-parallax_demo";
import { FadeComponent } from "@/components/ui/fade-component";
export function Home() {
  return (
    <div className="home">
      <FadeComponent
        direction="left"
        framerProps={{
          show: { transition: { delay: 0.5 } },
        }}
        className="circle_fade"
      >
        <OrbitingCirclesDemo></OrbitingCirclesDemo>
      </FadeComponent>
      <HeroParallaxDemo></HeroParallaxDemo>
    </div>
  );
}
