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
(15, 'Express', 'express'),

-- Databases
(16, 'PostgreSQL', 'postgresql'),
(17, 'Neo4j', 'neo4j'),
(18, 'Redis', 'redis'),

-- DevOps and Cloud 
(19, 'Docker', 'docker'),
(20, 'GCP', 'googlecloud'),
(21, 'Azure', 'azure'),
(22, 'Nginx', 'nginx'),
(23, 'Git', 'git'),

-- IoT and Hardware
(24, 'Arduino / ESP8266', 'arduino')

ON CONFLICT (id) DO NOTHING;


SELECT setval(
    pg_get_serial_sequence('technology', 'id'), 
    COALESCE(MAX(id), 1)
) FROM technology;