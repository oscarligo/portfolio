import { addMessages, init } from 'svelte-i18n';
import en from './en.json' with { type: 'json' };
import es from './es.json' with { type: 'json' };

addMessages('es', es);
addMessages('en', en);

function normalizeLocale(locale: string | null | undefined) {
	const value = locale?.toLowerCase() ?? '';

	if (value.startsWith('en')) {
		return 'en';
	}

	return 'es';
}

const initialLocale =
	typeof window === 'undefined'
		? 'es'
		: normalizeLocale(window.localStorage.getItem('locale') ?? window.navigator.language);

init({
	fallbackLocale: 'es',
	initialLocale
});
