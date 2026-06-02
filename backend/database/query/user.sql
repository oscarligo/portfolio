-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: CreateUser :one
-- Útil para tu script CLI inicial de creación de administrador
INSERT INTO users (
    username, password_hash
) VALUES (
    $1, $2
) RETURNING *;