package services

import (
	"context"

	"github.com/deside01/effective_mobile/internal/database/db"
	"github.com/deside01/effective_mobile/internal/handlers/dto"
)

type SubscriptionService interface {
	Create(context.Context, dto.SubscriptionBody) (*db.Subscription, error)
}
