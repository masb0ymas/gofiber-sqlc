-- name: GetRoles :many
SELECT * FROM "role" WHERE "deleted_at" IS NULL OFFSET $1 LIMIT $2;

-- name: CountRole :one
SELECT COUNT(*) FROM "role" WHERE "deleted_at" IS NULL;

-- name: GetRole :one
SELECT * FROM "role" WHERE id = $1;

-- name: NewRole :one
INSERT INTO "role" (name) VALUES ($1) RETURNING *;

-- name: UpdateRole :one
UPDATE "role" SET name = $1 WHERE id = $2 AND "deleted_at" = NULL RETURNING *;

-- name: RestoreRole :exec
UPDATE "role" SET "deleted_at" = NULL WHERE id = $1;

-- name: SoftDeleteRole :exec
UPDATE "role" SET "deleted_at" = now() WHERE id = $1;

-- name: ForceDeleteRole :exec
DELETE FROM "role" WHERE id = $1;

-- name: RestoreRoles :exec
UPDATE "role" SET "deleted_at" = NULL WHERE id IN ($1);

-- name: SoftDeleteRoles :exec
UPDATE "role" SET "deleted_at" = now() WHERE id IN ($1);

-- name: ForceDeleteRoles :exec
DELETE FROM "role" WHERE id IN ($1);