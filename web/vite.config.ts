import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import VueDevTools from "vite-plugin-vue-devtools";
import { resolve } from "path";

export default defineConfig({
  plugins: [vue(), VueDevTools()],
  resolve: {
    alias: {
      "@": resolve(__dirname, "./src"),
    },
  },
  server: {
    proxy: {
      "/slink": {
        target: "http://localhost:8090",
        changeOrigin: true,
      },
      "/s/": {
        target: "http://localhost:3000",
        changeOrigin: true,
      },
    },
  },
});