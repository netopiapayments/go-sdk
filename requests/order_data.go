package requests

import (
	"errors"
	"fmt"
	"time"
)

type OrderData struct {
	NtpID        string          `json:"ntpID"`
	PosSignature string          `json:"posSignature"`
	DateTime     string          `json:"dateTime"`
	Description  string          `json:"description"`
	OrderID      string          `json:"orderID"`
	Amount       float64         `json:"amount"`
	Currency     string          `json:"currency"`
	Billing      BillingShipping `json:"billing"`
	Shipping     BillingShipping `json:"shipping"`
	Products     []Product       `json:"products"`
	Installments struct {
		Selected  int   `json:"selected"`
		Available []int `json:"available"`
	} `json:"installments"`
	Data interface{} `json:"data"`
}

func (o *OrderData) Validate() error {
	if o.PosSignature == "" {
		return errors.New("posSignature is required")
	}
	if _, err := time.Parse(time.RFC3339, o.DateTime); err != nil {
		return errors.New("dateTime must be in RFC3339 format")
	}
	if o.Description == "" {
		return errors.New("description is required")
	}
	if o.OrderID == "" {
		return errors.New("orderID is required")
	}
	if o.Amount < 0 {
		return errors.New("amount cannot be negative")
	}
	if len(o.Currency) != 3 {
		return errors.New("currency must be 3-letter ISO code")
	}

	if err := o.Billing.Validate(); err != nil {
		return fmt.Errorf("billing validation failed: %w", err)
	}

	for i, p := range o.Products {
		if err := p.Validate(); err != nil {
			return fmt.Errorf("product %d validation failed: %w", i, err)
		}
	}

	return nil
}
