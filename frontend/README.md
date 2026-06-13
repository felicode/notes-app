# Notes App — Frontend

React + TypeScript single-page app built
with [Vite](https://vite.dev/).
This is the web client for the Go
notes-app backend (see the repo root).

## Tech stack
- React 19
- TypeScript
- Vite (dev server + build)
- ESLint

## Prerequisites
- Node.js (version pinned in the repo-root `.tool-versions`, managed via [asdf](https://asdf-vm.com/))

## Setup
```bash
npm install
```

Development

Start the Vite dev server with hot-module
reload:
npm run dev
The app runs at http://localhost:5173 by
default.

Build

Type-check and produce a production build
in dist/:
npm run build

Preview
Serve the production build locally to
verify it:
npm run preview

Lint
npm run lint

Project structure

frontend/
├── public/          # static assets
served as-is
├── src/
│   ├── assets/      # imported
images/svgs
│   ├── App.tsx      # root component
│   ├── main.tsx     # app entry point
│   └── *.css        # styles
├── index.html       # Vite HTML entry
├── vite.config.ts   # Vite configuration
└── tsconfig*.json   # TypeScript
configuration

## Verification
- `cat frontend/README.md` shows the new
content.
- Commands listed match
`frontend/package.json` scripts (`dev`,
`build`, `lint`, `preview`).