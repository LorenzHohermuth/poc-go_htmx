package main

import (
	"lorenz/go-htmx/internal/model"
	"lorenz/go-htmx/internal/routes"
)

func main() {
	model.Setup()
	routes.SetupAndRun()
}
