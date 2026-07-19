package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	errs "github.com/deside01/effective_mobile/internal/errors"
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
		utils.Error(w, http.StatusBadRequest, fmt.Sprint("unable to decode body: ", err))
		return
	}

	sub, err := sh.svc.CreateSubscription(r.Context(), body)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(w, http.StatusCreated, sub)
}

func (sh *SubscriptionHandler) GetSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	sub, err := sh.svc.GetSubscriptionByID(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrNotFound):
			utils.Error(w, http.StatusNotFound, err.Error())
		default:
			utils.Error(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.JSON(w, http.StatusOK, sub)
}

func (sh *SubscriptionHandler) GetSubscriptionsPage(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	subs, err := sh.svc.GetSubscriptionsPage(r.Context(), query)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(w, http.StatusOK, subs)
}

func (sh *SubscriptionHandler) UpdateSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	var body dto.SubscriptionBody

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()

	decoder.DisallowUnknownFields()

	if err := decoder.Decode(&body); err != nil {
		utils.Error(w, http.StatusBadRequest, fmt.Sprint("unable to decode body: ", err))
		return
	}

	err = sh.svc.UpdateSubscriptionByID(r.Context(), id, body)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.JSON(w, http.StatusNoContent, "")
}

func (sh *SubscriptionHandler) DeleteSubscription(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Error(w, http.StatusBadRequest, "invalid id")
		return
	}

	err = sh.svc.DeleteSubscriptionByID(r.Context(), id)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrNotFound):
			utils.Error(w, http.StatusNotFound, err.Error())
		default:
			utils.Error(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.JSON(w, http.StatusNoContent, "")
}

func (sh *SubscriptionHandler) GetSummary(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	total_cost, err := sh.svc.GetUserSummary(r.Context(), query)
	if err != nil {
		switch {
		case errors.Is(err, errs.ErrInvalidInput):
			utils.Error(w, http.StatusBadRequest, err.Error())
		default:
			utils.Error(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	utils.JSON(w, http.StatusOK, map[string]int32{
		"total_cost": total_cost,
	})
}
