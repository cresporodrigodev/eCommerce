package main

import (
	"github.com/cresporodrigodev/eCommerce/app/backend/main/boostrap"
	"log"
)

func main() {
	if err := boostrap.RunServer(); err != nil {
		log.Fatal(err)
	}
}
