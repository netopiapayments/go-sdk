package main

import (
	"github.com/netopiapayments/go-sdk/requests"

	netopia "github.com/netopiapayments/go-sdk"
)

func main() {

	cfg := netopia.Config{
		ApiKey:          "gfcqLrjzcJGxxwiSDToGP9kOo3j1SemD_2cu13gwpmblql1QCeSfAhitc0o=",
		PosSignature:    "2TEU-SJYL-Q9J8-2DMG-XRW3",
		IsLive:          false,
		NotifyURL:       "https://google.ro/",
		RedirectURL:     "https://google.ro/",
		PublicKey:       []byte("-----BEGIN CERTIFICATE-----MIIC3zCCAkigAwIBAgIBATANBgkqhkiG9w0BAQsFADCBiDELMAkGA1UEBhMCUk8xEjAQBgNVBAgTCUJ1Y2hhcmVzdDESMBAGA1UEBxMJQnVjaGFyZXN0MRAwDgYDVQQKEwdORVRPUElBMSEwHwYDVQQLExhORVRPUElBIERldmVsb3BtZW50IHRlYW0xHDAaBgNVBAMTE25ldG9waWEtcGF5bWVudHMucm8wHhcNMjUwMjI4MDY0MTAyWhcNMzUwMjI2MDY0MTAyWjCBiDELMAkGA1UEBhMCUk8xEjAQBgNVBAgTCUJ1Y2hhcmVzdDESMBAGA1UEBxMJQnVjaGFyZXN0MRAwDgYDVQQKEwdORVRPUElBMSEwHwYDVQQLExhORVRPUElBIERldmVsb3BtZW50IHRlYW0xHDAaBgNVBAMTE25ldG9waWEtcGF5bWVudHMucm8wgZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBALwh0/NhEpZFuKvghZ9N75CXba05MWNCh422kcfFKbqP5YViCUBg3Mc5ZYd1e0Xi9Ui1QI2Z/jvvchrDZGQwjarApr3S9bowHEkZH81ZolOoPHBZbYpA28BIyHYRcaTXjLtiBGvjpwuzljmXeBoVLinIaE0IUpMen9MLWG2fGMddAgMBAAGjVzBVMA4GA1UdDwEB/wQEAwIFoDATBgNVHSUEDDAKBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MB0GA1UdDgQWBBQ9yXChMGxzUzQflmkXT1oyIBoetTANBgkqhkiG9w0BAQsFAAOBgQCYzIorfGX/sehOpLJp0HDC3JUmqTeyZOO7k7FTaxegHzpN2cv54/RRTUiHo1Geufb9/EhveilD3dRgWQDPtkHLy1sF0+wD5trHsNP4NUdaLr8gv19VKd4u3sbgBfBcQ4gdcyCW+p0RbQkL5TbzWl/lF64yLhsZHysKOe1CazRlrg==-----END CERTIFICATE-----"),
		PosSignatureSet: []string{""},
	}

	logger := &netopia.DefaultLogger{}

	client, err := netopia.NewPaymentClient(cfg, logger)
	if err != nil {
		logger.Errorf("Failed to initialize Netopia Payment Client:", err)
		return
	}

	startReq := &requests.StartPaymentRequest{
		Config: &requests.ConfigData{
			Language:    "ro",
			NotifyURL:   cfg.NotifyURL,
			RedirectURL: cfg.RedirectURL,
		},
		Payment: &requests.PaymentData{
			Options: requests.PaymentOptions{
				Installments: 0,
				Bonus:        0,
			},
			Instrument: requests.PaymentInstrument{
				Type:       "card",
				Account:    "9900004810225098",
				ExpMonth:   12,
				ExpYear:    2025,
				SecretCode: "111",
				Token:      "",
			},
			Data: map[string]string{
				"BROWSER_USER_AGENT": "Mozilla/5.0",
				"BROWSER_TZ":         "Europe/Bucharest",
			},
		},
		Order: &requests.OrderData{
			PosSignature: cfg.PosSignature,
			DateTime:     "2024-12-12T14:30:00Z",
			Description:  "Test Payment",
			OrderID:      "ORDER-12347683",
			Amount:       100.50,
			Currency:     "RON",
			Billing: requests.BillingShipping{
				Email:       "client@example.com",
				Phone:       "0741234567",
				FirstName:   "John",
				LastName:    "Doe",
				City:        "Bucuresti",
				Country:     642,
				CountryName: "Romania",
				State:       "Bucuresti",
				PostalCode:  "010101",
				Details:     "Test",
			},
			Shipping: requests.BillingShipping{
				Email:       "client@example.com",
				Phone:       "0741234567",
				FirstName:   "John",
				LastName:    "Doe",
				City:        "Bucuresti",
				Country:     642,
				CountryName: "Romania",
				State:       "Bucuresti",
				PostalCode:  "010101",
				Details:     "Test",
			},
			Products: []requests.Product{
				{
					Name:     "Test Product",
					Code:     "PROD001",
					Category: "Test Category",
					Price:    100.50,
					Vat:      19,
				},
			},
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
