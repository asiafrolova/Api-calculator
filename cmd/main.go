package main

import (
	"github.com/asiafrolova/Api-calculator/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()

}
