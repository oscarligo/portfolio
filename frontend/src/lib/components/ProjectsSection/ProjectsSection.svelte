<script lang="ts">
	import { onMount } from 'svelte';
	import { _ } from 'svelte-i18n';
	import ProjectCard from '../ProjectCard/ProjectCard.svelte';
	import { projectService } from '../../services/projectService.ts';
	import type { Project } from '../../types/project.js';
	import './ProjectSection.css';

	let {
		eyebrow = '',
		title,
		description,
		featuredOnly = false,
		limit
	} = $props<{
		eyebrow?: string;
		title: string;
		description: string;
		featuredOnly?: boolean;
		limit?: number;
	}>();

	let projects = $state<Project[]>([]);
	let isLoading = $state(true);
	let errorMessage = $state('');

	async function loadProjects() {
		isLoading = true;
		errorMessage = '';

		try {
			const response = featuredOnly
				? await projectService.getFeatured()
				: await projectService.getAll();

			projects = typeof limit === 'number' ? response.slice(0, limit) : response;
		} catch (error) {
			projects = [];
			errorMessage =
				error instanceof Error ? error.message : $_('projects.unexpectedError');
		} finally {
			isLoading = false;
		}
	}

	onMount(() => {
		void loadProjects();
	});
</script>

<section class="projects-section">
	<div class="section-heading">
		{#if eyebrow}
			<p class="eyebrow">{eyebrow}</p>
		{/if}

		<h2>{title}</h2>
		<p>{description}</p>
	</div>

	{#if isLoading}
		<div class="state-card">
			<p>{$_('projects.loading')}</p>
		</div>
	{:else if errorMessage}
		<div class="state-card state-error">
			<p class="state-title">{$_('projects.errorTitle')}</p>
			<p class="state-detail">{errorMessage}</p>
			<button type="button" onclick={() => void loadProjects()}>
				{$_('projects.retry')}
			</button>
		</div>
	{:else if projects.length === 0}
		<div class="state-card">
			<p>{$_('projects.empty')}</p>
		</div>
	{:else}
		<div class="projects-grid">
			{#each projects as project (project.id)}
				<ProjectCard {project} />
			{/each}
		</div>
	{/if}
</section>
