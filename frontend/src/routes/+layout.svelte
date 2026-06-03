<script lang="ts">
	import '../app.css';
	import '../i18n/index.ts';
	import Navbar from '../lib/components/Navbar/Navbar.svelte';

	let { children } = $props();
</script>

<nav>
	<div class="app-shell">
		<Navbar />

		<main class="main-content">
			{@render children()}
		</main>
	</div>
</nav>