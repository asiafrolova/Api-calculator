package main

import (
	"github.com/asiafrolova/Api-calculator/internal/application"
	//"github.com/asiafrolova/Calculator/rpn/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()

}
