-- name: ListProjects :many
SELECT * FROM project
ORDER BY created_at DESC;

-- name: GetProjectWithTechnologies :many
SELECT 
    p.id as project_id, p.title, p.description_short, p.featured,
    t.id as tech_id, t.name as tech_name, t.icon_slug
FROM project p
LEFT JOIN project_technology pt ON p.id = pt.project_id
LEFT JOIN technology t ON pt.technology_id = t.id
WHERE p.id = $1;