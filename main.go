package main

import (
	"fmt"

	"github.com/jbattistella/capstone-project/engine"
)

func main() {

	var res = engine.AppResponse{}

	res = engine.Engine()

	if res.Msg1 != "" {
		fmt.Println(res.Msg1)
	}
	if res.Msg2 != "" {
		fmt.Println(res.Msg2)
	}
	if res.Msg3 != "" {
		fmt.Println(res.Msg3)
	}
	if res.Vegetables != nil {
		for veg := range res.Vegetables {
			fmt.Println(res.Vegetables[veg].CommonName)
		}
	}

	//set up router

	//add handler func

	//update DB funcs
}
