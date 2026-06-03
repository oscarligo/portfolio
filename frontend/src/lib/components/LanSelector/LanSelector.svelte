<script lang="ts">
	import { locale, _ } from 'svelte-i18n';
	import './LanSelector.css';

	const supportedLanguages = [
		{ code: 'es', label: 'ES' },
		{ code: 'en', label: 'EN' }
	] as const;

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
