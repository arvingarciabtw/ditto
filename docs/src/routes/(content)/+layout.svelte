<script>
	import Sidebar from "$lib/components/sidebar.svelte";

	let { children } = $props();
</script>

<main>
	<Sidebar />
	<article>
		{@render children()}
	</article>
</main>

<style>
	main {
		display: grid;
		grid-template-columns: 16rem 1fr;
		gap: 2rem;
	}

	@media (max-width: 850px) {
		main {
			grid-template-columns: 1fr;
		}
	}

	/* to style our markdown */
	:global(article) {
		min-width: 0;
	}
	:global(article p) {
		margin-top: 1rem;
		color: var(--subtext0);
	}
	:global(article del) {
		color: var(--overlay1);
	}
	:global(article img) {
		margin: 2rem 0;
		border: 1px dashed var(--base);
	}
	:global(article h1) {
		font-size: 1rem;
		font-weight: 400;
		color: var(--flamingo);
	}
	:global(article h1::before) {
		content: "# ";
	}
	:global(article h2) {
		margin-top: 4rem;
		font-size: 1rem;
		font-weight: 400;
		color: var(--pink);
	}
	:global(article h2::before) {
		content: "## ";
	}
	:global(article code) {
		padding: 0.125rem 0.25rem 0.175rem;
		background: var(--base);
		color: var(--text);
		border: 1px dashed var(--surface0);
	}
	:global(article pre) {
		margin: 1rem 0 2rem;
		padding: 1.5rem;
		background: var(--mantle);
		overflow-x: auto;
		border: 1px dashed var(--surface0);
	}
	:global(article pre code) {
		border: none;
	}
	:global(article table) {
		width: 100%;
		margin: 1rem 0 2rem;
		border: 1px dashed var(--surface0);
		border-collapse: collapse;
		background: var(--mantle);
	}
	:global(article table code) {
		background: transparent;
		border: none;
		color: var(--subtext0);
	}
	:global(article th),
	:global(article td) {
		padding: 0.75rem 1rem;
		text-align: left;
		vertical-align: top;
		border-right: 1px dashed var(--surface0);
		border-bottom: 1px dashed var(--surface0);
	}
	:global(article th) {
		background: var(--base);
		color: var(--sapphire);
		font-weight: 400;
	}
	:global(article td) {
		color: var(--subtext0);
	}
	:global(article tbody tr:nth-child(even)) {
		background: var(--base);
	}
	:global(article tr:last-child td) {
		border-bottom: 0;
	}
	:global(article th:last-child),
	:global(article td:last-child) {
		border-right: 0;
	}
	:global(article blockquote) {
		margin-top: 1rem;
		padding-left: 1rem;
		border-left: 4px solid var(--surface0);
	}

	@media (max-width: 500px) {
		:global(article th),
		:global(article td) {
			padding: 0.625rem 0.75rem;
		}
	}
</style>
