package main

import (
	"fmt"

	"github.com/GiovannaK/gopportunities/config"
	"github.com/GiovannaK/gopportunities/router"
)

func main() {
	err := config.Init()

	if err != nil {
		fmt.Println(err)
		return
	}

	router.Initialize()
}
