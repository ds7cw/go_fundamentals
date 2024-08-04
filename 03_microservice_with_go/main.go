package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/ds7cw/go_fundamentals/03_microservice_with_go/application"
)

func main() {
	app := application.New()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	// cancel called at the end of the main func execution
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start  app:", err)
	}
	// use 'curl -X POST localhost:3000/hello -v'
	// 'curl localhost:3000/orders' - GET
	// 'curl -X POST localhost:3000/orders' - POST

	// Error Message when redis server is not running:
	//  failed to start  app: failed to connect to redis:
	//  No connection could be made because the target machine
	//  actively refused it.

	// redis-cli/ redis-cli shutdown
	// sudo service redis-server stop/ start/ restart

}
