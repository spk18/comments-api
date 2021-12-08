package main

import (
	"fmt"
	"net/http"

	transportHTTP "github.com/spk18/comments-api/internal/transport/http"
)

// App - the struct which contains all the pointers to databases etc
type App struct {}


// Run - sets up our application
func (a *App) Run() error {

	fmt.Println("Setting up App")

	handler := transportHTTP.NewHandler()
	handler.SetupRoutes()

	if err := http.ListenAndServe(":3000", handler.Router); err != nil {
		return err
	}
	

	return nil
}

func main() {
	
	fmt.Println("Hello")

	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error starting rest api")
		fmt.Println(err)
	}

}