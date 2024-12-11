
  

# NETOPIA Payments Go SDK

The **NETOPIA Payments Go SDK** provides seamless integration with the **NETOPIA Payments API v2**. It allows developers to handle payments, IPN verification, and status retrieval efficiently within their Go applications.  

---

## **Table of Contents**

- [Features](#features)

- [Installation](#installation)

- [Getting Started](#getting-started)

- [Initialization](#initialization)

- [Configuration](#configuration)

- [API Reference](#api-reference)

- [PaymentClient Methods](#paymentclient-methods)

- [StartPayment](#startpayment)

- [GetStatus](#getstatus)

- [VerifyAuth](#verifyauth)

- [IPN Verification](#ipn-verification)

- [Error Handling](#error-handling)

- [Examples](#examples)

- [Useful Links](#useful-links)

---

## **Dependencies**

- github.com/golang-jwt/jwt/v5" 

## **Features**

- Start a payment with customer details, products, and billing/shipping data.

- Retrieve the status of an order.

- Verify 3DSecure authentications.

- Validate IPNs (Instant Payment Notifications) for order updates.

- Compatible with both Sandbox and Live environments.  

---

## **Installation**

To install the SDK, use `go get` to include it in your project:

```sh

go  get  github.com/x/x

```
---

## **Getting Started**

### **Initialization**

To initialize the NETOPIA Payments client, configure it with your API Key and POS Signature. 

The API Key can be initialized either as an environment variable or through configuration object:

```go
// By setting the environment variable:
os.Setenv("NETOPIA_API_KEY", "YOUR_API_KEY")

// By hydration of the request's config
netopia.Config{
	ApiKey: "YOUR_API_KEY",
}

// WARNING! Setting the request's config will override the environment variable
```

```go

import  "github.com/x/x"

cfg := netopia.Config {
	PosSignature: "XXXX-XXXX-XXXX-XXXX",
	ApiKey: "your-api-key",
	IsLive: false, // True = production | False = sandbox
	NotifyURL: "https://yourdomain.com/ipn",
	RedirectURL: "https://yourdomain.com/redirect",
	PublicKey: []byte("-----PUBLIC KEY-----"),
	ActiveKey: "your-active-key",
	PosSignatureSet: []string{"XXXX-XXXX-XXXX-XXXX"},
	HashMethod: "sha512", // Default hash method
}

client, err := netopia.NewPaymentClient(cfg)

if  err != nil {
	panic(err)
}

```

### **Configuration**

The `Config` struct contains the following fields:
  

| Field | Type | Required | Description |

| `PosSignature` | `string` | Yes | POS Signature for your NETOPIA account |

| `ApiKey` | `string` | Yes | API key generated in NETOPIA's admin panel |

| `IsLive` | `bool` | No | Whether to use the live environment (`true`) or sandbox (`false`) |

| `NotifyURL` | `string` | Yes | The URL where IPNs (order updates) will be sent |

| `RedirectURL` | `string` | Yes | The URL to redirect the customer after payment |

| `PublicKey` | `[]byte` | Yes | RSA public key provided by NETOPIA for signature verification |

| `ActiveKey` | `string` | Yes | Active key for your merchant account |

| `PosSignatureSet`| `[]string`| Yes | List of allowed POS Signatures |

| `HashMethod` | `string` | No | Hash method for IPN validation (default: `sha512`) |

---

## **API Reference**

### **PaymentClient Methods** 

#### **StartPayment**

Initiates a payment transaction.

```go

startReq := &requests.StartPaymentRequest{
		Config: &requests.ConfigData{
		NotifyURL: cfg.NotifyURL,
		RedirectURL: cfg.RedirectURL,
		Language: "ro", // ro is the default language if not provided
	},
	Payment: &requests.PaymentData{
		Options: requests.PaymentOptions{
			Installments: 1,
			Bonus: 0,
		},
		Instrument: requests.PaymentInstrument{
			Type: "card",
			Account: "9900009184214768",
			ExpMonth: 11,
			ExpYear: 2025,
			SecretCode: "111",
		},
	},
	Order: &requests.OrderData{
		OrderID: "R973i8Stza46n0me152oidgnr_492",
		Amount: 0,
		Currency: "RON",
		Description: "Sample order",
		Billing: requests.BillingShipping{
			Email: "client@example.com",
			Phone: "123456789",
			FirstName: "John",
			LastName: "Doe",
			City: "Bucharest",
			Country: 642,
		},
	Products: []requests.Product{
			{Name: "T-shirt Alfa", Code: "D276C05398EO14", Price: 17, Vat: 19},
		},
	},
}

startResp, err := client.StartPayment(startReq)

if  err != nil {
	fmt.Println("Error:", err)
	return
}

// Get the response from the API
fmt.Printf("Start payment response: %+v", startResp)

```

---

#### **GetStatus**

Retrieves the status of an order.

```go

statusResp, err := client.GetStatus("ntpID-123456", "orderID-12345")

if  err != nil {	
	fmt.Println("Error:", err)
	return
}

fmt.Printf("Order status: %+v", statusResp)

```
---

#### **VerifyAuth**

Handles 3D Secure verification for transactions.

```go

verifyResp, err := client.VerifyAuth("authToken123", "ntpID-123456", "paResData")

if  err != nil {
	fmt.Println("Error:", err)
	return
}

fmt.Printf("VerifyAuth response: %+v", verifyResp)

```
---

### **IPN Verification**

Validates the authenticity and integrity of IPNs.

```go

func  ipnHandler(w  http.ResponseWriter, r *http.Request) {
	result, err := client.VerifyIPN(r)
	if  err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("IPN Verification Result: %+v", result)
	w.WriteHeader(http.StatusOK)
}

```
---
## **Running tests**
To ensure the SDK functions as expected, comprehensive unit tests are provided. You can run unit tests suit using the following command:

```bash
cd project-dir
go test -v ./tests
```

---
## **Error Handling**

The SDK provides structured error handling with pre-defined error variables. Some common errors include:


| Error Name | Description |

| `ErrMissingAPIKey` | API Key is not provided. |

| `ErrInvalidPublicKey` | The provided public key is invalid. |

| `ErrInvalidIssuer` | JWT token issuer (`iss`) is not "NETOPIA Payments". |

| `ErrPayloadHashMismatch` | Hash of the payload does not match `sub` in the JWT.|

| `ErrInvalidToken` | The JWT token is invalid. |

 
 ### **Handle errors:** 

```go

if  err != nil {
	switch  err {
		case  netopia.ErrMissingAPIKey:
			fmt.Println("API Key is missing!")
		default:
			fmt.Println("An error occurred:", err)
	}
}

```
---

## **Examples**

Clone the sample repository and explore for complete use cases:

```bash

git clone  https://github.com/x/sample.git
cd  x/sample
go  run main.go

```
---

## **Useful Links**

- [NETOPIA Payments Documentation](https://doc.netopia-payments.com/)

- [API Reference](https://secure.sandbox.netopia-payments.com/spec)

- [Demo Cards for Sandbox](https://support.netopia-payments.com/en-us/article/52-carduri-de-test)

---