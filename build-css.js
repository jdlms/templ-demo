const fs = require("fs");
const path = require("path");
const tailwindcss = require("@tailwindcss/postcss");
const postcss = require("postcss");

async function buildCSS() {
  const inputCSS = fs.readFileSync("./input.css", "utf8");

  const result = await postcss([tailwindcss]).process(
    inputCSS,
    { from: "./input.css", to: "./web/static/assets/styles.css" }
  );

  fs.writeFileSync("./web/static/assets/styles.css", result.css);
  console.log("✅ Tailwind CSS built successfully");
}

buildCSS().catch(console.error);
