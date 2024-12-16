package netopia

import (
	"crypto/sha512"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	STATUS_NEW                  = 1
	STATUS_OPENED               = 2
	STATUS_PAID                 = 3
	STATUS_CANCELED             = 4
	STATUS_CONFIRMED            = 5
	STATUS_PENDING              = 6
	STATUS_SCHEDULED            = 7
	STATUS_CREDIT               = 8
	STATUS_CHARGEBACK_INIT      = 9
	STATUS_CHARGEBACK_ACCEPT    = 10
	STATUS_ERROR                = 11
	STATUS_DECLINED             = 12
	STATUS_FRAUD                = 13
	STATUS_PENDING_AUTH         = 14
	STATUS_3D_AUTH              = 15
	STATUS_CHARGEBACK_REPRESENT = 16
	STATUS_REVERSED             = 17
	STATUS_PENDING_ANY          = 18
	STATUS_PROGRAMMED_RECURRENT = 19
	STATUS_CANCELED_PROGRAMMED  = 20
	STATUS_TRIAL_PENDING        = 21
	STATUS_TRIAL                = 22
	STATUS_EXPIRED              = 23
)

type IPNVerificationResult struct {
	ErrorType    int             `json:"errorType"`
	ErrorCode    int             `json:"errorCode"`
	ErrorMessage string          `json:"errorMessage"`
	Payload      json.RawMessage `json:"payload"`
	Status       int             `json:"status,omitempty"`
	Message      string          `json:"message,omitempty"`
}

type IPNData struct {
	Payment struct {
		Status int `json:"status"`
	} `json:"payment"`
}

func (c *PaymentClient) VerifyIPN(r *http.Request) (*IPNVerificationResult, error) {
	verificationToken := r.Header.Get("Verification-token")
	if verificationToken == "" {
		return nil, ErrMissingVerification
	}

	parts := strings.Split(verificationToken, ".")
	if len(parts) != 3 {
		return nil, ErrWrongVerificationToken
	}

	publicKey, err := x509.ParsePKIXPublicKey(c.cfg.PublicKey)
	if err != nil {
		return nil, ErrInvalidPublicKey
	}

	bodyData, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	token, err := jwt.Parse(verificationToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, ErrUnexpectedSigningToken
		}
		return publicKey, nil
	}, jwt.WithLeeway(0), jwt.WithTimeFunc(time.Now))
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	if claims["iss"] != "NETOPIA Payments" {
		return nil, ErrInvalidIssuer
	}

	audVal := claims["aud"]
	var actualAud string
	switch v := audVal.(type) {
	case []interface{}:
		if len(v) == 0 {
			return nil, ErrEmptyAudience
		}
		firstAud, _ := v[0].(string)
		if firstAud == "" {
			return nil, ErrEmptyAudience
		}
		actualAud = firstAud
	case string:
		actualAud = v
	default:
		return nil, ErrEmptyAudience
	}

	if actualAud != c.cfg.PosSignature {
		return nil, ErrInvalidAudience
	}

	found := false
	for _, ps := range c.cfg.PosSignatureSet {
		if ps == actualAud {
			found = true
			break
		}
	}

	if !found {
		return nil, ErrAudienceNotInSet
	}

	payloadHash, err := computeHash(c.cfg.HashMethod, bodyData)
	if err != nil {
		return nil, err
	}
	sub, _ := claims["sub"].(string)
	if payloadHash != sub {
		return nil, ErrPayloadHashMismatch
	}

	var ipnData IPNData

	if err := json.Unmarshal(bodyData, &ipnData); err != nil {
		return nil, ErrFailedPayloadParsing
	}

	result := &IPNVerificationResult{
		ErrorType:    0,
		ErrorCode:    0,
		ErrorMessage: "",
		Payload:      bodyData,
		Status:       ipnData.Payment.Status,
		Message:      statusMessage(ipnData.Payment.Status),
	}

	return result, nil
}

func computeHash(hashMethod string, data []byte) (string, error) {
	switch strings.ToLower(hashMethod) {
	case "sha512":
		sum := sha512.Sum512(data)
		return base64.StdEncoding.EncodeToString(sum[:]), nil
	default:
		return "", ErrUnsupportedHashMethod
	}
}

func statusMessage(status int) string {
	switch status {
	case STATUS_NEW, STATUS_PAID, STATUS_CONFIRMED:
		return "payment was confirmed; deliver goods"
	case STATUS_CREDIT:
		return "a previously confirmed payment was refunded; cancel goods delivery"
	case STATUS_CANCELED:
		return "payment was cancelled; do not deliver goods"
	case STATUS_PENDING_AUTH:
		return "update payment status, last modified date&time in your system"
	case STATUS_FRAUD:
		return "payment in reviewing"
	default:
		return "no specific action"
	}
}
