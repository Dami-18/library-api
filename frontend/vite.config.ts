import { defineConfig } from 'vite'
import tailwindcss from '@tailwindcss/vite'
import react from '@vitejs/plugin-react'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    react(),
    tailwindcss(),
  ],
  // server: {
  //   proxy: {
  //     '/lib-api': {
  //       target: 'http://localhost:8080', //  Go backend
  //       changeOrigin: true,
  //       secure: false,
  //     },
  //   },
  // },
})
