# demo-frontend

Compact React + TypeScript frontend for the serverless workshop.

## Files

This demo intentionally uses a flat structure (no nested folders):

- `index.html`
- `app.tsx`
- `styles.css`
- `package.json`
- `tsconfig.json`

After build, runtime output is also minimal:

- `index.html`
- `app.js`
- `app.css`

## Run

```bash
npm install
npm run build
```

Optional while editing:

```bash
npm run dev
```

Preview built files:

```bash
npm run preview
```

Build output is in `dist/`.

## API wiring

No manual API settings in UI.

Vite proxy is preconfigured:

- `/weather` -> Yandex Cloud Function
- `/api/*` -> API Gateway

This keeps the demo clean and avoids CORS issues in local development.

## What is included in the UI

- Weather by city (`GET /weather?city=...`)
- Add to favorites (`POST /api/favorites`)
- List favorites (`GET /api/favorites`)
- NES.css visual style focused on weather card + favorites board
