# Templ Demo

An experiment to use React in Go Templ templates with Tailwind CSS v4.

## Setup & Build Commands

```bash
# 1. Install dependencies
npm install

# 2. Build Tailwind CSS
npm run build-css

# 3. Build React components
npm run build

# 4. Generate templ Go files from templates
templ generate

# 5. Run the server
go run cmd/server/main.go
```

## Development Commands

```bash
# Watch and rebuild CSS on changes
npm run dev-css

# Watch and rebuild React components on changes
npm run dev

# Watch and regenerate templ files on changes
templ generate --watch

# Run all in separate terminals for full development setup
```

## Complete Build Process

To build the entire project from scratch:

```bash
npm install
npm run build-css
npm run build
templ generate
go run cmd/server/main.go
```

## Project Structure

```
templ-demo/
├── cmd/
│ └── server/
│ └── main.go
├── internal/
│ └── handlers/
│ └── handlers.go
├── web/
│ ├── templates/
│ │ ├── main.templ
│ │ └── hello.templ
│ ├── components/
│ │ ├── Counter.jsx
│ │ └── desk/
│ │ ├── Container.tsx
│ │ ├── Loader.tsx
│ │ └── Spinner.tsx
│ ├── static/
│ │ ├── models/
│ │ │ └── Desk.glb
│ │ └── dist/
│ │ └── (built files)
│ └── index.js
├── input.css
├── postcss.config.js
├── go.mod
├── go.sum
├── package.json
└── vite.config.js
```
