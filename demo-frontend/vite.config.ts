import { defineConfig } from "vite"
import react from "@vitejs/plugin-react"

export default defineConfig({
  plugins: [react()],
  server: {
    proxy: {
      "/weather": {
        target: "https://d5db5ut10n51lpusmbjr.uvah0e6r.apigw.yandexcloud.net",
        changeOrigin: true,
      },
      "/api": {
        target: "https://d5db5ut10n51lpusmbjr.uvah0e6r.apigw.yandexcloud.net",
        changeOrigin: true
      }
    }
  },
  build: {
    cssCodeSplit: false,
    rollupOptions: {
      output: {
        entryFileNames: "app.js",
        chunkFileNames: "app.js",
        assetFileNames: (assetInfo) => {
          if (assetInfo.name?.endsWith(".css")) {
            return "app.css"
          }

          return "[name][extname]"
        }
      }
    }
  }
})
