-- name: GetUsers :many
SELECT u.id, u.created_at, u.updated_at, u.deleted_at, u.fullname, u.email, u.phone, u.address, u.is_active, u.is_blocked, u.role_id, r.name as role_name
  FROM "user" u
  LEFT JOIN "role" r ON u.role_id = r.id
  WHERE u.deleted_at IS NULL 
  OFFSET $1 LIMIT $2;

-- name: CountUser :one
SELECT COUNT(*) FROM "user" WHERE "deleted_at" IS NULL;

-- name: GetUser :one
SELECT * FROM "user" WHERE id = $1;

-- name: GetUserWithRelation :one
SELECT u.id, u.created_at, u.updated_at, u.deleted_at, u.fullname, u.email, u.phone, u.address, u.is_active, u.is_blocked, u.role_id, r.name as role_name
  FROM "user" u
  LEFT JOIN "role" r ON u.role_id = r.id
  WHERE u.id = $1;

-- name: NewUser :one
INSERT INTO "user" (fullname, email, password, phone, token_verify, address, is_active, is_blocked, role_id)
  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING *;

-- name: UpdateUser :one
UPDATE "user" SET fullname=$2, email=$3, phone=$4, address=$5, is_active=$6, is_blocked=$7, role_id=$8
  WHERE id=$1 AND "deleted_at"=NULL RETURNING id;

-- name: RestoreUser :exec
UPDATE "user" SET "deleted_at" = NULL WHERE id = $1;

-- name: SoftDeleteUser :exec
UPDATE "user" SET "deleted_at" = now() WHERE id = $1;

-- name: ForceDeleteUser :exec
DELETE FROM "user" WHERE id = $1;

-- name: RestoreUsers :exec
UPDATE "user" SET "deleted_at" = NULL WHERE id IN ($1);

-- name: SoftDeleteUsers :exec
UPDATE "user" SET "deleted_at" = now() WHERE id IN ($1);

-- name: ForceDeleteUsers :exec
DELETE FROM "user" WHERE id IN ($1);