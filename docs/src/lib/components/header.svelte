<script lang="ts">
	import Version from "@lucide/svelte/icons/git-fork";
	import Star from "@lucide/svelte/icons/star";
	import Menu from "@lucide/svelte/icons/menu";

	let showMobileNav = $state(false);

	function handleClick() {
		showMobileNav = !showMobileNav;
	}
</script>

<header>
	<div class="left">
		<a class="name" href="/">ditto</a>
		<a
			href="https://github.com/arvingarciabtw/ditto"
			target="_blank"
			rel="external noopener noreferrer"
		>
			<span class="version">
				<Version size={14} />
				<p class="label">v1.3.3</p>
			</span>
		</a>
		<a
			href="https://github.com/arvingarciabtw/ditto"
			target="_blank"
			rel="external noopener noreferrer"
		>
			<span class="stars">
				<Star size={14} />
				<p class="label">120+</p>
			</span>
		</a>
	</div>
	<div class="right">
		<button class="icon" onclick={handleClick}>
			<Menu size={20} />
		</button>
	</div>
	<div class="sidebar-mobile-nav" class:show-up={showMobileNav}>
		<div class="top">
			<p>page navigation</p>
			<button onclick={handleClick}>close</button>
		</div>
		<div class="nav-sections">
			<section class="getting-started">
				<p>getting started</p>
				<ul>
					<li><a href="/motivation" onclick={handleClick}>motivation</a></li>
					<li><a href="/installation" onclick={handleClick}>installation</a></li>
					<li><a href="/permissions" onclick={handleClick}>permissions</a></li>
					<li><a href="/config" onclick={handleClick}>config</a></li>
				</ul>
			</section>
			<section class="usage">
				<p>usage</p>
				<ul>
					<li><a href="/modes" onclick={handleClick}>modes</a></li>
					<li><a href="/visual" onclick={handleClick}>visual</a></li>
					<li><a href="/lists" onclick={handleClick}>lists</a></li>
					<li><a href="/lock" onclick={handleClick}>lock</a></li>
					<li><a href="/custom-layouts" onclick={handleClick}>custom layouts</a></li>
				</ul>
			</section>
		</div>
	</div>
</header>

<style>
	header {
		display: flex;
		align-items: center;
		justify-content: space-between;
	}
	.left {
		display: grid;
		grid-template-columns: repeat(3, max-content);
		align-items: center;
		gap: 2rem;
	}
	.right {
		display: none;
	}
	.sidebar-mobile-nav {
		position: fixed;
		inset: 0;
		padding: 1.5rem;
		z-index: 10;
		transform: translateX(100%);
		transition: transform 300ms ease;
		background: var(--crust);

		.top {
			margin-bottom: 2rem;
			display: flex;
			align-items: center;
			justify-content: space-between;
			color: var(--mauve);
		}
		.top button {
			padding: 0;
			background: transparent;
			border: none;
			color: var(--red);
		}

		.nav-sections {
			section {
				margin-bottom: 4rem;
			}
			p {
				color: var(--teal);
			}
			ul {
				margin-top: 2rem;
				display: grid;
				gap: 0.5rem;
			}
			a {
				color: var(--text);
			}
		}
	}
	.sidebar-mobile-nav.show-up {
		transform: translateX(0);
	}

	button.icon {
		padding: 0;
		background: transparent;
		border: none;
		color: var(--sapphire);
	}

	.name:hover {
		text-decoration: underline dotted;
	}
	.version,
	.stars {
		display: grid;
		grid-template-columns: repeat(2, max-content);
		align-items: center;
		gap: 0.75rem;
	}
	.version {
		color: var(--green);
	}
	.stars {
		color: var(--yellow);
	}
	.version:hover,
	.stars:hover {
		text-decoration: underline dotted;
	}

	a {
		color: var(--mauve);
		text-decoration: none;
	}

	@media (max-width: 850px) {
		.right {
			display: block;
		}
	}
</style>
