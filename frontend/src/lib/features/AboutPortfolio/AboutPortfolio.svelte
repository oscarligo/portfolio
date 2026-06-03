<script lang="ts">
	import { _ } from 'svelte-i18n';
	import './AboutPortfolio.css';

	const frontendItems = ['framework', 'language', 'i18n', 'tooling'] as const;
	const backendItems = ['api', 'persistence', 'auth', 'storage'] as const;
	const architecturePoints = ['contract', 'data', 'delivery'] as const;
</script>

<section class="about-portfolio">
	<section class="page-card intro-card">
		<p class="eyebrow">{$_('pages.aboutPortfolio.eyebrow')}</p>
		<h1>{$_('pages.aboutPortfolio.title')}</h1>
		<p>{$_('pages.aboutPortfolio.message')}</p>
	</section>

	<div class="stack-grid">
		<article class="stack-card">
			<p class="stack-label">{$_('pages.aboutPortfolio.frontend.label')}</p>
			<h2>{$_('pages.aboutPortfolio.frontend.title')}</h2>
			<p>{$_('pages.aboutPortfolio.frontend.description')}</p>

			<div class="stack-items">
				{#each frontendItems as item}
					<div class="stack-item">
						<h3>{$_(`pages.aboutPortfolio.frontend.items.${item}.title`)}</h3>
						<p>{$_(`pages.aboutPortfolio.frontend.items.${item}.description`)}</p>
					</div>
				{/each}
			</div>
		</article>

		<article class="stack-card">
			<p class="stack-label">{$_('pages.aboutPortfolio.backend.label')}</p>
			<h2>{$_('pages.aboutPortfolio.backend.title')}</h2>
			<p>{$_('pages.aboutPortfolio.backend.description')}</p>

			<div class="stack-items">
				{#each backendItems as item}
					<div class="stack-item">
						<h3>{$_(`pages.aboutPortfolio.backend.items.${item}.title`)}</h3>
						<p>{$_(`pages.aboutPortfolio.backend.items.${item}.description`)}</p>
					</div>
				{/each}
			</div>
		</article>
	</div>

	<section class="page-card architecture-card">
		<p class="stack-label">{$_('pages.aboutPortfolio.architecture.label')}</p>
		<h2>{$_('pages.aboutPortfolio.architecture.title')}</h2>
		<p>{$_('pages.aboutPortfolio.architecture.description')}</p>

		<div class="architecture-list">
			{#each architecturePoints as point}
				<div class="architecture-point">
					<h3>{$_(`pages.aboutPortfolio.architecture.points.${point}.title`)}</h3>
					<p>{$_(`pages.aboutPortfolio.architecture.points.${point}.description`)}</p>
				</div>
			{/each}
		</div>
	</section>
</section>
