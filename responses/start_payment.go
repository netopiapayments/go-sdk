package responses

type StartPaymentResponse struct {
	ApiResponse
	Payment        *Payment          `json:"payment,omitempty"`
	CustomerAction *Action           `json:"customerAction,omitempty"`
	Error          *ErrorWithDetails `json:"error,omitempty"`
}

type Action struct {
	Type                string     `json:"type,omitempty"`
	URL                 string     `json:"url,omitempty"`
	AuthenticationToken string     `json:"authenticationToken,omitempty"`
	FormData            Attributes `json:"formData,omitempty"`
}
