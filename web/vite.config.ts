import tailwindcss from "@tailwindcss/vite";
import {sveltekit} from "@sveltejs/kit/vite";
import {defineConfig} from "vite";

export default defineConfig({
    plugins: [tailwindcss(), sveltekit()],
    resolve: process.env.VITEST
        ? {
            conditions: ["browser"]
        }
        : undefined,
    test: {
        expect: {requireAssertions: true},
        projects: [
            {
                extends: "./vite.config.ts",
                test: {
                    name: "client",
                    environment: "jsdom",
                    include: ["src/**/*.svelte.{test,spec}.{js,ts}"],
                    exclude: ["src/lib/server/**"],
                    setupFiles: ["./vitest-setup-client.ts"]
                }
            },
        ]
    }
});
