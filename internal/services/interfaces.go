package services

import (
	"context"
	"net/url"

	"github.com/deside01/effective_mobile/internal/database/db"
	"github.com/deside01/effective_mobile/internal/handlers/dto"
)

type SubscriptionService interface {
	CreateSubscription(context.Context, dto.SubscriptionBody) (*db.Subscription, error)
	GetSubscriptionByID(context.Context, int64) (*db.Subscription, error)
	GetSubscriptionsPage(context.Context, url.Values) ([]db.Subscription, error)
	UpdateSubscriptionByID(context.Context, int, dto.SubscriptionBody) error
}
