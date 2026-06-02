CREATE TABLE IF NOT EXISTS project (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description_short TEXT NOT NULL,
    description_long TEXT NOT NULL,
    repo_url TEXT NOT NULL,
    live_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(), 
    featured BOOLEAN NOT NULL DEFAULT FALSE 
); 

CREATE TABLE IF NOT EXISTS technology (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    icon_slug TEXT NOT NULL
); 

CREATE TABLE IF NOT EXISTS project_technology ( 
    project_id INTEGER NOT NULL,
    technology_id INTEGER NOT NULL,
    PRIMARY KEY (project_id, technology_id),
    FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE,
    FOREIGN KEY (technology_id) REFERENCES technology (id) ON DELETE CASCADE
); 

CREATE TABLE IF NOT EXISTS project_images (
    id SERIAL PRIMARY KEY,
    project_id INTEGER NOT NULL,
    image_url TEXT NOT NULL,
    FOREIGN KEY (project_id) REFERENCES project (id) ON DELETE CASCADE
); 
