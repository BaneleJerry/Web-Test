-- name: InsertAPIKey :one
INSERT INTO api_keys (id, user_id, api_key_value, salt, expiration_time, created_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;


-- name: SelectAPIKeysByUserID :one
SELECT * FROM api_keys WHERE user_id = $1;

-- name: UpdateAPIKeyExpirationTime :exec
UPDATE api_keys SET expiration_time = $2 WHERE id = $1;

-- name: DeleteAPIKeyByID :exec
DELETE FROM api_keys WHERE id = $1;

