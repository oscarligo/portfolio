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
        class:open={isMenuOpen}
        onclick={toggleMenu} 
        aria-label="Toggle menu"
        aria-expanded={isMenuOpen}
    >
        <span class="menu-toggle-icon" aria-hidden="true">
            <span></span>
            <span></span>
            <span></span>
        </span>
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
