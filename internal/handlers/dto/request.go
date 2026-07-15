package dto

type SubscriptionBody struct {
	ServiceName string `json:"service_name,omitempty"`
	Price       int    `json:"price,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	SubDate     string `json:"sub_date,omitempty"`
	ExpDate     string `json:"exp_date,omitempty"`
}
