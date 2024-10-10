import { Navleft } from "../nav_left/nav_left";
import { Navright } from "../nav_right/nav_right";
import "./nav.css";
export function Nav() {
  return (
    <div className="nav">
      <Navleft></Navleft>
      <Navright></Navright>
    </div>
  );
}
