package router

import (
	"github.com/deside01/effective_mobile/internal/handlers"
	"github.com/go-chi/chi/v5"
)

func NewRouter(sh *handlers.SubscriptionHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Route("/api/v1", func(r chi.Router) {
		r.Get("/subscriptions/{id}", sh.GetSubscription)

		r.Post("/subscriptions", sh.CreateSubscription)
		r.Get("/subscriptions", sh.GetSubscriptionsPage)
		r.Put("/subscriptions/{id}", sh.UpdateSubscription)
	})

	return r
}
