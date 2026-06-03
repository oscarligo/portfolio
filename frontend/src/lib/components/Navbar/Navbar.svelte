<script lang="ts">
	// Importamos las llaves compiladas directamente
	import * as m from '$paraglide/messages.js';
	import { languageTag, setLanguageTag } from '$paraglide/runtime.js';

	let { currentRoute, onNavigate } = $props();

	function toggleLanguage() {
		const nextLang = languageTag() === 'es' ? 'en' : 'es';
		setLanguageTag(nextLang);
	}
</script>

<nav class="navbar">
	<h1 onclick={() => onNavigate('home')}>Oscar Rompich</h1>
	
	<div class="nav-links">
		<button onclick={() => onNavigate('home')}>{m.menu_home()}</button>
		<button onclick={() => onNavigate('about')}>{m.menu_about()}</button>
		
		<button onclick={toggleLanguage}>
			{languageTag() === 'es' ? '🇺🇸 EN' : '🇪🇸 ES'}
		</button>
	</div>
</nav>