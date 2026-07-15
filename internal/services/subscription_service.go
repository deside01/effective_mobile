package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/deside01/effective_mobile/internal/database/db"
	"github.com/deside01/effective_mobile/internal/handlers/dto"
	"github.com/deside01/effective_mobile/internal/repos"
	"github.com/deside01/effective_mobile/internal/utils"
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

func (ss *SubscriptionSvc) CreateSubscription(ctx context.Context, body dto.SubscriptionBody) (*db.Subscription, error) {
	var userID pgtype.UUID
	if err := userID.Scan(body.UserID); err != nil {
		return nil, fmt.Errorf("invalid user_id")
	}

	if body.ServiceName == "" {
		return nil, fmt.Errorf("empty service_name")
	}

	if body.Price <= 0 {
		return nil, fmt.Errorf("price must be greater than 0")
	}

	subDate := utils.PGDateParse(body.SubDate)
	if !subDate.Valid {
		return nil, fmt.Errorf("invalid sub_date")

	}

	expDate := utils.PGDateParse(body.ExpDate)

	return ss.repo.Create(ctx, db.CreateSubscriptionParams{
		ServiceName: body.ServiceName,
		Price:       int32(body.Price),
		UserID:      userID,
		SubDate:     subDate,
		ExpDate:     expDate,
	})
}

func (ss *SubscriptionSvc) GetSubscriptionByID(ctx context.Context, id int64) (*db.Subscription, error) {
	sub, err := ss.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("not found")
		}
		return nil, err
	}

	return sub, nil
}

func (ss *SubscriptionSvc) GetSubscriptionsPage(ctx context.Context, query url.Values) ([]db.Subscription, error) {
	limitStr := query.Get("limit")
	offsetStr := query.Get("page")

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 100
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
	}
	subs, err := ss.repo.GetAll(ctx, db.GetSubscriptionsPageParams{
		Limit:  int32(limit),
		Offset: int32(offset) * 10,
	})
	if err != nil {
		return nil, fmt.Errorf("err")
	}

	return subs, nil
}

func (ss *SubscriptionSvc) UpdateSubscriptionByID(ctx context.Context, id int, body dto.SubscriptionBody) error {
	serviceName := utils.PGTextParse(body.ServiceName)
	price := utils.PGInt4Parse(body.Price)
	subDate := utils.PGDateParse(body.SubDate)
	expDate := utils.PGDateParse(body.ExpDate)

	params := db.UpdateSubscriptionByIDParams{
		ID:          int64(id),
		ServiceName: serviceName,
		Price:       price,
		SubDate:     subDate,
		ExpDate:     expDate,
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now().UTC(),
			Valid: true,
		},
	}

	return ss.repo.UpdateByID(ctx, params)
}
