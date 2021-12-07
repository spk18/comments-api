package main

import "fmt"


// App - the struct which contains all the pointers to databases etc 
type App struct{

}

// Run - sets up our application
func (a *App) Run() error {
	fmt.Println("Setting up App")
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