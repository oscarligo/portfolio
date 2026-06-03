<script lang="ts">
	import './TechTag.css';

	let { label, icon_slug = label.toLowerCase() } = $props<{
		label: string;
		icon_slug?: string;
	}>();

	const techKey = $derived(icon_slug.trim().toLowerCase());
</script>

<div class="tech-tag" data-tech={techKey}>
	{#if icon_slug}
		<img src="/logos/{icon_slug}.svg" alt="Icono de {label}" class="tech-icon" />
	{/if}
	<span>{label}</span>
</div>
