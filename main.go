package main

import (
	"basics/util"
	"fmt"
)

func main() {
	m := util.NewCustomMap[int, string]()
	m.Insert(1, "one")
	m.Insert(2, "two")
	fmt.Println(m)
}
