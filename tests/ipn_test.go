package tests

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha512"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"math/big"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	netopia "github.com/netopiapayments/go-sdk"
)

func TestVerifyIPN(t *testing.T) {
	privateKey, certPEM, err := GenTestKeyPair()
	if err != nil {
		t.Fatalf("failed to generate test keypair: %v", err)
	}

	cfg := netopia.Config{
		PosSignature:    "TEST-POS-SIG",
		ApiKey:          "fake-api-key",
		IsLive:          false,
		NotifyURL:       "https://example.com/ipn",
		RedirectURL:     "https://example.com/return",
		PublicKey:       certPEM,
		PosSignatureSet: []string{"TEST-POS-SIG"},
		HashMethod:      "sha512",
	}
	logger := &netopia.DefaultLogger{}

	client, err := netopia.NewPaymentClient(cfg, logger)
	if err != nil {
		t.Fatalf("failed to init PaymentClient: %v", err)
	}

	ipnBody := map[string]interface{}{
		"payment": map[string]interface{}{
			"status": 3,
		},
	}
	bodyBytes, _ := json.Marshal(ipnBody)

	verificationToken, err := signIPNBodyWithRSA(bodyBytes, "TEST-POS-SIG", privateKey)
	if err != nil {
		t.Fatalf("failed to sign IPN body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/ipn", bytes.NewReader(bodyBytes))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Verification-token", verificationToken)

	result, err := client.VerifyIPN(req)
	if err != nil {
		t.Fatalf("VerifyIPN failed: %v", err)
	}

	if result.Status != 3 {
		t.Errorf("Expected status=3, got %d", result.Status)
	}
	t.Logf("VerifyIPN OK! Status=%d, Message=%s", result.Status, result.Message)
}

func GenTestKeyPair() (*rsa.PrivateKey, []byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}
	template := x509.Certificate{
		SerialNumber: big.NewInt(1),
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(time.Hour),
		Subject:      pkix.Name{CommonName: "Self-Signed Test"},
		KeyUsage:     x509.KeyUsageDigitalSignature,
	}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return nil, nil, err
	}
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})
	return privateKey, certPEM, nil
}

func signIPNBodyWithRSA(body []byte, posSignature string, privateKey *rsa.PrivateKey) (string, error) {
	bodySum := sha512.Sum512(body)
	bodyHash := base64.StdEncoding.EncodeToString(bodySum[:])

	now := time.Now()
	claims := jwt.MapClaims{
		"iss": "NETOPIA Payments",
		"aud": []string{posSignature},
		"sub": bodyHash,
		"iat": now.Unix(),
		"nbf": now.Unix(),
		"exp": now.Add(5 * time.Minute).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	return token.SignedString(privateKey)
}
