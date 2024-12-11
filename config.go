package netopia

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
