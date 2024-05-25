package model

type Request struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Creator string `json:"creator"`
}
