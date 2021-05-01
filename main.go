package main

import (
	"log"

	app "huddlet/app"
)

func main() {
	err := app.Run()
	if err != nil {
		log.Fatal(err.Error())
	}
}
