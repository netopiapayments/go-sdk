package netopia

type PaymentClient struct {
	cfg    Config
	logger Logger
}

func NewPaymentClient(cfg Config, logger Logger) (*PaymentClient, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	if logger == nil {
		logger = &DefaultLogger{}
	}

	return &PaymentClient{cfg: cfg, logger: logger}, nil
}

func (c *PaymentClient) GetLogger() Logger {
	return c.logger
}

func (c *PaymentClient) BaseURL() string {
	if c.cfg.IsLive {
		return "https://secure.netopia-payments.com/api/"
	}
	return "https://secure-sandbox.netopia-payments.com"
}
