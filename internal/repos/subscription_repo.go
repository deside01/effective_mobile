package repos

import (
	"context"

	"github.com/deside01/effective_mobile/internal/database/db"
)

type SubscriptionRepo struct {
	repo db.Querier
}

func NewSubscriptionRepo(repo db.Querier) *SubscriptionRepo {
	return &SubscriptionRepo{
		repo: repo,
	}
}

func (sr *SubscriptionRepo) Create(ctx context.Context, params db.CreateSubscriptionParams) (*db.Subscription, error) {
	sub, err := sr.repo.CreateSubscription(ctx, params)

	return &sub, err
}

func (sr *SubscriptionRepo) GetByID(ctx context.Context, id int64) (*db.Subscription, error) {
	sub, err := sr.repo.GetSubscriptionByID(ctx, id)
	return &sub, err
}

func (sr *SubscriptionRepo) GetAll(ctx context.Context, params db.GetSubscriptionsPageParams) ([]db.Subscription, error) {
	return sr.repo.GetSubscriptionsPage(ctx, params)
}

func (sr *SubscriptionRepo) UpdateByID(ctx context.Context, params db.UpdateSubscriptionByIDParams) error {
	return sr.repo.UpdateSubscriptionByID(ctx, params)
}
