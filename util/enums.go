package util

import "fmt"

func getDamageType(weaponType string) int {
	switch weaponType {
	case "axe":
		return 90
	case "sword":
		return 80
	case "mace":
		return 70
	default:
		panic("Unknown weapon type")
	}
}

func enums() {
	var weaponType string = "axe"
	var damage int = getDamageType(weaponType)
	fmt.Println(damage)
}
