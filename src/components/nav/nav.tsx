import { Navleft } from "../nav_left/nav_left";
import { Navright } from "../nav_right/nav_right";
import "./nav.css";
export function Nav() {
  return (
    <div className="nav">
      <div className="left">
        <Navleft></Navleft>
      </div>
      <div className="right">
        <Navright></Navright>
      </div>
    </div>
  );
}
