<script lang="ts">
	import { onMount } from 'svelte';
	import { isLoading } from 'svelte-i18n';
	import Navbar from './lib/components/Navbar/Navbar.svelte';
	import AboutPortfolioPage from './routes/AboutPortfolio/+page.svelte';
	import HomePage from './routes/+page.svelte';

	type Route = 'home' | 'about';

	const routeHashes: Record<Route, string> = {
		home: '#/',
		about: '#/about-portfolio'
	};

	let currentRoute = $state<Route>('home');

	function getRouteFromHash(hash: string): Route {
		return hash === routeHashes.about ? 'about' : 'home';
	}

	function syncRoute(hash: string) {
		currentRoute = getRouteFromHash(hash);
	}

	function navigate(route: Route) {
		const nextHash = routeHashes[route];

		if (window.location.hash !== nextHash) {
			window.location.hash = nextHash;
		}

		syncRoute(nextHash);
	}

	onMount(() => {
		const knownHashes = new Set(Object.values(routeHashes));

		if (!knownHashes.has(window.location.hash)) {
			window.location.hash = routeHashes.home;
		}

		syncRoute(window.location.hash);

		const handleHashChange = () => {
			syncRoute(window.location.hash);
		};

		window.addEventListener('hashchange', handleHashChange);

		return () => {
			window.removeEventListener('hashchange', handleHashChange);
		};
	});
</script>

{#if $isLoading}
	<div class="loading-screen">Loading...</div>
{:else}
	<div class="app-shell">
		<Navbar {currentRoute} onNavigate={navigate} />

		<main class="main-content">
			{#if currentRoute === 'home'}
				<HomePage />
			{:else}
				<AboutPortfolioPage />
			{/if}
		</main>
	</div>
{/if}

<style>
	:global(body) {
		margin: 0;
		min-height: 100vh;
		background:
			radial-gradient(circle at top, rgba(14, 165, 233, 0.12), transparent 35%),
			var(--bg-surface);
		color: var(--text-primary);
	}

	.loading-screen {
		display: grid;
		place-items: center;
		min-height: 100vh;
		font-size: 1rem;
		color: var(--text-secondary);
	}

	.app-shell {
		min-height: 100vh;
	}

	.main-content {
		max-width: 960px;
		margin: 0 auto;
		padding: 3rem 1.5rem;
		box-sizing: border-box;
	}
</style>
