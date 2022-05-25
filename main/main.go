package main

import (
	"fmt"
	"deck"
)

/*** 
go standard library
https://pkg.go.dev/std
***/

type book string;

func (b book) printTitle(){
	fmt.Println(b)
}


func main(){
	cards := newDeck();
	cards.printAll();
	fmt.Println("++++++++++++++++++++++")
	hand, remainingDeck := cards.deal(5);
	fmt.Println("in hand: ")
	hand.printAll();
	fmt.Println("remaining: ")
	remainingDeck.printAll();

	hand.saveToFile("handDeck.txt")

	deckFromFile := newDeckFromFile("handDeck.txt")
	fmt.Println("deck from file:", deckFromFile)
	// fmt.Println(cards)
	deckFromFile.shuffle()
	fmt.Println("shuffled deck from file: ", deckFromFile)
}

func createDeck () deck{
	return deck {"Ace of Diamonds",
		"Joker of Hearts",
		"Five of Spade"}
}

func createCard() string{
	return "Five of Diamonds"
}