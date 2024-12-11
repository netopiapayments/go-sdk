package main

import (
	"encoding/json"
	"fmt"
	"netopia" //NETOPIA SDK
	"netopia/requests"
	"os"
	"time"
)

func main() {

	// Set the API Key as an environment variable (optional), can also be set directly in the config.
	os.Setenv("NETOPIA_API_KEY", "G85osrR3aP2zFoCJPMmTtQSRwaqtbLmnf2SznyEcSkWGTi4SRWeYO36xsvc=")

	// Initialize Netopia Client Configuration
	cfg := netopia.Config{
		PosSignature:    "2PU4-RWFV-BR5X-0AM4-LQSL", // POS Signature
		IsLive:          false,                      // false = sandbox | true = production
		NotifyURL:       "http://yourdomain.com/ipn",
		RedirectURL:     "http://yourdomain.com/back", // Redirrect URL after payment
		PublicKey:       []byte("-----BEGIN PUBLIC KEY-----\n...publickey...\n-----END PUBLIC KEY-----"),
		ActiveKey:       "active_key",
		PosSignatureSet: []string{"2PU4-RWFV-BR5X-0AM4-LQSL"}, // A list of POS Signatures (allowed)
	}

	// Initialize Netopia Payment Client based on the configuration above
	client, err := netopia.NewPaymentClient(cfg)

	// Handle errors
	if err != nil {
		switch err {
		case netopia.ErrMissingAPIKey:
			fmt.Println("API Key is missing!")
		default:
			panic(err)
		}
	}

	// Prepare the StartPayment Request with necessary details
	startReq := &requests.StartPaymentRequest{
		Config: &requests.ConfigData{
			EmailTemplate: "",              // Email template for notifications
			EmailSubject:  "",              // Subject for the notification email
			NotifyURL:     cfg.NotifyURL,   // Notification URL
			RedirectURL:   cfg.RedirectURL, // Redirect URL
			Language:      "ro",            // Language for notifications
		},
		Payment: &requests.PaymentData{
			Options: requests.PaymentOptions{
				Installments: 0,
				Bonus:        0,
			},
			Instrument: requests.PaymentInstrument{
				Type:       "card",             // Payment type (e.g., card)
				Account:    "9900004810225098", // Card number
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
			OrderID:      "",                                              // Unique order ID
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
			Data: map[string]string{ // Custom data fields
				"property1": "string",
				"property2": "string",
			},
		},
	}

	// Validate the StartPaymentRequest before sending it to ensure all required fields are correctly populated
	err = startReq.Validate()
	if err != nil {
		fmt.Println("Request validation failed:", err)
		return
	}

	// Send the payment request and handle the response (start payment action)
	startResp, err := client.StartPayment(startReq)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// For example purpose or/and debugging, you can see JSON Request sent to the enpoint:
	jsonData, err := json.MarshalIndent(startReq, "", "  ")
	if err != nil {
		fmt.Println("Serialization error: ", err)
		return
	}
	///////////////////////////////////////////////

	fmt.Println(string(jsonData))
	fmt.Printf("Start payment response: %+v\n", startResp)
}
