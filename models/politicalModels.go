package models

type PoliticModels struct {
	ID            string
	Nome          string
	Partido       string
	Cargo         string
	Estado        string
	Municipio     string
	MandatoInicio int
	MandatoFim    int
	Email         string
	RedesSociais  map[string]string // ex: {"instagram": "insta.com/nome"}
	Projetos      []BillModels
	Votacoes      []Voting
	Gastos        []PublicSpending
}
