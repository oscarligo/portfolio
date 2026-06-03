<script lang="ts">
	import { _ } from 'svelte-i18n';
	import ProjectsSection from '../../components/ProjectsSection/ProjectsSection.svelte';
	import './Home.css';
</script>

<section class="home-page">
	<section class="page-card intro-card">
		<p class="eyebrow">{$_('pages.home.eyebrow')}</p>
		<h1>{$_('pages.home.title')}</h1>
		<p>{$_('pages.home.message')}</p>
	</section>

	<ProjectsSection
		eyebrow={$_('pages.home.projects.eyebrow')}
		title={$_('pages.home.projects.title')}
		description={$_('pages.home.projects.description')}
	/>
</section>
