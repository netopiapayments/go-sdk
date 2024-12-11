package requests

type PaymentData struct {
	Options    PaymentOptions    `json:"options"`
	Instrument PaymentInstrument `json:"instrument"`
	Data       map[string]string `json:"data"`
}

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

type ConfigData struct {
	EmailTemplate string `json:"emailTemplate"`
	EmailSubject  string `json:"emailSubject"`
	NotifyURL     string `json:"notifyUrl"`
	RedirectURL   string `json:"redirectUrl"`
	Language      string `json:"language"`
}

type Installments struct {
	Selected  int   `json:"selected"`
	Available []int `json:"available"`
}

type OrderData struct {
	NtpID        string          `json:"ntpID"`
	PosSignature string          `json:"posSignature"`
	DateTime     string          `json:"dateTime"`
	Description  string          `json:"description"`
	OrderID      string          `json:"orderID"`
	Amount       float64         `json:"amount"`
	Currency     string          `json:"currency"`
	Billing      BillingShipping `json:"billing"`
	Shipping     BillingShipping `json:"shipping"`
	Products     []Product       `json:"products"`
	Installments struct {
		Selected  int   `json:"selected"`
		Available []int `json:"available"`
	} `json:"installments"`
	Data interface{} `json:"data"`
}

type PaymentInstrument struct {
	Type       string `json:"type"`
	Account    string `json:"account"`
	ExpMonth   int    `json:"expMonth"`
	ExpYear    int    `json:"expYear"`
	SecretCode string `json:"secretCode"`
	Token      string `json:"token"`
}

type PaymentOptions struct {
	Installments int `json:"installments"`
	Bonus        int `json:"bonus"`
}

type Product struct {
	Name     string  `json:"name"`
	Code     string  `json:"code"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Vat      float64 `json:"vat"`
}

type StartPaymentRequest struct {
	Config  *ConfigData  `json:"config"`
	Payment *PaymentData `json:"payment"`
	Order   *OrderData   `json:"order"`
}

type StatusRequest struct {
	PosID   string `json:"posID"`
	NtpID   string `json:"ntpID"`
	OrderID string `json:"orderID"`
}

type VerifyAuthRequest struct {
	AuthenticationToken string            `json:"authenticationToken"`
	NtpID               string            `json:"ntpID"`
	FormData            map[string]string `json:"formData"`
}
