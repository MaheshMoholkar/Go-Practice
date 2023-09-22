package main

import (
	"basics/util"
	"fmt"
)

func main() {
	var color util.Color
	var result = color.GetColor(2)
	fmt.Println(result)
}
