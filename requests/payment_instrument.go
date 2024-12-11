package requests

import (
	"errors"
	"regexp"
	"time"
)

type PaymentInstrument struct {
	Type       string `json:"type"`
	Account    string `json:"account"`
	ExpMonth   int    `json:"expMonth"`
	ExpYear    int    `json:"expYear"`
	SecretCode string `json:"secretCode"`
	Token      string `json:"token"`
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
		match, _ := regexp.MatchString(`^[0-9]{3,4}$`, inst.SecretCode)
		if !match {
			return errors.New("secretCode must be 3 or 4 digits")
		}
	}

	return nil
}
