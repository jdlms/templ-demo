import { Suspense } from "react";
import { Canvas, extend } from "@react-three/fiber";
import { OrbitControls } from "@react-three/drei";
import { Loader } from "./Loader";
import { Spinner } from "./Spinner";
import * as THREE from "three";

// Extend fiber with Three.js objects
extend(THREE);

export function Container() {
  return (
    <div className=" flex justify-center h-52 mt-8 mb-12 md:h-80 md:mt-0">
      <Canvas
        style={{ height: "80%", width: "80%" }}
        camera={{ position: [0, 7, -10], fov: 75 }}
        className="cursor-pointer rounded-full"
      >
        {/* @ts-ignore */}
        <ambientLight intensity={2} />
        {/* @ts-ignore */}
        <pointLight
          position={[0, 18, -10]}
          intensity={500}
          color="#fff"
          distance={0}
        />
        <Suspense fallback={<Spinner />}>
          <Loader />
        </Suspense>
        <OrbitControls
          enableZoom={false}
          enablePan={false}
          autoRotate={true}
          autoRotateSpeed={4.5}
          target={[2, 2.9, 1]}
        />
      </Canvas>
    </div>
  );
}
