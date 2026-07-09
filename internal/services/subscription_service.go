package services

import (
	"context"
	"fmt"
	"time"

	"github.com/deside01/effective_mobile/internal/database/db"
	"github.com/deside01/effective_mobile/internal/handlers/dto"
	"github.com/deside01/effective_mobile/internal/repos"
	"github.com/jackc/pgx/v5/pgtype"
)

type SubscriptionSvc struct {
	repo *repos.SubscriptionRepo
}

func NewSubscriptionService(repo *repos.SubscriptionRepo) *SubscriptionSvc {
	return &SubscriptionSvc{
		repo: repo,
	}
}

func (ss *SubscriptionSvc) Create(ctx context.Context, body dto.SubscriptionBody) (*db.Subscription, error) {
	var userID pgtype.UUID
	if err := userID.Scan(body.UserID); err != nil {
		return nil, fmt.Errorf("invalid user_id")
	}

	subDate, err := time.Parse("01-2006", body.SubDate)
	if err != nil {
		return nil, fmt.Errorf("invalid sub_date")
	}

	expDate, err := time.Parse("01-2006", body.ExpDate)
	if err != nil {
		return nil, fmt.Errorf("invalid exp_date")
	}

	return ss.repo.Create(ctx, db.CreateSubscriptionParams{
		ServiceName: body.ServiceName,
		Price:       int32(body.Price),
		UserID:      userID,
		SubDate: pgtype.Date{
			Time:  subDate,
			Valid: true,
		},
		ExpDate: pgtype.Date{
			Time:  expDate,
			Valid: true,
		},
	})
}
