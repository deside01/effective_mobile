package repos

import (
	"context"

	"github.com/deside01/effective_mobile/internal/database/db"
)

type SubscriptionRepository interface {
	Create(context.Context, db.CreateSubscriptionParams) (*db.Subscription, error)
	GetByID(context.Context, int64) (*db.Subscription, error)
	GetAll(context.Context, db.GetSubscriptionsPageParams) ([]db.Subscription, error)
	UpdateByID(context.Context, db.UpdateSubscriptionByIDParams) error
}
