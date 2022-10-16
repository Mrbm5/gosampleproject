package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of 'deck'

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// receiver function
// any variable of type "deck" now gets access to the print method
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Multiple value return from a function
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// receiver function
// any variable of type "deck" now gets access to the print method
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// receiver function
// any variable of type "deck" now gets access to the print method
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error::", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for index := range d {
		newPosition := r.Intn(len(d) - 1)
		d[index], d[newPosition] = d[newPosition], d[index]
	}
}
