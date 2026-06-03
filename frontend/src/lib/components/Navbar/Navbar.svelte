<script lang="ts">
    import { page } from '$app/state';
    import { _ } from 'svelte-i18n';
    import LanSelector from '../LanSelector/LanSelector.svelte';
    import './Navbar.css';

    let isMenuOpen = $state(false);

    const links = [
        { href: '/about', key: 'nav.aboutPortfolio' }
    ] as const;

    function isActive(href: string) {
        return page.url.pathname === href;
    }

    function toggleMenu() {
        isMenuOpen = !isMenuOpen;
    }

    function closeMenu() {
        isMenuOpen = false;
    }
</script>

<nav class="navbar" aria-label={$_('nav.ariaLabel')}>
    <a class="brand" href="/" onclick={closeMenu}>
        {$_('nav.brand')}
    </a>

    <button 
        class="menu-toggle" 
        onclick={toggleMenu} 
        aria-label="Toggle menu"
        aria-expanded={isMenuOpen}
    >
        {#if isMenuOpen}
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M18 6 6 18"/><path d="m6 6 12 12"/></svg>
        {:else}
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="4" x2="20" y1="12" y2="12"/><line x1="4" x2="20" y1="6" y2="6"/><line x1="4" x2="20" y1="18" y2="18"/></svg>
        {/if}
    </button>

    <div class="navbar-actions" class:show-mobile={isMenuOpen}>
        <div class="nav-links">
            {#each links as link (link.href)}
                <a 
                    class:active={isActive(link.href)} 
                    class="nav-link" 
                    href={link.href}
                    onclick={closeMenu}
                >
                    {$_(link.key)}
                </a>
            {/each}
        </div>

        <LanSelector />
    </div>
</nav>