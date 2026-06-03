<script lang="ts">
	import { locale, _ } from 'svelte-i18n';
	import type { Project } from '../../types/project.js';
	import TechTag from '../TechTag/TechTag.svelte';
	import './ProjectCard.css';

	let { project } = $props<{
		project: Project;
	}>();

	const coverImage = $derived(project.images[0]);
	const currentLocale = $derived($locale === 'en' ? 'en-US' : 'es-ES');

	function formatDate(value: string, languageTag: string) {
		const date = new Date(value);

		if (Number.isNaN(date.getTime())) {
			return value;
		}

		return new Intl.DateTimeFormat(languageTag, {
			year: 'numeric',
			month: 'short'
		}).format(date);
	}
</script>

<article class="project-card">
	<div class="project-media">
		{#if coverImage?.image_url}
			<img src={coverImage.image_url} alt={`${project.title} preview`} loading="lazy" />
		{:else}
			<div class="project-placeholder">
				<span>{project.title}</span>
			</div>
		{/if}

		{#if project.featured}
			<span class="featured-badge">{$_('projects.featured')}</span>
		{/if}
	</div>

	<div class="project-body">
		<div class="project-header">
			<div>
				<h3>{project.title}</h3>
				<p class="project-meta">
					{$_('projects.published')} {formatDate(project.created_at, currentLocale)}
				</p>
			</div>
		</div>

		{#if project.technologies.length}
			<div class="project-tags" aria-label={$_('projects.technologies')}>
				{#each project.technologies as technology (technology.id)}
					<TechTag label={technology.name} title={technology.icon_slug || technology.name} />
				{/each}
			</div>
		{/if}

		<div class="project-actions">
			{#if project.repo_url}
				<a href={project.repo_url} target="_blank" rel="noreferrer">
					{$_('projects.actions.repository')}
				</a>
			{/if}

			{#if project.live_url}
				<a href={project.live_url} target="_blank" rel="noreferrer">
					{$_('projects.actions.liveDemo')}
				</a>
			{/if}

			{#if project.video_url}
				<a href={project.video_url} target="_blank" rel="noreferrer">
					{$_('projects.actions.video')}
				</a>
			{/if}
		</div>
	</div>
</article>
