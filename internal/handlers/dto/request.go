package dto

type SubscriptionBody struct {
	ServiceName string `json:"service_name"`
	Price       int    `json:"price"`
	UserID      string `json:"user_id"`
	SubDate     string `json:"sub_date"`
	ExpDate     string `json:"exp_date"`
}
