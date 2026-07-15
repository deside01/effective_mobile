package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/deside01/effective_mobile/internal/handlers/dto"
	"github.com/deside01/effective_mobile/internal/services"
	"github.com/deside01/effective_mobile/internal/utils"
	"github.com/go-chi/chi/v5"
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
	defer r.Body.Close()

	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		utils.Error(w, 401, fmt.Sprint("unable decode body: ", err))
		return
	}

	sub, err := sh.svc.CreateSubscription(r.Context(), body)
	if err != nil {
		utils.Error(w, 500, err.Error())
		return
	}

	utils.JSON(w, 201, sub)
}

func (sh *SubscriptionHandler) GetSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, 401, "invalid id")
		return
	}

	sub, err := sh.svc.GetSubscriptionByID(r.Context(), int64(id))
	if err != nil {
		utils.Error(w, 403, err.Error())
		return
	}

	utils.JSON(w, 200, sub)
}

func (sh *SubscriptionHandler) GetSubscriptionsPage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	subs, err := sh.svc.GetSubscriptionsPage(r.Context(), query)
	if err != nil {
		utils.Error(w, 500, err.Error())
		return
	}

	utils.JSON(w, 200, subs)
}

func (sh *SubscriptionHandler) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, 401, "invalid id")
		return
	}

	var body dto.SubscriptionBody

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		utils.Error(w, 401, fmt.Sprint("unable decode body: ", err))
		return
	}

	err = sh.svc.UpdateSubscriptionByID(r.Context(), id, body)
	if err != nil {
		utils.Error(w, 500, err.Error())
		return
	}

	utils.JSON(w, 204, "")
}
