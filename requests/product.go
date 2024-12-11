package requests

import "errors"

type Product struct {
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Vat      float64 `json:"vat"`
}

func (p *Product) Validate() error {
	if p.Name == "" {
		return errors.New("product name is required")
	}
	if p.Code == "" {
		return errors.New("product code is required")
	}
	if p.Price <= 0 {
		return errors.New("price must be greater than 0 or 0")
	}
	return nil
}
