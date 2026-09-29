<script lang="ts">
	import "../global.css";
	import favicon from "$lib/assets/favicon.png";
	import Header from "$lib/components/header.svelte";
	import Footer from "$lib/components/footer.svelte";

	let { children } = $props();
</script>

<svelte:head>
	<link rel="icon" href={favicon} />
	<link
		rel="preload"
		href="/commit-mono-vf.woff2"
		as="font"
		type="font/woff2"
		crossorigin="anonymous"
	/>
	<title>ditto</title>
</svelte:head>

<Header />
{@render children()}
<Footer />
