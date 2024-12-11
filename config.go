package netopia

import "fmt"

type Config struct {
	PosSignature    string
	ApiKey          string
	IsLive          bool
	NotifyURL       string
	RedirectURL     string
	PublicKey       []byte
	ActiveKey       string
	PosSignatureSet []string
	HashMethod      string
}

func (c *Config) Validate() error {
	if c.ApiKey == "" {
		return fmt.Errorf("missing apiKey")
	}
	if c.PosSignature == "" {
		return fmt.Errorf("missing posSignature")
	}
	if len(c.PosSignatureSet) == 0 {
		return fmt.Errorf("posSignatureSet must not be empty")
	}
	if c.NotifyURL == "" || c.RedirectURL == "" {
		return fmt.Errorf("both notifyURL and redirectURL must be set")
	}
	if len(c.PublicKey) == 0 {
		return fmt.Errorf("publicKey must be provided")
	}
	if c.ActiveKey == "" {
		return fmt.Errorf("activeKey must be provided")
	}
	if c.HashMethod == "" {
		c.HashMethod = "sha512" // implicit
	}
	return nil
}
