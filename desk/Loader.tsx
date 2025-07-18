import React, { useMemo } from "react";
import { useLoader } from "@react-three/fiber";
import { GLTFLoader } from "three/examples/jsm/Addons.js";

export function Loader() {
  const { materials, scene } = useLoader(GLTFLoader, "/models/Desk.glb");
  useMemo(() => {
    Object.values(materials).forEach((material: any) => {
      material.color.set("#efd8cb");
      material.needsUpdate = true;
    });
  }, [materials]);

  return React.createElement("primitive", { object: scene });
}
