<script lang="ts">
	import '../app.css';
	import '../i18n/index.ts';
	import Navbar from '../lib/components/Navbar/Navbar.svelte';
	import { _ } from 'svelte-i18n'
	import { theme } from '../lib/utils/ThemeStore.svelte.ts';
	import { onMount } from 'svelte';

	let { children } = $props();
	onMount(() => {
        theme.initTheme(); 
    });
</script>

<svelte:head>
    <title>{$_('header.title')}</title>
    <meta name="description" content={$_('header.description')} />
</svelte:head>

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