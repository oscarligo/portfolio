INSERT INTO technology (id, name, icon_slug) VALUES

-- Programming Languages
(1, 'Go', 'go'),
(2, 'Rust', 'rust'),
(3, 'TypeScript', 'typescript'),
(4, 'JavaScript', 'javascript'),
(5, 'Python', 'python'),
(7, 'Java', 'java'),
(8, 'C++', 'cpp'),
(9, 'Kotlin', 'kotlin'),
(10, 'PHP', 'php'),

-- Frameworks and Libraries
(11, 'Svelte', 'svelte'),
(12, 'React', 'react'),
(13, 'Next.js', 'nextdotjs'),
(14, 'Tailwind', 'tailwindcss'),

-- Databases
(15, 'PostgreSQL', 'postgresql'),
(16, 'Neo4j', 'neo4j'),
(17, 'Redis', 'redis'),

-- DevOps and Cloud 
(18, 'Docker', 'docker'),
(19, 'GCP', 'googlecloud'),
(20, 'Azure', 'azure'),
(21, 'Nginx', 'nginx'),
(22, 'Git', 'git'),

-- IoT and Hardware
(23, 'Arduino / ESP8266', 'arduino')

ON CONFLICT (id) DO NOTHING;


SELECT setval(
    pg_get_serial_sequence('technology', 'id'), 
    COALESCE(MAX(id), 1)
) FROM technology;