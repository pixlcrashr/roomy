package web

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

//go:embed dist/*
var app embed.FS

func RegisterRoutes(router fiber.Router) {
	distFS, _ := fs.Sub(app, "dist")

	router.Use("/", filesystem.New(filesystem.Config{
		Root:         http.FS(distFS),
		Index:        "index.html",
		NotFoundFile: "index.html",
	}))
}
