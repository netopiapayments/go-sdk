package responses

type ApiResponse struct {
	Code    *string           `json:"code,omitempty"`
	Message *string           `json:"message,omitempty"`
	Error   *ErrorWithDetails `json:"error,omitempty"`
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

type StatusResponse struct {
	ApiResponse
	Bnpl           *BnplOptions `json:"bnpl,omitempty"`
	Merchant       *Merchant    `json:"merchant,omitempty"`
	Config         *Config      `json:"config,omitempty"`
	Order          *Order       `json:"order,omitempty"`
	Payment        *Payment     `json:"payment,omitempty"`
	CustomerAction *Action      `json:"customerAction,omitempty"`
}

type BnplOptions struct {
	Oney *Oney `json:"oney,omitempty"`
}

type Oney struct {
	MerchantUID    string       `json:"merchantUID,omitempty"`
	PaymentOptions []OneyOption `json:"paymentOptions,omitempty"`
}

type OneyOption struct {
	BusinessTransactionCode string `json:"businessTransactionCode,omitempty"`
	Title                   string `json:"title,omitempty"`
	Instalments             int    `json:"instalments,omitempty"`
}

type Merchant struct {
	MerchantName string `json:"merchantName,omitempty"`
	PosUrl       string `json:"posUrl,omitempty"`
	PosName      string `json:"posName,omitempty"`
	PosID        int    `json:"posID,omitempty"`
	QrType       int    `json:"qrType,omitempty"`
	ShowCancel   bool   `json:"showCancel,omitempty"`
	PosType      int    `json:"posType,omitempty"`
}

type Config struct {
	EmailTemplate string `json:"emailTemplate,omitempty"`
	EmailSubject  string `json:"emailSubject,omitempty"`
	CancelURL     string `json:"cancelUrl,omitempty"`
	NotifyURL     string `json:"notifyUrl,omitempty"`
	RedirectURL   string `json:"redirectUrl,omitempty"`
	Language      string `json:"language,omitempty"`
}

type Order struct {
	NtpID           string           `json:"ntpID,omitempty"`
	PosSignature    string           `json:"posSignature,omitempty"`
	DateTime        string           `json:"dateTime,omitempty"`
	Description     string           `json:"description,omitempty"`
	OrderID         string           `json:"orderID,omitempty"`
	Amount          float64          `json:"amount,omitempty"`
	Currency        string           `json:"currency,omitempty"`
	Billing         *Address         `json:"billing,omitempty"`
	Shipping        *ShippingAddress `json:"shipping,omitempty"`
	Products        []Product        `json:"products,omitempty"`
	Installments    *Installments    `json:"installments,omitempty"`
	Data            Attributes       `json:"data,omitempty"`
	ClientID        string           `json:"clientID,omitempty"`
	ScaExemptionInd string           `json:"scaExemptionInd,omitempty"`
}

type Installments struct {
	Selected  int   `json:"selected,omitempty"`
	Available []int `json:"available,omitempty"`
}

type Address struct {
	Email       string `json:"email"`
	Phone       string `json:"phone"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	City        string `json:"city"`
	Country     int    `json:"country"`
	CountryName string `json:"countryName"`
	State       string `json:"state"`
	PostalCode  string `json:"postalCode"`
	Details     string `json:"details"`
}

type ShippingAddress struct {
	Email      string `json:"email,omitempty"`
	Phone      string `json:"phone,omitempty"`
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	City       string `json:"city,omitempty"`
	Country    int    `json:"country,omitempty"`
	State      string `json:"state,omitempty"`
	PostalCode string `json:"postalCode,omitempty"`
	Details    string `json:"details,omitempty"`
}

type Product struct {
	Name     string  `json:"name,omitempty"`
	Code     string  `json:"code,omitempty"`
	Category string  `json:"category,omitempty"`
	Price    float64 `json:"price,omitempty"`
	Vat      float64 `json:"vat,omitempty"`
}

type StartPaymentResponse struct {
	ApiResponse
	Payment        *Payment `json:"payment,omitempty"`
	CustomerAction *Action  `json:"customerAction,omitempty"`
}

type Action struct {
	Type                string     `json:"type,omitempty"`
	URL                 string     `json:"url,omitempty"`
	AuthenticationToken string     `json:"authenticationToken,omitempty"`
	FormData            Attributes `json:"formData,omitempty"`
}

type VerifyAuthResponse struct {
	ApiResponse
	Payment *Payment `json:"payment,omitempty"`
}
