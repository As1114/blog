import {fileURLToPath, URL} from 'node:url'
import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
// https://vitejs.dev/config/
export default defineConfig(({mode}) => {
    const serverUrl = 'http://127.0.0.1:8888'
    return {
        plugins: [
            vue(),
        ],
        envDir: "./",
        resolve: {
            alias: {
                '@': fileURLToPath(new URL('./src', import.meta.url))
            }
        },
        server: {
            host: "0.0.0.0",
            port: 81,
            proxy: {
                "/api": {
                    target: serverUrl,
                    changeOrigin: true,
                },
                "/uploads": {
                    target: serverUrl,
                    changeOrigin: true,
                },
            }
        }
    }
})
