<aside>
	<section class="getting-started">
		<p class="label">getting started</p>
		<ul>
			<li><a href="/motivation">motivation</a></li>
			<li><a href="/installation">installation</a></li>
			<li><a href="/permissions">permissions</a></li>
			<li><a href="/config">config</a></li>
		</ul>
	</section>
	<section class="usage">
		<p class="label">usage</p>
		<ul>
			<li><a href="/modes">modes</a></li>
			<li><a href="/visual">visual</a></li>
			<li><a href="/lists">lists</a></li>
			<li><a href="/lock">lock</a></li>
			<li><a href="/custom-layouts">custom layouts</a></li>
		</ul>
	</section>
</aside>

<style>
	section {
		margin-bottom: 4rem;
	}
	a {
		color: var(--text);
		text-decoration: none;
	}
	a:hover {
		text-decoration: underline dotted;
	}
	ul {
		margin-top: 2rem;
		display: grid;
		gap: 0.75rem;
	}
	p.label {
		color: var(--teal);
	}

	@media (max-width: 850px) {
		aside {
			display: none;
		}
	}
</style>
