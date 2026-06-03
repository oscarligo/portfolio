<script lang="ts">
	import './i18n/index.ts';
	import Navbar from './lib/components/Navbar/Navbar.svelte';
	import HomePage from './routes/+page.svelte';
</script>

<div class="app-shell">
	<Navbar />

	<main class="main-content">
		<HomePage />
	</main>
</div>


