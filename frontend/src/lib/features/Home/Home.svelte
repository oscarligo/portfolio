<script lang="ts">
	import { _ } from 'svelte-i18n';
	import './Home.css';
</script>

<section class="page-card">
	<p class="eyebrow">Home</p>
	<h1>{$_('pages.home.title')}</h1>
	<p>{$_('pages.home.message')}</p>
</section>
