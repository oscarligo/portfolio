<script lang="ts">
	import './ContactTag.css';

	let { type, label, value } = $props<{
		type: 'email' | 'github' | 'linkedin';
		label: string;
		value: string;
	}>();

	const contactKey = $derived(type.trim().toLowerCase());
	const href = $derived(
		contactKey === 'email'
			? `mailto:${value}`
			: value.startsWith('http')
				? value
				: `https://${value}`
	);
</script>

<a
	{href}
	class="contact-tag"
	data-contact={contactKey}
	target={contactKey === 'email' ? '_self' : '_blank'}
	rel="noopener noreferrer"
>
	<span class="contact-icon" aria-hidden="true"></span>
	<span>{label}</span>
</a>
