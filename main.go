package main

import "fmt"

func main() {

	cards := newDeck()
	// callig receiver function to shuffle & print the card deck
	cards.shuffle()
	fmt.Println("suffled cards..!")
	cards.print()

	// save cards into the file
	cards = newDeck()
	cards.saveToFile("my_cards")
	// read card deck from file
	fmt.Println("New card deck from file")
	Cards := newDeckFromFile("my_cards")
	Cards.print()

	// hand, remainingDeck := deal(Cards, 10)

}
