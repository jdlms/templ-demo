import { LineWave } from "react-loader-spinner";
import { Html } from "@react-three/drei";

export function Spinner() {
  return (
    <Html prepend center>
      <div className="bg-neutral-800 w-44 h-44 md:h-64 md:w-64 cursor-pointer rounded-full  shadow-inner">
        <LineWave
          visible={true}
          height="90"
          width="90"
          firstLineColor="#ad7753"
          middleLineColor="#ffd700"
          lastLineColor="#ff6b6b"
          ariaLabel="line-wave-loading"
          wrapperStyle={{
            display: "flex",
            justifyContent: "center",
            alignItems: "center",
            marginBottom: "12rem",
          }}
          wrapperClass="spinner"
        />
      </div>
    </Html>
  );
}
