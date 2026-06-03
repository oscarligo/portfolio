<script lang="ts">
	import { locale, _ } from 'svelte-i18n';

	const supportedLanguages = [
		{ code: 'es', label: 'Español' },
		{ code: 'en', label: 'English' }
	];

	let currentLang = $derived(
		supportedLanguages.find((language) => language.code === $locale) ?? supportedLanguages[0]
	);

	function handleLanguageChange(event: Event) {
		const select = event.currentTarget as HTMLSelectElement;
		const nextLocale = select.value;

		$locale = nextLocale;
		window.localStorage.setItem('locale', nextLocale);
	}
</script>

<label class="lang-selector">
	<span class="lang-label">{$_('nav.language')}</span>

	<select
		value={$locale}
		onchange={handleLanguageChange}
		aria-label={$_('nav.changeLanguage')}
		title={currentLang.label}
	>
		{#each supportedLanguages as language (language.code)}
			<option value={language.code}>{language.label}</option>
		{/each}
	</select>
</label>

<style>
	.lang-selector {
		display: inline-flex;
		align-items: center;
		gap: 0.65rem;
		color: var(--text-secondary);
		font-size: 0.95rem;
		font-weight: 600;
	}

	.lang-label {
		white-space: nowrap;
	}

	select {
		border: 1px solid var(--border-color);
		border-radius: 999px;
		padding: 0.55rem 0.9rem;
		background: var(--bg-primary);
		color: var(--text-primary);
		font: inherit;
		cursor: pointer;
	}

	@media (max-width: 640px) {
		.lang-selector {
			justify-content: space-between;
			width: 100%;
		}
	}
</style>
