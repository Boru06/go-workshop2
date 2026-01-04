package models

type Register struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	LineID   string `json:"line_id"`
	Tel      string `json:"tel"`
	Business string `json:"business"`
	Website  string `json:"website"`
}