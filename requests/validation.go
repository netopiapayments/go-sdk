package requests

import (
	"errors"
	"fmt"
	"net/url"
	"time"
)

func (req *StartPaymentRequest) Validate() error {
	if req.Config == nil {
		return errors.New("config is required")
	}
	if err := req.Config.Validate(); err != nil {
		return fmt.Errorf("config validation failed: %w", err)
	}

	if req.Payment == nil {
		return errors.New("payment is required")
	}
	if err := req.Payment.Validate(); err != nil {
		return fmt.Errorf("payment validation failed: %w", err)
	}

	if req.Order == nil {
		return errors.New("order is required")
	}
	if err := req.Order.Validate(); err != nil {
		return fmt.Errorf("order validation failed: %w", err)
	}

	return nil
}

func (c *ConfigData) Validate() error {
	if c.NotifyURL == "" || !isValidURL(c.NotifyURL) {
		return errors.New("notifyUrl is invalid")
	}
	if c.RedirectURL == "" || !isValidURL(c.RedirectURL) {
		return errors.New("redirectUrl is invalid")
	}
	if len(c.Language) != 2 {
		return errors.New("language must be ISO 639-1 code (2 chars)")
	}
	return nil
}

func (p *PaymentData) Validate() error {
	if err := p.Instrument.Validate(); err != nil {
		return fmt.Errorf("instrument validation failed: %w", err)
	}

	if p.Options.Installments < 0 {
		return errors.New("installments cannot be negative")
	}

	if p.Options.Bonus < 0 {
		return errors.New("bonus cannot be negative")
	}
	return nil
}

func (inst *PaymentInstrument) Validate() error {
	if inst.Token != "" {
		if len(inst.Token) < 5 {
			return errors.New("token is too short")
		}
		return nil
	}

	if inst.Type == "card" {
		if len(inst.Account) < 12 || len(inst.Account) > 19 {
			return errors.New("invalid card number length")
		}
		if inst.ExpMonth < 1 || inst.ExpMonth > 12 {
			return errors.New("expMonth must be between 1 and 12")
		}
		if inst.ExpYear < time.Now().Year() {
			return errors.New("expYear cannot be in the past")
		}
		if len(inst.SecretCode) < 3 || len(inst.SecretCode) > 4 {
			return errors.New("secretCode must be 3 or 4 digits")
		}
	}
	return nil
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
		return errors.New("currency must be a 3-letter ISO code")
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

func (b *BillingShipping) Validate() error {
	if b.Email == "" {
		return errors.New("email is required")
	}
	if b.FirstName == "" || b.LastName == "" {
		return errors.New("firstName and lastName are required")
	}
	if b.City == "" {
		return errors.New("city is required")
	}
	if b.Country <= 0 {
		return errors.New("country must be positive")
	}
	return nil
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("product name is required")
	}
	if p.Code == "" {
		return errors.New("product code is required")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than 0")
	}
	return nil
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
