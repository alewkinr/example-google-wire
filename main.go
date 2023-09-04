package main

import "github.com/alewkinr/example-google-wire/internal"

func main() {
	message := internal.NewMessage()
	greeter := internal.NewGreeter(message)
	event := internal.NewEvent(greeter)

	event.Start()
}
