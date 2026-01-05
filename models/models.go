package models

import (
	

	"gorm.io/gorm"
)

type Register struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
	LineID   string `json:"line_id"`
	Tel      string `json:"tel"`
	Business string `json:"business"`
	Website  string `json:"website"`
}

type Dogs struct {
	gorm.Model
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
}

type DogsRes struct {
	Name  string `json:"name"`
	DogID int    `json:"dog_id"`
	Type  string `json:"type"`
}

type ResultData struct {
	Count int         `json:"count"`
	Data  []DogsRes `json:"data"`
	Name  string      `json:"name"`
	Red     int `json:"Sum_red"`
	Green   int `json:"Sum_green"`
	Pink    int `json:"Sum_pink"`
	NoColor int `json:"Sum_nocolor"`
}


type Company struct {
	gorm.Model
	Name string `json:"company_name"`
	Address string `json:"address"`
	Tel string `json:"tel"`
	Email string `json:"email"`
	Business string `json:"business"`
}
