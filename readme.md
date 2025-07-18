# Templ Demo

## Development Commands

```bash
# Install dependencies
npm install

# Build Tailwind CSS
npm run build-css

# Build React components
npm run build

# Generate templ Go files
templ generate

# Run the server
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
├── go.mod
├── go.sum
├── package.json
└── vite.config.js (or keep esbuild setup)
```
