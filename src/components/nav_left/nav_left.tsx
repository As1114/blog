import "./nav_left.css";
import { NavbarDemo } from "../test/navbar_demo";
import { ConfettiDemo } from "../test/confettid_demo";

export function Navleft() {
  return (
    <div className="nav_left">
      <div className="slogan">
        <ConfettiDemo></ConfettiDemo>
      </div>
      <div className="function">
        <NavbarDemo></NavbarDemo>
      </div>
    </div>
  );
}
