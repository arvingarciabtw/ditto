<script>
	import Version from "@lucide/svelte/icons/git-fork";
	import Star from "@lucide/svelte/icons/star";
</script>

<header>
	<a class="name" href="/">ditto</a>
	<span class="version">
		<div class="icon">
			<Version />
		</div>
		<p class="label">v1.3.3</p>
	</span>
	<span class="stars">
		<div class="icon">
			<Star />
		</div>
		<p class="label">120+</p>
	</span>
</header>

<style>
	header {
		display: grid;
		grid-template-columns: repeat(3, max-content);
		align-items: center;
		gap: 2rem;
	}

	.version,
	.stars {
		display: grid;
		grid-template-columns: repeat(2, max-content);
		gap: 0.75rem;
		.icon {
			width: 14px;
			height: 14px;
		}
	}
	.version {
		color: var(--green);
	}
	.stars {
		color: var(--yellow);
	}

	a {
		color: var(--mauve);
		text-decoration: none;
	}
</style>
