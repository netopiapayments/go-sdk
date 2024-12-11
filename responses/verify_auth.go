package responses

type VerifyAuthResponse struct {
	ApiResponse
	Payment *Payment          `json:"payment,omitempty"`
	Error   *ErrorWithDetails `json:"error,omitempty"`
}
