package models

type Voting struct {
	ID        string
	ProjectID string
	PoliticID string
	Vote      string // ex: "Sim", "Não", "Abstenção"
	Data      string
}
