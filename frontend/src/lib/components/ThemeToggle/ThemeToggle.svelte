<script lang="ts">
    import { theme } from '../../utils/ThemeStore.svelte.ts';
    import './ThemeToggle.css'
    import Sun from '@lucide/svelte/icons/sun';
	import Moon from '@lucide/svelte/icons/moon';
</script>

<button 
    class="theme-toggle" 
    onclick={() => theme.toggleTheme()} 
    aria-label="Cambiar a modo {theme.isDark ? 'claro' : 'oscuro'}"
    title="Cambiar tema"
>
    {#if !theme.isDark}
        <Sun class="sun-icon" aria-hidden="true" />
    {:else}
        <Moon class="moon-icon" aria-hidden="true" />
    {/if}
</button>
