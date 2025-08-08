-- name: CreateFeedFollow :one
WITH inserted_feeds_follow AS (
    INSERT INTO feeds_follows (id, created_at, updated_at, user_id, feed_id)
    VALUES ($1, $2, $3, $4, $5)
    RETURNING *
)
SELECT
    inserted_feeds_follow.*,
    feeds.name AS feed_name,
    users.name AS user_name
FROM inserted_feeds_follow
INNER JOIN feeds ON inserted_feeds_follow.feed_id = feeds.id
INNER JOIN users ON inserted_feeds_follow.user_id = users.id;
--

-- name: GetFeedFollowsForUser :many
SELECT feeds_follows.*, feeds.name AS feed_name, users.name AS user_name
FROM feeds_follows
INNER JOIN feeds ON feeds_follows.feed_id = feeds.id
INNER JOIN users ON feeds_follows.user_id = users.id
WHERE feeds_follows.user_id = $1;
--

-- name: DeleteFeedFollow :exec
DELETE FROM feeds_follows WHERE feed_id = $1 AND user_id = $2;
--
