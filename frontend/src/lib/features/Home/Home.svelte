<script lang="ts">
	import { _ } from 'svelte-i18n';
	import ProjectsSection from '../../components/ProjectsSection/ProjectsSection.svelte';
	import './Home.css';
	import AboutMeCard from './components/AboutMeCard/AboutMeCard.svelte';
</script>

<section class="home-page">
	<AboutMeCard />

	<ProjectsSection
		eyebrow={$_('pages.home.projects.eyebrow')}
		title={$_('pages.home.projects.title')}
		description={$_('pages.home.projects.description')}
	/>
</section>
