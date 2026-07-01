package main

import (
	"fmt"

	"github.com/RyanDev-21/pkg/projector"
)

func main() {
	opts, err := projector.GetOpts()
	if err != nil {
		panic(err)
	}
	config, err := projector.NewConfig(opts)
	if err != nil {
		panic(err)
	}

	fmt.Printf("opts:%+v", opts)
	fmt.Printf("config:%+v", config)
}
