package main

import "Yacalc/internal/application"

func main() {
	err := application.RunServer()
	if err != nil {
		panic(err)
	}
}
