import React from "react";
import { createRoot } from "react-dom/client";
import Counter from "./components/Counter.jsx";
import { Container } from "./components/desk/Container.tsx";

// Export function to render the Counter component
export function renderCounter(element) {
  const root = createRoot(element);
  root.render(React.createElement(Counter));
}

// Export function to render the Desk component
export function renderDesk(element) {
  const root = createRoot(element);
  root.render(React.createElement(Container));
}
