<script lang="ts">
	import { _ } from 'svelte-i18n';
	import './AboutMeCard.css';
    import TechTag from '$lib/components/TechTag/TechTag.svelte';
    import type { ProjectTechnology } from '../../../../types/project.ts';
    import ContactTag from '$lib/components/ContactTag/ContactTag.svelte';

    const Stack: ProjectTechnology[] = [
    { id: 1, name: 'Go', icon_slug: 'go' },
    { id: 2, name: 'Rust', icon_slug: 'rust' },
    { id: 3, name: 'TypeScript', icon_slug: 'typescript' },
    { id: 4, name: 'JavaScript', icon_slug: 'javascript' },
    { id: 5, name: 'Python', icon_slug: 'python' },
    { id: 6, name: 'Java', icon_slug: 'java' },
    { id: 7, name: 'C++', icon_slug: 'cpp' },
    { id: 8, name: 'Kotlin', icon_slug: 'kotlin' },
    { id: 12, name: 'Svelte', icon_slug: 'svelte' },
    { id: 13, name: 'React', icon_slug: 'react' },
    { id: 14, name: 'Next.js', icon_slug: 'nextdotjs' },
    { id: 16, name: 'Express', icon_slug: 'express' },
    { id: 17, name: 'PostgreSQL', icon_slug: 'postgresql' },
    { id: 18, name: 'Neo4j', icon_slug: 'neo4j' },
    { id: 19, name: 'Redis', icon_slug: 'redis' },
    { id: 20, name: 'Docker', icon_slug: 'docker' },
    { id: 21, name: 'Google Cloud Platform', icon_slug: 'googlecloud' },
    { id: 22, name: 'Microsoft Azure', icon_slug: 'azure' },
    { id: 23, name: 'Nginx', icon_slug: 'nginx' },
    { id: 24, name: 'Git', icon_slug: 'git' },
    { id: 25, name: 'Arduino / ESP8266', icon_slug: 'arduino' }
];

</script>

<div class="about-me-card">
    <p class="eyebrow">{$_('pages.home.aboutMe.eyebrow')}</p>
    <h2>{$_('pages.home.aboutMe.title')}</h2>
    <p>{$_('pages.home.aboutMe.description')}</p>

    <div class="contact-container">
    <ContactTag 
        type="email" 
        label="Email" 
        value="oscarligo6@gmail.com" 
    />
    <ContactTag 
        type="github" 
        label="GitHub Profile" 
        value="github.com/oscarligo" 
    />
    <ContactTag 
        type="linkedin" 
        label="LinkedIn" 
        value="linkedin.com/in/oscar-rompich-367569284" 
    />
    </div>
    <div class="tech-stack-title">{$_('pages.home.aboutMe.techStack')}</div>


    <div class="tech-stack">

        {#each Stack as tech}
            <TechTag label={tech.name} icon_slug={tech.icon_slug}/>
        {/each}
    </div>
</div>



