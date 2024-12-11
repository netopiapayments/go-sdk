package netopia

import (
	"fmt"
	"netopia/requests"
	"netopia/responses"
)

func (c *PaymentClient) StartPayment(req *requests.StartPaymentRequest) (*responses.StartPaymentResponse, error) {
	if req.Order == nil {
		return nil, ErrInvalidOrder
	}

	req.Order.PosSignature = c.cfg.PosSignature
	req.Config.NotifyURL = c.cfg.NotifyURL
	req.Config.RedirectURL = c.cfg.RedirectURL

	if req.Config.Language == "" {
		req.Config.Language = "ro"
	}

	url := fmt.Sprintf("%s/payment/card/start", c.BaseURL())
	return sendJSON[responses.StartPaymentResponse](url, c.cfg.ApiKey, req)
}

func (c *PaymentClient) GetStatus(ntpID, orderID string) (*responses.StatusResponse, error) {
	req := requests.StatusRequest{
		PosID:   c.cfg.PosSignature,
		NtpID:   ntpID,
		OrderID: orderID,
	}

	url := fmt.Sprintf("%s/operation/status", c.BaseURL())
	return sendJSON[responses.StatusResponse](url, c.cfg.ApiKey, req)
}

func (c *PaymentClient) VerifyAuth(authToken, ntpID, paRes string) (*responses.VerifyAuthResponse, error) {
	req := requests.VerifyAuthRequest{
		AuthenticationToken: authToken,
		NtpID:               ntpID,
		FormData:            map[string]string{"paRes": paRes},
	}

	url := fmt.Sprintf("%s/payment/card/verify-auth", c.BaseURL())
	return sendJSON[responses.VerifyAuthResponse](url, c.cfg.ApiKey, req)
}
