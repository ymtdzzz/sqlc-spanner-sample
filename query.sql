-- name: GetUser :one
SELECT * FROM Users
WHERE ID = ? LIMIT 1;

-- name: GetUserWithAddresses :many
SELECT * FROM Users u
JOIN Addresses a ON u.ID = a.UserID
WHERE u.ID = ?;

-- name: GetCounter :one
SELECT * FROM Counters
WHERE ID = ? LIMIT 1;

-- name: UpdateCounter :exec
UPDATE Counters SET Count = ?
WHERE ID = ?;
