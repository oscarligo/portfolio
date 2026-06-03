<script lang="ts">
	import '../app.css';
	import '../i18n/index.ts';
	import Navbar from '../lib/components/Navbar/Navbar.svelte';
	import { _ } from 'svelte-i18n'

	let { children } = $props();
</script>

<nav>
	<div class="app-shell">
		<Navbar />

		<main class="main-content">
			{@render children()}
		</main>
	</div>

	<footer class="footer">
		<p>{$_('footer.text')}
			<a href="https://github.com/oscarligo/portfolio.git" target="_blank" rel="noopener noreferrer">GitHub</a>
		</p>
	</footer>	
</nav>