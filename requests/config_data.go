package requests

import (
	"errors"
	"net/url"
)

type ConfigData struct {
	EmailTemplate string `json:"emailTemplate"`
	EmailSubject  string `json:"emailSubject"`
	NotifyURL     string `json:"notifyUrl"`
	RedirectURL   string `json:"redirectUrl"`
	Language      string `json:"language"`
}

func (c *ConfigData) Validate() error {
	if c.NotifyURL == "" || !isValidURL(c.NotifyURL) {
		return errors.New("notifyUrl is required and must be a valid URL")
	}
	if c.RedirectURL == "" || !isValidURL(c.RedirectURL) {
		return errors.New("redirectUrl is required and must be a valid URL")
	}
	if len(c.Language) != 2 {
		return errors.New("language must be 2 characters (ISO 639-1 code)")
	}
	return nil
}

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	return err == nil
}
