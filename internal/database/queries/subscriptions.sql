-- name: CreateSubscription :one
INSERT INTO subscriptions
    (service_name, price, user_id, sub_date, exp_date)
VALUES
    ($1, $2, $3, $4, $5)
RETURNING *;
