-- name: GetRoles :many
SELECT * FROM "role" WHERE "deleted_at" IS NULL OFFSET $1 LIMIT $2;

-- name: CountRole :one
SELECT COUNT(*) FROM "role" WHERE "deleted_at" IS NULL;

-- name: GetRole :one
SELECT * FROM "role" WHERE id = $1;

-- name: NewRole :one
INSERT INTO "role" (name) VALUES ($1) RETURNING *;

-- name: UpdateRole :one
UPDATE "role" SET name = $1 WHERE id = $2 RETURNING *;

-- name: DeleteRole :exec
DELETE FROM "role" WHERE id = $1;
