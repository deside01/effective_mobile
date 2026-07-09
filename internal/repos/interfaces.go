package repos

import (
	"context"

	"github.com/deside01/effective_mobile/internal/database/db"
)

type SubscriptionRepository interface {
	Create(context.Context, db.CreateSubscriptionParams) (*db.Subscription, error)
}
