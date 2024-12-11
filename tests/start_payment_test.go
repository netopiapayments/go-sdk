package tests

import (
	"testing"
	"time"

	"netopia/requests"
)

func TestStartPaymentRequest_Validate(t *testing.T) {
	validRequest := &requests.StartPaymentRequest{
		Config: &requests.ConfigData{
			NotifyURL:   "https://example.com/notify",
			RedirectURL: "https://example.com/redirect",
			Language:    "ro",
		},
		Payment: &requests.PaymentData{
			Options: requests.PaymentOptions{
				Installments: 0,
				Bonus:        0,
			},
			Instrument: requests.PaymentInstrument{
				Type:       "card",
				Account:    "4111111111111111",
				ExpMonth:   12,
				ExpYear:    time.Now().Year() + 1,
				SecretCode: "123",
			},
		},
		Order: &requests.OrderData{
			PosSignature: "XXXX-XXXX",
			DateTime:     time.Now().UTC().Format(time.RFC3339),
			Description:  "Test order",
			OrderID:      "order_123",
			Amount:       10.5,
			Currency:     "RON",
			Billing: requests.BillingShipping{
				Email:     "user@example.com",
				FirstName: "John",
				LastName:  "Doe",
				City:      "Bucuresti",
				Country:   642,
			},
			Products: []requests.Product{
				{
					Name:  "test product",
					Code:  "SKU123",
					Price: 10.5,
				},
			},
		},
	}

	if err := validRequest.Validate(); err != nil {
		t.Errorf("expected no error for valid request, got: %v", err)
	}

	// Test no config
	invalidRequest := &requests.StartPaymentRequest{
		Payment: validRequest.Payment,
		Order:   validRequest.Order,
	}
	if err := invalidRequest.Validate(); err == nil {
		t.Errorf("expected error when config is missing")
	}

	// Test no payment
	invalidRequest = &requests.StartPaymentRequest{
		Config: validRequest.Config,
		Order:  validRequest.Order,
	}
	if err := invalidRequest.Validate(); err == nil {
		t.Errorf("expected error when payment is missing")
	}

	// Test no order
	invalidRequest = &requests.StartPaymentRequest{
		Config:  validRequest.Config,
		Payment: validRequest.Payment,
	}
	if err := invalidRequest.Validate(); err == nil {
		t.Errorf("expected error when order is missing")
	}
}

func TestConfigData_Validate(t *testing.T) {
	cfg := &requests.ConfigData{
		NotifyURL:   "https://example.com/notify",
		RedirectURL: "https://example.com/redirect",
		Language:    "en",
	}

	if err := cfg.Validate(); err != nil {
		t.Errorf("expected no error for valid config, got: %v", err)
	}

	cfg.NotifyURL = ""
	if err := cfg.Validate(); err == nil {
		t.Errorf("expected error for missing notifyUrl")
	}

	cfg.NotifyURL = "https://example.com/notify"
	cfg.RedirectURL = ""
	if err := cfg.Validate(); err == nil {
		t.Errorf("expected error for missing redirectUrl")
	}

	cfg.RedirectURL = "https://example.com/redirect"
	cfg.Language = "english"
	if err := cfg.Validate(); err == nil {
		t.Errorf("expected error for invalid language length")
	}
}

func TestPaymentData_Validate(t *testing.T) {
	payment := &requests.PaymentData{
		Options: requests.PaymentOptions{
			Installments: 0,
			Bonus:        0,
		},
		Instrument: requests.PaymentInstrument{
			Type:       "card",
			Account:    "4111111111111111",
			ExpMonth:   10,
			ExpYear:    time.Now().Year() + 1,
			SecretCode: "123",
		},
	}

	if err := payment.Validate(); err != nil {
		t.Errorf("expected no error for valid payment, got: %v", err)
	}

	payment.Options.Installments = -1
	if err := payment.Validate(); err == nil {
		t.Errorf("expected error for negative installments")
	}
	payment.Options.Installments = 0
	payment.Options.Bonus = -10
	if err := payment.Validate(); err == nil {
		t.Errorf("expected error for negative bonus")
	}
}

