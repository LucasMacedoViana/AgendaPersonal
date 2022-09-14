package models

import (
	"github.com/gofiber/fiber/v2"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Alunos struct {
	gorm.Model
	Id 			uint	`gorm:"primaryKey"`
	Nome 		string
	CPF 		string
	Nascimento 	string
	Plano 		string
	Valor 		float64
}

//Fazer a Query com o banco de dados
func BuscandoAlunos(c *fiber.Ctx) error {
	return c.SendString("Buscando todos os alunos")
}

func BuscandoAluno(c *fiber.Ctx) error {
	return c.SendString("Buscando um aluno")
}

func CriandoAluno(c *fiber.Ctx) error {
	return c.SendString("Criando um aluno")
}

func EditandoAluno(c *fiber.Ctx) error {
	return c.SendString("Editando um aluno")
}

func ApagandoAluno(c *fiber.Ctx) error {
	return c.SendString("Excluindo um aluno")
}