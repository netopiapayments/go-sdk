package main

import (
	"time"

	"github.com/netopiapayments/go-sdk/requests"

	netopia "github.com/netopiapayments/go-sdk"
)

func main() {

	// Initialize Netopia Client Configuration
	cfg := netopia.Config{
		ApiKey:          "APIKEY",
		PosSignature:    "POSID", // POS Signature
		IsLive:          false,   // false = sandbox | true = production
		NotifyURL:       "x",
		RedirectURL:     "x", // Redirrect URL after payment
		PublicKey:       []byte("-----BEGIN PUBLIC KEY-----\n...publickey...\n-----END PUBLIC KEY-----"),
		ActiveKey:       "ACTIVEKEY",
		PosSignatureSet: []string{""}, // A list of POS Signatures (allowed)
	}

	logger := &netopia.DefaultLogger{}

	client, err := netopia.NewPaymentClient(cfg, logger)
	if err != nil {
		logger.Errorf("Failed to initialize Netopia Payment Client:", err)
		return
	}

	// Prepare the StartPayment Request with necessary details
	startReq := &requests.StartPaymentRequest{
		Config: &requests.ConfigData{
			EmailTemplate: "",                   // Email template for notifications
			EmailSubject:  "",                   // Subject for the notification email
			NotifyURL:     "http://google.com/", // Notification URL
			RedirectURL:   "http://google.com/", // Redirect URL
			Language:      "ro",                 // Language for notifications
		},
		Payment: &requests.PaymentData{
			Options: requests.PaymentOptions{
				Installments: 0,
				Bonus:        0,
			},
			Instrument: requests.PaymentInstrument{
				Type:       "card",             // Payment type (e.g., card)
				Account:    "9900518572831942", // Card number
				ExpMonth:   11,                 // Card expiration month
				ExpYear:    2025,               // Card expiration year
				SecretCode: "111",              // Card CVC/CVV
				Token:      "",
			},
			Data: map[string]string{
				// Browser and device data (required for 3DSecure payments)
				"BROWSER_USER_AGENT":    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/90.0.4430.212 Safari/537.36",
				"BROWSER_TZ":            "Europe/Bucharest",
				"BROWSER_COLOR_DEPTH":   "32",
				"BROWSER_JAVA_ENABLED":  "true",
				"BROWSER_LANGUAGE":      "en-US,en;q=0.9",
				"BROWSER_TZ_OFFSET":     "0",
				"BROWSER_SCREEN_WIDTH":  "1200",
				"BROWSER_SCREEN_HEIGHT": "1400",
				"BROWSER_PLUGINS":       "Chrome PDF Plugin, Chrome PDF Viewer, Native Client",
				"MOBILE":                "false",
				"SCREEN_POINT":          "false",
				"OS":                    "macOS",
				"OS_VERSION":            "10.15.7 (32-bit)",
				"IP_ADDRESS":            "127.0.0.1",
			},
		},
		Order: &requests.OrderData{
			NtpID:        "",                                              // Leave empty for a new transaction
			PosSignature: cfg.PosSignature,                                // POS signature (also known as POS ID) for this order
			DateTime:     time.Now().UTC().Format("2006-01-02T15:04:05Z"), // Current date and time but can be any date in future
			Description:  "DEMO API FROM WEB - SDK",                       // Order description
			OrderID:      "8",                                             // Unique order ID
			Amount:       0,                                               // Payment amount
			Currency:     "RON",                                           // Currency
			Billing: requests.BillingShipping{
				Email:       "client@test.com",
				Phone:       "0000000",
				FirstName:   "ClientPrenume",
				LastName:    "ClientNume",
				City:        "Bucuresti",
				Country:     642,
				CountryName: "Romania",
				State:       "Bucuresti",
				PostalCode:  "246513",
				Details:     "Fara Detalii",
			},
			Shipping: requests.BillingShipping{
				Email:       "client@test.com",
				Phone:       "0000000",
				FirstName:   "ClientPrenume",
				LastName:    "ClientNume",
				City:        "Bucuresti",
				Country:     642,
				CountryName: "Romania",
				State:       "Bucuresti",
				PostalCode:  "246513",
				Details:     "Fara Detalii",
			},
			Products: []requests.Product{ // Product list for the order
				{Name: "string", Code: "SKU", Category: "category", Price: 1, Vat: 19},
			},
			// Not Used
			Installments: struct {
				Selected  int   `json:"selected"`
				Available []int `json:"available"`
			}{
				Selected:  0,
				Available: []int{0},
			},
			Data: map[string]string{
				"property1": "string",
				"property2": "string",
			},
		},
	}

	if err := startReq.Validate(); err != nil {
		logger.Errorf("StartPaymentRequest validation failed:", err)
		return
	}

	startResp, err := client.StartPayment(startReq)
	if err != nil {
		logger.Errorf("Failed to start payment:", err)
		return
	}

	if startResp == nil {
		logger.Errorf("Invalid response: StartPaymentResponse is nil.")
		return
	}

	if startResp.Payment == nil {
		logger.Debugf("API Message: %s\n", *startResp.Message)
		return
	}

	if startResp.Error.Code == "00" {
		logger.Infof("Payment initiated successfully!")
		logger.Infof("\nPlease visit the Payment URL to complete your transaction.")
	} else {
		logger.Infof("Payment has errors!")
		logger.Debugf("Error message: %s\n", startResp.Error.Message)
	}

	logger.Debugf("Payment URL: %s\n", startResp.Payment.PaymentURL)
	logger.Debugf("Payment Token: %s\n", startResp.Payment.Token)
	logger.Debugf("Payment Status: %d\n", startResp.Payment.Status)

	if startResp.Payment.Binding != nil {
		logger.Debugf("Card Binding ExpireYear: %d\n", startResp.Payment.Binding.ExpireYear)
	} else {
		logger.Debugf("No binding information available.")
	}
}
