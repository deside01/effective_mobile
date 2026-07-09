package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deside01/effective_mobile/internal/handlers/dto"
	"github.com/deside01/effective_mobile/internal/services"
	"github.com/deside01/effective_mobile/internal/utils"
)

type SubscriptionHandler struct {
	svc services.SubscriptionService
}

func NewSubscriptionHandler(svc services.SubscriptionService) *SubscriptionHandler {
	return &SubscriptionHandler{
		svc: svc,
	}
}

func (sh *SubscriptionHandler) CreateSubscription(w http.ResponseWriter, r *http.Request) {
	var body dto.SubscriptionBody

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		utils.Error(w, 401, fmt.Sprint("unable decode body: ", err))
	}
}
