"use client";

import React from "react";
import { CardBody, CardContainer, CardItem } from "../ui/3d-card";

export function ThreeDCardDemo() {
  return (
    <CardContainer className="inter-var">
      <CardBody className="bg-gray-50 relative group/card border-black/[0.1] w-auto h-auto rounded-xl p-6 border  ">
        <CardItem translateZ="60">Make things float in air</CardItem>
        <CardItem as="p" translateZ="60">
          Hover over this card to unleash the power of CSS perspective
        </CardItem>
        <CardItem translateZ="60">
          -----------------------------------------------------------------
        </CardItem>
      </CardBody>
    </CardContainer>
  );
}
