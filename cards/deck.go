package main

import (
	"fmt"
)

type deck []string

func (cards deck) print() {
	for idx, card := range cards {
		fmt.Println(idx, ": ", card)
	}
}
