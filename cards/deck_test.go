package main

import (
    "testing"
    "os"
)

func TestNewDeck(t *testing.T){
    d := newDeck()

    if len(d) != 16 {
        t.Error("Expected deck length 16, but got ",len(d))
    }

    if d[0] != "Ace of Spades"{
        t.Error("Expected first card of Ace of Spades, but got ",d[0])
    }

    if d[len(d) - 1] != "Four of Clubs"{
        t.Error("Expected last card of Four of Clubs, but got ",d[len(d) - 1])
    }
}

func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T){
    os.Remove("_decktesting")

    d := newDeck()
    d.save("_deck")

    loadedDeck := deckFromFile("_deck")

    if len(loadedDeck) != 16 {
        t.Error("Expected 16 cards in deck, but got ",len(loadedDeck))
    }

    os.Remove("_deck")
}