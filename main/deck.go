package main

import (
	"fmt"
	"strings"
	"io/ioutil"
)

/***
Go is not OO language, there's no concept of class
custom defined type deck, it is an extension of slice string
***/
type deck []string;

/***
don't use receiver function in newDeck
this is to create new type of deck
receiver function is to accept and work with existing type
***/
func newDeck () deck {
	//declare empty deck
	cards := deck{}
	cardSuits := []string {"Club", "Diamond", "Heart", "Spade"}
	cardValue := []string {"Ace", "Two", "Three", "Joker", "Queen", "King"}
	//loop through suites and value
	//looping index i/j is not used, only required for syntax, declare the unused var as "_" to inform compiler ignore it
	//in Go, any unused variable will cause compile error
	for _, suite := range cardSuits {
		for _, value := range cardValue {
			//append(source, value) function will append the value to last of the source array and return a new array
			cards = append(cards, value + " of " + suite)
		}
	}
	return cards
}

/***
Go allow multiple return values, separated by ","
***/
func (dk deck) deal(handSize int) (deck, deck){
	return dk[:handSize], dk[handSize:]
}

func (dk deck) tostring() string {
	return strings.Join(dk, ",");
}

func (dk deck) saveToFile(fileName string) {
	ioutil.WriteFile(fileName, []byte(dk.tostring()), 0666);
}

/***
receiver function:
printAll function will receive type deck
all variables of type deck automatically have the ability to use this function 
***/
func (dk deck) printAll() {
	for i, d := range dk {
		fmt.Println(i, d)
	}
}