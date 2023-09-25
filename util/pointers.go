package util

import "fmt"

type Character struct {
	HP int
}

func (character *Character) TakeDamage(amount int) {
	character.HP -= amount
	fmt.Println("Damage Taken:", character.HP)
}
