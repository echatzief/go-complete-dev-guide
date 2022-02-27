package main

import "fmt"

func main() {
	cards := newDeck()
	fmt.Println("------------------------ Initial Deck ------------------------ ")
	cards.print()
	firstDeck, secDeck := cards.deal(4)
	fmt.Println("------------------------ First Deck ------------------------ ")
	firstDeck.print()
	fmt.Println("------------------------ Second Deck ------------------------ ")
	secDeck.print()
	secDeck.save("secDeck.deck")
    fmt.Println("------------------------ Third Deck ------------------------ ")
    thirdDeck := deckFromFile("secDeck.deck")
	thirdDeck.print()
	fmt.Println("------------------------ Third Deck(Shuffled) ------------------------ ")
	thirdDeck.shuffle()
	thirdDeck.print()
}
