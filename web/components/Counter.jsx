import React, { useState } from "react";

export default function Counter() {
  const [count, setCount] = useState(0);
  return (
    <button
      style={{
        border: "red",
        padding: "12px 24px",
        fontSize: "18px",
        margin: "16px",
        cursor: "pointer",
      }}
      onClick={() => setCount(count + 1)}
    >
      Count: {count}
    </button>
  );
}
