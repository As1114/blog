"use client";
import React, { useState } from "react";
import { ProductItem, Menu, MenuItem } from "../ui/navbar-menu";
import { cn } from "@/lib/utils";

export function NavbarDemo() {
  return <Navbar />;
}

function Navbar({ className }: { className?: string }) {
  const [active, setActive] = useState<string | null>(null);
  return (
    <div className={cn("", className)}>
      <Menu setActive={setActive}>
        <MenuItem setActive={setActive} active={active} item="Blog">
          <ProductItem title="Blog" src="" description="" />
        </MenuItem>
        <MenuItem setActive={setActive} active={active} item="Record">
          <ProductItem title="Record" src="" description="" />
        </MenuItem>
        <MenuItem setActive={setActive} active={active} item="Timeline">
          <ProductItem title="Timeline" src="" description="" />
        </MenuItem>
      </Menu>
    </div>
  );
}
