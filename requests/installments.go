package requests

type Installments struct {
	Selected  int   `json:"selected"`
	Available []int `json:"available"`
}
