package main

import "fmt"

type Player struct {
	name   string
	health int
	attack float64
}

func (player Player) getHealth() int {
	return player.health
}

func demo() {
	// player := Player{
	// 	name:   "Gopher",
	// 	health: 100,
	// 	attack: 1.5,
	// } // infer Player type

	// users: make(map[string]int)

	users := map[string]int{
		"John": 22,
	}
	users["Tim"] = 23

	// age, ok := users["Tim"]
	// if ok {
	// 	fmt.Printf("Tim is %d years old\n", age)
	// } else {
	// 	fmt.Println("Tim not found")
	// }

	for key, value := range users {
		fmt.Printf("%s is %d years old\n", key, value)
	}

	// fmt.Printf("Users: %+v\n", users) // %+v prints the struct

	// delete(users, "Tim") // delete works only for map

	// fmt.Printf("Health: %+v\n", player.getHealth())
}
