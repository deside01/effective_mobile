-- name: CreateSubscription :one
INSERT INTO subscriptions
    (service_name, price, user_id, sub_date, exp_date)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetSubscriptionsPage :many
SELECT
    id,
    service_name,
    price,
    user_id,
    sub_date,
    exp_date,
    created_at,
    updated_at
FROM subscriptions
ORDER BY id DESC
LIMIT $1
OFFSET $2;

-- name: GetSubscriptionByID :one
SELECT
    id,
    service_name,
    price,
    user_id,
    sub_date,
    exp_date,
    created_at,
    updated_at
FROM subscriptions
WHERE id = $1;

-- name: UpdateSubscriptionByID :exec
UPDATE subscriptions
SET
    service_name = COALESCE(sqlc.narg('service_name'), service_name),
    price = COALESCE(sqlc.narg('price'), price),
    sub_date = COALESCE(sqlc.narg('sub_date'), sub_date),
    exp_date = COALESCE(sqlc.narg('exp_date'), exp_date),
    updated_at = COALESCE(sqlc.narg('updated_at'), updated_at)
WHERE id = sqlc.arg('id');

-- name: DeleteSubscriptionByID :one
DELETE FROM subscriptions
WHERE id = $1
RETURNING id;

-- name: GetUserSummary :one
SELECT COALESCE(SUM(price), 0)::integer AS total_cost
FROM subscriptions
WHERE user_id = $1
  AND (sqlc.narg('service_name')::text IS NULL OR service_name = sqlc.narg('service_name'))
  AND (sqlc.arg('sub_date')::date IS NULL OR sub_date >= sqlc.arg('sub_date'))
  AND (sqlc.arg('exp_date')::date IS NULL OR exp_date <= sqlc.arg('exp_date'));
