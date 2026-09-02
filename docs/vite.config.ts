import { escapeSvelte, mdsvex } from "mdsvex";
import adapter from "@sveltejs/adapter-static";
import { sveltekit } from "@sveltejs/kit/vite";
import { codeToHtml, type BundledLanguage } from "shiki";
import { defineConfig } from "vite";

export default defineConfig({
	plugins: [
		sveltekit({
			compilerOptions: {
				// Force runes mode for the project, except for libraries. Can be removed in svelte 6.
				runes: ({ filename }) =>
					filename.split(/[/\\]/).includes("node_modules") ? undefined : true
			},
			adapter: adapter(),
			preprocess: [
				mdsvex({
					extensions: [".svx", ".md"],
					highlight: {
						highlighter: async (code, lang) => {
							const html = await codeToHtml(code, {
								lang: (lang ?? "text") as BundledLanguage,
								theme: "catppuccin-mocha"
							});

							return `{@html \`${escapeSvelte(html)}\`}`;
						}
					}
				})
			],
			extensions: [".svelte", ".svx", ".md"]
		})
	]
});
