package main

import (
	cf "book/config"
	"book/app"
)

func main() {
	config := cf.Load()

	app.Run(config)
}
