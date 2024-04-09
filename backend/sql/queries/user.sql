-- name: CreateUser :one
-- CreateUser creates a new user in the database.
INSERT INTO users (id, username, email, password_hash, salt, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING * ;
 
-- name: GetUserByID :one
-- GetUserByID retrieves a user from the database by ID.
SELECT id, username, email, password_hash, salt, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetUserByUsername :one 
-- GetUserByUsername retrieves a user from the database by username.
SELECT * FROM users WHERE username = $1;

-- name: UpdateUser :exec
-- UpdateUser updates an existing user in the database.
UPDATE users
SET username = $2, email = $3, password_hash = $4, salt = $5, updated_at = $6
WHERE id = $1;

-- name: DeleteUser :exec
-- DeleteUser deletes a user from the database by ID.
DELETE FROM users
WHERE id = $1;