package db_model

type SubscribeAddressRequest struct {
	Address string `json:"address"`
}

type SubscribeAddressResponse struct {
	SubscribeSuccess bool `json:"subscribe_success"`
}
