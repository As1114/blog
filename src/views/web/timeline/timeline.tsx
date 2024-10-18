import { TimelineDemo } from "@/components/test/timeline_demo";
import "./timeline.css";
import BackgroundMedia from "@/components/ui/backgroundmedia";

export function Timeline() {
    return (
        <div className="timeline">
            <BackgroundMedia
                type="image"
                variant="none"
                src="/public/images/mountains.png"
            ></BackgroundMedia>
            <TimelineDemo></TimelineDemo>
        </div>
    );
}
