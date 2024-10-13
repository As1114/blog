import { cn } from "@/lib/utils";
import Marquee from "@/components/ui/marquee";
import { AnimatedPinDemo } from "./3d_animated_pin_demo";

const ReviewCard = ({
  title,
  description,
  link,
  thumbnail,
}: {
  title: string;
  description?: string;
  link: string;
  thumbnail: string;
}) => {
  return (
    <AnimatedPinDemo
      title={title}
      description={description}
      link={link}
      thumbnail={thumbnail}
    ></AnimatedPinDemo>
  );
};

export const MarqueeDemo = ({
  reviews,
}: {
  reviews: {
    title: string;
    description?: string;
    link: string;
    thumbnail: string;
  }[];
}) => {
  const firstRow = reviews.slice(0, reviews.length / 2);
  const secondRow = reviews.slice(reviews.length / 2);
  return (
    <div className="relative flex h-[550px] w-[1200px] flex-col items-center justify-center overflow-hidden rounded-lg bg-background">
      <Marquee pauseOnHover className="[--duration:20s]">
        {firstRow.map((review) => (
          <ReviewCard key={review.title} {...review} />
        ))}
      </Marquee>
      <Marquee reverse pauseOnHover className="[--duration:20s]">
        {secondRow.map((review) => (
          <ReviewCard key={review.title} {...review} />
        ))}
      </Marquee>
      {/* <div className="pointer-events-none absolute inset-y-0 left-0 w-1/4 bg-gradient-to-r from-white"></div>
      <div className="pointer-events-none absolute inset-y-0 right-0 w-1/4 bg-gradient-to-l from-white"></div> */}
    </div>
  );
};
