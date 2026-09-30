<script>
	import Keyboard from "$lib/components/keyboard.svelte";
	import Star from "@lucide/svelte/icons/star";
</script>

<main>
	<section class="keyboard">
		<Keyboard />
	</section>

	<section class="keyboard-disclosure">
		<blockquote>
			on larger viewports, this page shows an interactive keyboard. see the site on desktop if you
			wanna mess around with it! (⌐■ᴗ■)
		</blockquote>
	</section>

	<section class="description">
		<p class="subtext">
			ditto is a cross-platform and system-wide keyboard visualizer and keycaster in the terminal
			that works even when it isn't in focus. automatically syncs with your native terminal color
			scheme. intended for folks who want to prettify their terminal setup.
		</p>
	</section>

	<section>
		<p class="subtext">
			to get started, see the
			<a class="docs" href="/motivation">motivation</a>
			behind it and see if you find it interesting.
		</p>
	</section>

	<section>
		<p class="subtext">
			if you quickly want to see what it looks like, take a peek at the
			<a class="lists" href="/lists">lists</a>,
			<a class="visual" href="/visual">visual</a>, and
			<a class="modes" href="/modes">modes</a> pages.
		</p>
	</section>

	<section>
		<p class="subtext">
			see the
			<a
				class="releases"
				href="https://github.com/arvingarciabtw/ditto/releases"
				target="_blank"
				rel="noopener noreferrer">releases</a
			>
			to install. if you're on linux, there are other
			<a class="options" href="/installation">installation</a> options.
		</p>
	</section>

	<section>
		<p class="subtext">
			if you like it, a <span class="star"><Star /></span> is appreciated!
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
		margin: 2rem 0;
	}
	section.keyboard-disclosure {
		display: none;
	}
	.docs {
		color: var(--mauve);
	}
	.lists,
	.visual,
	.modes {
		color: var(--sapphire);
	}
	.releases {
		color: var(--green);
	}
	.options {
		color: var(--yellow);
	}
	.star {
		display: inline;
		vertical-align: middle;
		color: var(--yellow);

		& :global(svg) {
			display: inline;
			width: 1em;
			height: 1em;
			fill: var(--yellow);
		}
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
		section.keyboard-disclosure {
			display: block;
			padding-left: 1rem;
			margin: 0rem;
			border-left: 4px solid var(--surface1);
			color: var(--overlay1);
			line-height: 2;
		}
		main {
			place-content: center;
		}
	}
</style>
