package main

import (
	"chash/consistent"
	"fmt"
)

func main() {

	c := consistent.New(4, nil)

	c.Add("aaaaa-1")
	c.Add("aaaaa-2")
	c.Add("aaaaa-3")

	for i := 0; i < 50; i++ {
		fmt.Println(c.Get(fmt.Sprintf("%d", i)))
	}
}
