import { OrbitingCirclesDemo } from "@/components/test/orbiting_circles_demo";
import "./home.css";
import { HeroParallaxDemo } from "@/components/test/hero-parallax_demo";
export function Home() {
  return (
    <div className="Home">
      <OrbitingCirclesDemo></OrbitingCirclesDemo>
      <HeroParallaxDemo></HeroParallaxDemo>
    </div>
  );
}
