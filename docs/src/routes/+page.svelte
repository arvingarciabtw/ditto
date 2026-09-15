<script>
	import Keyboard from "$lib/components/keyboard.svelte";
</script>

<main>
	<section class="keyboard">
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

	<section class="docs">
		<p class="subtext">
			to get started, see the
			<a href="/motivation">docs</a>
			if it fits your use case or if you'd like to mess around and have fun with it.
		</p>
	</section>

	<section class="docs">
		<p class="subtext">
			see the
			<a
				href="https://github.com/arvingarciabtw/ditto/releases"
				target="_blank"
				rel="noopener noreferrer">releases</a
			>
			for the installation. if you're on linux, there are other
			<a href="/installation">options</a>.
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
	section.keyboard {
		margin-top: 2rem;
		margin-bottom: 4rem;
	}

	@media (max-width: 850px) {
		section.keyboard {
			margin-top: 1rem;
		}
	}

	@media (max-width: 650px) {
		section.keyboard {
			display: none;
		}
		main {
			place-content: center;
		}
	}
</style>
