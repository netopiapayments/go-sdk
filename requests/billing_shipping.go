package requests

import "errors"

type BillingShipping struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	City        string `json:"city"`
	Country     int    `json:"country"`
	State       string `json:"state"`
	PostalCode  string `json:"postalCode"`
	CountryName string `json:"contryName"`
	Details     string `json:"details"`
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
