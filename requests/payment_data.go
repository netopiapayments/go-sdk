package requests

import (
	"errors"
	"fmt"
)

type PaymentData struct {
	Options    PaymentOptions    `json:"options"`
	Instrument PaymentInstrument `json:"instrument"`
	Data       map[string]string `json:"data"`
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
