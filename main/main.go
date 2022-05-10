package main

import "fmt"

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
	// fmt.Println(cards)
}

func createDeck () deck{
	return deck {"Ace of Diamonds",
		"Joker of Hearts",
		"Five of Spade"}
}

func createCard() string{
	return "Five of Diamonds"
}