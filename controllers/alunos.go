package controllers

import "github.com/gofiber/fiber/v2"

//Fazer a conexão com o front

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