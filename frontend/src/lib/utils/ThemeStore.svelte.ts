class ThemeStore {
	#isDark = $state(false);
	#systemThemeQuery: MediaQueryList | null = null;
	#boundSystemThemeListener: ((event: MediaQueryListEvent) => void) | null = null;

	get isDark() {
		return this.#isDark;
	}

	initTheme() {
		if (typeof window === 'undefined') {
			return;
		}

		const savedTheme = window.localStorage.getItem('theme');
		this.#systemThemeQuery = window.matchMedia('(prefers-color-scheme: dark)');
		this.#isDark = savedTheme ? savedTheme === 'dark' : this.#systemThemeQuery.matches;

		this.#applyTheme();
		this.#setupSystemThemeListener();
	}

	toggleTheme() {
		this.#isDark = !this.#isDark;

		if (typeof window !== 'undefined') {
			window.localStorage.setItem('theme', this.#isDark ? 'dark' : 'light');
		}

		this.#applyTheme();
	}

	#setupSystemThemeListener() {
		if (!this.#systemThemeQuery) {
			return;
		}

		if (this.#boundSystemThemeListener) {
			this.#systemThemeQuery.removeEventListener('change', this.#boundSystemThemeListener);
		}

		this.#boundSystemThemeListener = (event: MediaQueryListEvent) => {
			if (window.localStorage.getItem('theme')) {
				return;
			}

			this.#isDark = event.matches;
			this.#applyTheme();
		};

		this.#systemThemeQuery.addEventListener('change', this.#boundSystemThemeListener);
	}

	#applyTheme() {
		if (typeof document === 'undefined') {
			return;
		}

		document.documentElement.classList.toggle('dark', this.#isDark);
		this.#applyFavicon();
	}

	#applyFavicon() {
		const faviconEl = document.getElementById('app-favicon');

		if (!(faviconEl instanceof HTMLLinkElement)) {
			return;
		}

		const nextHref = this.#isDark ? faviconEl.dataset.darkHref : faviconEl.dataset.lightHref;

		if (nextHref) {
			faviconEl.href = nextHref;
		}
	}
}

export const theme = new ThemeStore();
