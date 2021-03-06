package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// create a new type of 'deck	'.
// which is slice of string

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Queen", "Jack", "King"}

	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suite)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	// file := ""
	// for _, str := range d {
	// 	file += str + "\n"
	// }

	return strings.Join([]string(d), "\n")
}

func (d deck) saveDeckToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	deck, err := os.ReadFile(filename)
	// os.Remove(filename)
	if err != nil {
		// log the error and return a call to newDeck()
		// OR
		// Log the error and quit the program
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	return strings.Split(string(deck), "\n")
}

func (d deck) shuffle() {
	for i := range d {
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
