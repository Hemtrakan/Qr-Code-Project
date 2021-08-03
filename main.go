package main

import (
	"log"
	"os"
	"os/signal"
	"qrcode/access"
	"qrcode/control"
	"qrcode/environment"
	"qrcode/present"
)

func main() {

	// -- Build Environment
	prop := environment.Build()
	if prop == nil {
		log.Panic("environment not exist")
	}

	// Init Access
	acc := access.Initial(prop)

	sign := make(chan os.Signal, 1)
	signal.Notify(sign, os.Interrupt, os.Kill)

	go present.APICreate(control.APICreate(acc))

	<-sign
}