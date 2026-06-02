-- name: GetProject :one
SELECT * FROM project
WHERE id = $1 LIMIT 1;

-- name: ListProjects :many
SELECT * FROM project
ORDER BY created_at DESC;

-- name: ListFeaturedProjects :many
SELECT * FROM project
WHERE featured = true
ORDER BY created_at DESC;

-- name: CreateProject :one
INSERT INTO project (
    title, description_short, description_long, repo_url, live_url, video_url, featured
) VALUES (
    $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: CreateProjectImage :one
INSERT INTO project_images (
    project_id, image_url
) VALUES (
    $1, $2
) RETURNING *;

-- name: AssociateProjectTechnology :exec
INSERT INTO project_technology (
    project_id, technology_id
) VALUES (
    $1, $2
);

-- name: ListProjectImages :many
SELECT * FROM project_images
WHERE project_id = $1;

-- name: ListProjectTechnologies :many
SELECT t.* FROM technology t
JOIN project_technology pt ON t.id = pt.technology_id
WHERE pt.project_id = $1;

-- name: DeleteProject :exec
DELETE FROM project
WHERE id = $1;