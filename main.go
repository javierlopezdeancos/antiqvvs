package main

import (
	"flag"
	"fmt"

	"github.com/javierlopezdeancos/antiqvvs/delete"
	"github.com/javierlopezdeancos/antiqvvs/environment"
	"github.com/javierlopezdeancos/antiqvvs/populate"
	"github.com/javierlopezdeancos/antiqvvs/update"
)

// Actions param from cli
var Actions = map[string]string{
	"delete":   "delete",
	"populate": "populate",
	"update":   "update",
}

// Resources param from cli
var Resources = map[string]string{
	"wines":  "wines",
	"prices": "prices",
	"wine":   "wine",
	"price":  "price",
}

func main() {
	err := environment.LoadEnv()

	if err != nil {
		panic(fmt.Sprintf("🔴 [ERROR] loading .env: %v", err))
	}

	action := flag.String("action", "", "a string")
	resource := flag.String("resource", "", "a string")

	flag.Parse()

	if *action == Actions["delete"] {
		fmt.Println("\n🔵 [INFO] Starting delete all wines in stripe...")

		err = delete.Start()

		if err != nil {
			fmt.Println(err)
		}
	} else if *action == Actions["populate"] {
		fmt.Println("\n🔵 [INFO] Starting populate products from JSON...")

		if *resource == Resources["prices"] {
			err = populate.CreatePrices()

			if err != nil {
				fmt.Println(err)
			}

			return
		} else if *resource == Resources["wines"] {
			err = populate.CreateWines()

			if err != nil {
				fmt.Println(err)
			}

			return
		}

		err = populate.Start()

		if err != nil {
			fmt.Println(err)
		}
	} else if *action == Actions["update"] {
		if *resource == Resources["price"] {
			err = update.Price()

			if err != nil {
				fmt.Println(err)
			}

			return
		} else if *resource == Resources["wine"] {
			err = update.Wine()

			if err != nil {
				fmt.Println(err)
			}

			return
		}

		err = update.Start()

		if err != nil {
			fmt.Println(err)
		}
	} else {
		fmt.Println("🔴 [ERROR] You need pass an action as param like populate or delete")
	}
}
