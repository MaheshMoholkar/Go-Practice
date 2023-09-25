package main

import (
	"basics/util"
	"fmt"
)

func main() {
	character := &util.Character{HP: 100}
	fmt.Println("Character HP:", character.HP)
	fmt.Println(character)
	fmt.Println(*character)
	character.TakeDamage(10)
	fmt.Println("Character HP:", character.HP)
}
