package requests

type StatusRequest struct {
	PosID   string `json:"posID"`
	NtpID   string `json:"ntpID"`
	OrderID string `json:"orderID"`
}
