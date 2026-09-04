<script>
	import Keyboard from "$lib/components/keyboard.svelte";
</script>

<main>
	<section class="docs">
		<p class="subtext">
			see the
			<a href="/motivation">docs</a>
			to get started. see
			<a
				href="https://github.com/arvingarciabtw/ditto/releases"
				target="_blank"
				rel="noopener noreferrer">releases</a
			>
			for the installation.
		</p>
	</section>

	<section class="keyboard">
		<Keyboard />
		<blockquote>this keyboard updates... try typing! •ᴗ•</blockquote>
	</section>

	<section class="description">
		<p class="subtext">
			ditto is a cross-platform, system-wide ascii keyboard visualizer and keycaster that mirrors
			your live keyboard inputs in real time, even when the terminal isn't in focus. it
			automatically syncs with your native terminal color scheme for seamless, interactive eye
			candy.
		</p>
	</section>

	<section class="motivation">
		<p class="subtext">
			see the
			<a href="/motivation">motivation</a>
			behind it!
		</p>
	</section>
</main>

<style>
	main {
		display: grid;
		grid-auto-rows: max-content;
		gap: 2rem;
	}
	section {
		max-width: 60rem;
	}
	section.keyboard blockquote {
		padding-left: 1rem;
		margin-top: 2rem;
		border-left: 4px solid var(--surface1);
		color: var(--overlay1);
	}
</style>
