package handler

import (
	"go-prima/repository"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	store repository.SQLStore
}

func NewServer(app *fiber.App, store *repository.SQLStore) *fiber.App {
	server := &Server{
		store: *store,
	}

	application := server.setupRouter(app)
	return application
}

func (server *Server) setupRouter(app *fiber.App) *fiber.App {

	app.Get("/", GetUsersHandler)
	app.Post("/login", server.LoginHandler)
	app.Post("/registration", server.RegistrationHandler)
	// app.Post("/forgot-password", handler.ForgotPasswordHandler)

	// // POST
	// app.Get("/post/create", handler.CreatePostHandler)
	// app.Get("/post/delete", handler.DeletePostHandler)
	// app.Get("/post/update", handler.UpdatePostHandler)

	// // PROFILE
	// app.Get("/profile", handler.GetUsersHandler)
	// app.Get("/profile/update", handler.GetUsersHandler)

	return app
}
