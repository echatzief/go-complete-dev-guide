package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

type deck []string

func (cards deck) print() {
	for idx, card := range cards {
		fmt.Println(idx, ": ", card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) deal(handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string{
    return strings.Join([]string(d), ";");
}

func (d deck) save(fileName string) error{
    return ioutil.WriteFile(fileName,[]byte(d.toString()), 0644)
}

func deckFromFile(fileName string) deck{
    content, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ",err)
		os.Exit(1)
	}
	return deck(strings.Split(string(content),";"))
}

func (d deck) shuffle() {
    source := rand.NewSource(time.Now().UnixNano())
    r := rand.New(source)

    for i:= range d{
        newPos := r.Intn(len(d) - 1)
        d[i], d[newPos] = d[newPos], d[i]
    }
}