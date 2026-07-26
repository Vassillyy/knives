-- internal/repository/queries/knife.sql
-- name: GetAll
SELECT id, name, description, price, material, blade_length, handle, brand, created_at, updated_at, deleted_at
FROM knives
WHERE deleted_at IS NULL;

-- name: GetByID
SELECT id, name, description, price, material, blade_length, handle, brand, created_at, updated_at, deleted_at
FROM knives
WHERE id = $1 AND deleted_at IS NULL;

-- name: Create
INSERT INTO knives (id, name, description, price, material, blade_length, handle, brand)
VALUES (gen_random_uuid(), $1, $2, $3, $4, $5, $6, $7)
RETURNING id, created_at, updated_at;

-- name: Update
UPDATE knives
SET
  name = COALESCE($1, name),
  description = COALESCE($2, description),
  price = COALESCE($3, price),
  material = COALESCE($4, material),
  blade_length = COALESCE($5, blade_length),
  handle = COALESCE($6, handle),
  brand = COALESCE($7, brand),
  updated_at = NOW()
WHERE id = $8 AND deleted_at IS NULL;

-- name: Delete
UPDATE knives
SET deleted_at = NOW()
WHERE id=$1 AND deleted_at IS NULL;
