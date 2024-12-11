package netopia

import (
	"os"
)

type PaymentClient struct {
	cfg Config
}

func NewPaymentClient(cfg Config) (*PaymentClient, error) {
	if cfg.ApiKey == "" {
		cfg.ApiKey = os.Getenv("NETOPIA_API_KEY")
	}

	if cfg.ApiKey == "" {
		return nil, ErrMissingAPIKey
	}

	if cfg.PosSignature == "" {
		return nil, ErrMissingPosSignature
	}

	return &PaymentClient{cfg: cfg}, nil
}

func (c *PaymentClient) BaseURL() string {
	if c.cfg.IsLive {
		return "https://secure.netopia-payments.com/api/"
	}
	return "https://secure-sandbox.netopia-payments.com"
}
