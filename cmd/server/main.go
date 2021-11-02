package main

import "fmt"

type App struct{}

func (app *App) Run() error {
	fmt.Println("Setting up the app")
	return nil
}

func main() {
	fmt.Printf("Go Rest API test")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting up the Rest API")
		fmt.Println(err)
	}
}
