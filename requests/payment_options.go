package requests

type PaymentOptions struct {
	Installments int `json:"installments"`
	Bonus        int `json:"bonus"`
}
