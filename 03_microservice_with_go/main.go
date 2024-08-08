package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/ds7cw/go_fundamentals/03_microservice_with_go/application"
)

func main() {
	app := application.New(application.LoadConfig())

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

	// curl -X POST -d '{"customer_id":"'$(uuidgen)'","line_items":[{"item_id":"'$(uuidgen)'","quantity":5,"price":1999}]}' localhost:3000/orders

	// SERVER_PORT=8080 go run main.go
	// 	curl localhost:8080
	// 	curl localhost:8080/orders | jq
	//  % Total    % Received % Xferd  Average Speed   Time    Time
	// 	Time  Current                  Dload  Upload   Total   Spent
	// 	Left  Speed
	// 	0     0    0     0    0     0      0      0 --:--:-- --:--:--
	// 	-100 12  100    12    0     0     37      0 --:--:-- --:--:--
	// 	--:--:--    37
}
