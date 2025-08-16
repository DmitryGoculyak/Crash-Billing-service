package main

import "Crash-Billing-service/internal/container"

func main() {
	app := container.Build()

	app.Run()
}
