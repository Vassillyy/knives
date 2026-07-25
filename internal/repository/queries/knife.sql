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
SET name=$1, description=$2, price=$3, material=$4, blade_length=$5, handle=$6, brand=$7, updated_at=NOW()
WHERE id=$8 AND deleted_at IS NULL;

-- name: Delete
UPDATE knives
SET deleted_at = NOW()
WHERE id=$1 AND deleted_at IS NULL;
