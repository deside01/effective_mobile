package router

import (
	"github.com/deside01/effective_mobile/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func NewRouter(sh *handlers.SubscriptionHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/subscriptions", sh.CreateSubscription)

	return r
}
