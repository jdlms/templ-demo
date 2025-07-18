import React, { useState, useEffect } from "react";

export default function Counter() {
  const [count, setCount] = useState(0);
  const [loading, setLoading] = useState(false);

  // Load initial count from API
  useEffect(() => {
    fetch('/api/counter')
      .then(res => res.json())
      .then(data => setCount(data.count))
      .catch(err => console.error('Failed to load counter:', err));
  }, []);

  const handleIncrement = async () => {
    if (loading) return; // Prevent double clicks
    
    setLoading(true);
    try {
      const response = await fetch('/api/counter/increment', {
        method: 'POST',
      });
      const data = await response.json();
      setCount(data.count);
    } catch (err) {
      console.error('Failed to increment counter:', err);
    } finally {
      setLoading(false);
    }
  };

  return (
    <button
      style={{
        border: "1px solid #ccc",
        padding: "12px 24px",
        fontSize: "18px",
        margin: "16px",
        cursor: loading ? "not-allowed" : "pointer",
        backgroundColor: loading ? "#f0f0f0" : "white",
        opacity: loading ? 0.7 : 1,
        transition: "opacity 0.1s ease",
      }}
      onClick={handleIncrement}
      disabled={loading}
    >
      Count: {count}
    </button>
  );
}
