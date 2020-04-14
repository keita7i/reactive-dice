package rest

type DiceResponse struct {
	Name  string   `json:"name"`
	Faces []string `json:"faces"`
}
