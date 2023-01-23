package main

import (
	"github.com/jbattistella/capstone-project/server"
)

//using flag to add cli functionality, how can CLI be used to add quality and value to program

// var flagvar string

// func init() {
// 	flag.StringVar(&flagvar, "flagname", "", "help message for flagname")

// 	flag.Parse()

// 	if flagvar != "" {
// 		fmt.Printf("flagvar: %v\n", flagvar)
// 	}
// }

func main() {

	server.GardenAPI()

}
