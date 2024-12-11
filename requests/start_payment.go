package requests

import (
	"errors"
	"fmt"
)

type StartPaymentRequest struct {
	Config  *ConfigData  `json:"config"`
	Payment *PaymentData `json:"payment"`
	Order   *OrderData   `json:"order"`
}

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
