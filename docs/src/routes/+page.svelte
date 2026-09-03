<script>
	import Keyboard from "$lib/components/keyboard.svelte";
</script>

<main>
	<section class="keyboard">
		<p class="subtext">try typing!</p>
		<Keyboard />
	</section>

	<section class="description">
		<p class="subtext">
			ditto is a cross-platform, system-wide ascii keyboard visualizer and keycaster that mirrors
			your live keyboard inputs in real time, even when the terminal isn't in focus. it
			automatically syncs with your native terminal color scheme for seamless, interactive eye
			candy.
		</p>
	</section>

	<section class="installation">
		<p class="subtext">
			see
			<a
				href="https://github.com/arvingarciabtw/ditto/releases"
				target="_blank"
				rel="noopener noreferrer">releases</a
			>
			for the installation.
		</p>
	</section>

	<section class="docs">
		<p class="subtext">
			see the
			<a href="/installation">docs</a>
			to get started.
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
	section.keyboard p {
		margin-bottom: 2rem;
	}
</style>
