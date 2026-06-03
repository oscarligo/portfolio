import { init, register } from 'svelte-i18n';

register('es', () => import('./es.json'));
register('en', () => import('./en.json'));

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
