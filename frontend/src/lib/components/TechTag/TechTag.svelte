<script lang="ts">
	import './TechTag.css';
	let { label, title = label } = $props<{
		label: string;
		title?: string;
	}>();
</script>

<span class="tech-tag" {title}>{label}</span>

