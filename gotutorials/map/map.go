package main

import (
	"fmt"
)

func main() {
	menu := map[string]float64{
		"soup":  4.99,
		"pie":   7.99,
		"salad": 6.99,
	}
	fmt.Println(menu)
	fmt.Println(menu["pie"])

	for k, v := range menu {
		fmt.Println(k, "-", v)
	}

	phoneBook := map[int]string{
		267584967: "mario",
		984759373: "luigi",
		845775485: "peach",
	}
	fmt.Println(phoneBook)

}
