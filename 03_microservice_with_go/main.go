package main

import (
	"context"
	"fmt"

	"github.com/ds7cw/go_fundamentals/03_microservice_with_go/application"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start  app:", err)
	}
	// use 'curl -X POST localhost:3000/hello -v'
	// 'curl localhost:3000/orders' - GET
	// 'curl -X POST localhost:3000/orders' - POST

}
