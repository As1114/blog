import "./home.css";
import { OrbitingCirclesDemo } from "@/components/test/orbiting_circles_demo";
import { HeroParallaxDemo } from "@/components/test/hero-parallax_demo";
import { FadeComponent } from "@/components/ui/fade-component";
export function Home() {
  return (
    <div className="home">
      <FadeComponent
        direction="left"
        framerProps={{
          show: { transition: { delay: 0.8 } },
        }}
        className="circle_fade"
      >
        <div>左 天气日历 数据</div>
        <OrbitingCirclesDemo></OrbitingCirclesDemo>
        <div>右 音乐</div>
      </FadeComponent>
      <FadeComponent>
        <HeroParallaxDemo></HeroParallaxDemo>
      </FadeComponent>
    </div>
  );
}
