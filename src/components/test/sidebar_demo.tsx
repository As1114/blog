"use client";

import { useState } from "react";
import { Sidebar, SidebarBody } from "../ui/sidebar";
import { cn } from "@/lib/utils";

export function SidebarDemo() {
  const [open, setOpen] = useState(false);

  return (
    <div
      className={cn(
        "rounded-md flex flex-col md:flex-row bg-white w-full flex-1 mx-auto",
        "border border-neutral-200 overflow-hidden h-screen"
      )}
    >
      <Sidebar open={open} setOpen={setOpen}>
        <SidebarBody className="justify-between gap-10"></SidebarBody>
      </Sidebar>
      <Dashboard />
    </div>
  );
}

// Dummy dashboard component with content
const Dashboard = () => (
  <div className="flex flex-1">
    <div className="p-2 md:p-5 border border-neutral-200 bg-white flex flex-col gap-2 flex-1 w-full h-full"></div>
  </div>
);
