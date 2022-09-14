package router

import (
	"agenda-personal/models"
	"github.com/gofiber/fiber/v2"
)

func IniciandoRotas(app *fiber.App) {
	perfil := app.Group("/aluno")
	perfil.Get("/", models.BuscandoAlunos)
	perfil.Get("/:id", models.BuscandoAluno)
	perfil.Post("/:id", models.CriandoAluno)
	perfil.Put("/:id", models.EditandoAluno)
	perfil.Delete("/:id", models.ApagandoAluno)
}