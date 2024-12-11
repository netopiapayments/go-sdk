package requests

type VerifyAuthRequest struct {
	AuthenticationToken string            `json:"authenticationToken"`
	NtpID               string            `json:"ntpID"`
	FormData            map[string]string `json:"formData"`
}
