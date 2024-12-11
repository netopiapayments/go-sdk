package responses

type ApiResponse struct {
	Code    *string `json:"code,omitempty"`
	Message *string `json:"message,omitempty"`
}

type Attributes map[string]string

type Payment struct {
	Method         string                   `json:"method,omitempty"`
	AllowedMethods []string                 `json:"allowedMethods,omitempty"`
	NtpID          string                   `json:"ntpID,omitempty"`
	Rrn            string                   `json:"rrn,omitempty"`
	Status         int                      `json:"status,omitempty"`
	Amount         float64                  `json:"amount,omitempty"`
	Currency       string                   `json:"currency,omitempty"`
	PaymentURL     string                   `json:"paymentURL,omitempty"`
	Token          string                   `json:"token,omitempty"`
	OperationDate  string                   `json:"operationDate,omitempty"`
	Options        *PaymentOptions          `json:"options,omitempty"`
	Binding        *PaymentBinding          `json:"binding,omitempty"`
	Instrument     *PaymentInstrumentNotify `json:"instrument,omitempty"`
	Data           Attributes               `json:"data,omitempty"`
}

type PaymentInstrumentNotify struct {
	PanMasked   string `json:"panMasked,omitempty"`
	PanCategory string `json:"panCategory,omitempty"`
	Issuer      string `json:"issuer,omitempty"`
	Country     int    `json:"country,omitempty"`
}

type PaymentOptions struct {
	Installments int                       `json:"installments,omitempty"`
	Bonus        int                       `json:"bonus,omitempty"`
	Split        []PaymentSplitDestination `json:"split,omitempty"`
}

type PaymentSplitDestination struct {
	PosID  int     `json:"posID,omitempty"`
	Amount float64 `json:"amount,omitempty"`
}

type PaymentBinding struct {
	Token       string `json:"token,omitempty"`
	ExpireMonth int    `json:"expireMonth,omitempty"`
	ExpireYear  int    `json:"expireYear,omitempty"`
}

type ErrorWithDetails struct {
	Code    string           `json:"code,omitempty"`
	Message string           `json:"message,omitempty"`
	Details []ErrorWithField `json:"details,omitempty"`
}

type ErrorWithField struct {
	Code       string     `json:"code,omitempty"`
	Message    string     `json:"message,omitempty"`
	Field      string     `json:"field,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
}
