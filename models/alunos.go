package models

import (
	"gorm.io/gorm"
)

type Alunos struct {
	gorm.Model `json:"gorm_model"`
	Id         	uint    `json:"id",gorm:"primaryKey"`
	Nome       	string  `json:"nome,omitempty"`
	CPF        	string  `json:"cpf,omitempty"`
	Email		string	`json:"email"`
	Nascimento 	string  `json:"nascimento,omitempty"`
	Plano      	string  `json:"plano,omitempty"`
	Valor      	float64 `json:"valor,omitempty"`
}

//Fazer a Query com o banco de dados
