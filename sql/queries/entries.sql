-- name: CreateEntry :one
INSERT INTO entries (category, friendly_name, host_name, ip_address, description)
VALUES (?, ?, ?, ?, ?)
RETURNING *;

-- name: GetEntries :many
SELECT * FROM entries
ORDER BY id ASC;

-- name: GetEntryByID :one
SELECT * FROM entries
WHERE ID = ?;

-- name: EditEntryByID :one
UPDATE entries
SET
    category = ?,
    friendly_name = ?,
    host_name = ?,
    ip_address = ?,
    description = ?
WHERE id = ?
RETURNING *;

-- name: GetEntriesByCategory :many
SELECT * FROM entries
WHERE category = ?;

-- name: DeleteEntry :exec
DELETE FROM entries
WHERE id = ?;
