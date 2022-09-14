package router

import (
	"agenda-personal/controllers"
	"github.com/gofiber/fiber/v2"
)

func IniciandoRotas(app *fiber.App) {
	perfil := app.Group("/aluno")
	perfil.Get("/", controllers.BuscandoAlunos)
	perfil.Get("/:id", controllers.BuscandoAluno)
	perfil.Post("/:id", controllers.CriandoAluno)
	perfil.Put("/:id", controllers.EditandoAluno)
	perfil.Delete("/:id", controllers.ApagandoAluno)
}