
-- name: URLGetAll :many
SELECT * FROM "url" LIMIT ? OFFSET ?;

-- name: URLGetByOriginal :one
SELECT * FROM "url" WHERE "original" = ?;

-- name: URLGetByID :one
SELECT * FROM "url" WHERE "id" = ?;

-- name: URLGetShortenCode :one
SELECT * FROM "url" WHERE "shortened" = ?;

-- name: URLDeleteUrlByID :exec
DELETE FROM "url" WHERE "id" = ?;

-- name: URLDeleteUrlByOriginal :exec
DELETE FROM "url" WHERE "original" = ?;

-- name: URLDeleteUrlByShortened :exec
DELETE FROM "url" WHERE "shortened" = ?;

-- name: URLCreate :one
INSERT INTO "url" ("id","original", "shortened", "clicks", "created")
VALUES (?, ? , ? ,? , ?) RETURNING *;

-- name: URLUpdateClicks :exec
UPDATE url SET "clicks" = "clicks" + 1 WHERE "id" = ?;
