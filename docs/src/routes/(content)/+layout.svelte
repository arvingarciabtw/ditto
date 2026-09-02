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
		grid-template-columns: 20rem 1fr;
		gap: 2rem;
	}

	/* to style our markdown */
	:global(article p) {
		margin-top: 1rem;
		color: var(--subtext0);
		text-wrap: balance;
	}
	:global(article p:last-of-type) {
		margin-top: 4rem;
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
	:global(article pre) {
		margin: 1rem 0 2rem;
		padding: 1.5rem;
		background: var(--mantle);
		overflow-x: auto;
	}
	:global(article blockquote) {
		margin-top: 1rem;
		padding-left: 1rem;
		border-left: 4px solid var(--surface0);
	}
</style>