func TestPaymentInstrument_Validate(t *testing.T) {
	inst := requests.PaymentInstrument{
		Type:       "card",
		Account:    "4111111111111111",
		ExpMonth:   12,
		ExpYear:    time.Now().Year() + 1,
		SecretCode: "123",
	}
	if err := inst.Validate(); err != nil {
		t.Errorf("expected no error for valid card instrument, got: %v", err)
	}

	inst.ExpMonth = 13
	if err := inst.Validate(); err == nil {
		t.Errorf("expected error for invalid expMonth")
	}
	inst.ExpMonth = 12

	inst.ExpYear = time.Now().Year() - 1
	if err := inst.Validate(); err == nil {
		t.Errorf("expected error for expired year")
	}
	inst.ExpYear = time.Now().Year() + 1

	inst.SecretCode = "abc"
	if err := inst.Validate(); err == nil {
		t.Errorf("expected error for non-numeric secretCode")
	}
}

func TestOrderData_Validate(t *testing.T) {
	order := &requests.OrderData{
		PosSignature: "POS-TEST",
		DateTime:     time.Now().UTC().Format(time.RFC3339),
		Description:  "Some description",
		OrderID:      "orderID",
		Amount:       100.5,
		Currency:     "RON",
		Billing: requests.BillingShipping{
			Email:     "user@example.com",
			FirstName: "John",
			LastName:  "Doe",
			City:      "Bucharest",
			Country:   642,
		},
		Products: []requests.Product{
			{
				Name:  "Test product",
				Code:  "SKU01",
				Price: 10,
			},
		},
	}
	if err := order.Validate(); err != nil {
		t.Errorf("expected no error for valid order, got: %v", err)
	}

	order.PosSignature = ""
	if err := order.Validate(); err == nil {
		t.Errorf("expected error for missing posSignature")
	}

	order.PosSignature = "POS-TEST"
	order.DateTime = "invalid date"
	if err := order.Validate(); err == nil {
		t.Errorf("expected error for invalid dateTime")
	}
}

func TestBillingShipping_Validate(t *testing.T) {
	b := &requests.BillingShipping{
		Email:     "user@example.com",
		FirstName: "Jane",
		LastName:  "Doe",
		City:      "Cluj",
		Country:   642,
	}
	if err := b.Validate(); err != nil {
		t.Errorf("expected no error for valid billing, got: %v", err)
	}

	b.Email = ""
	if err := b.Validate(); err == nil {
		t.Errorf("expected error for missing email")
	}

	b.Email = "user@example.com"
	b.FirstName = ""
	if err := b.Validate(); err == nil {
		t.Errorf("expected error for missing firstName")
	}

	b.FirstName = "Jane"
	b.LastName = ""
	if err := b.Validate(); err == nil {
		t.Errorf("expected error for missing lastName")
	}

	b.LastName = "Doe"
	b.City = ""
	if err := b.Validate(); err == nil {
		t.Errorf("expected error for missing city")
	}

	b.City = "Cluj"
	b.Country = 0
	if err := b.Validate(); err == nil {
		t.Errorf("expected error for non-positive country")
	}
}

func TestProduct_Validate(t *testing.T) {
	p := &requests.Product{
		Name:  "Test product",
		Code:  "SKU123",
		Price: 9.99,
	}
	if err := p.Validate(); err != nil {
		t.Errorf("expected no error for valid product, got: %v", err)
	}

	p.Name = ""
	if err := p.Validate(); err == nil {
		t.Errorf("expected error for missing product name")
	}

	p.Name = "Test product"
	p.Code = ""
	if err := p.Validate(); err == nil {
		t.Errorf("expected error for missing product code")
	}

	p.Code = "SKU123"
	p.Price = 0
	if err := p.Validate(); err == nil {
		t.Errorf("expected error for 0 price (should be > 0)")
	}
}
