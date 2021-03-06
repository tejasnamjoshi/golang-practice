package main

func main() {
	cards := newDeck()
	hand, _ := deal(cards, 10)
	hand.print()
	hand.shuffle()
	hand.print()

	// hand.print()
	// remainingDeck.print()
	// fmt.Println(cards.toString())
	// hand.saveDeckToFile("my_deck")
	// newHand := newDeckFromFile("my_deck")
	// newHand.print()

}

// 31 - Random Number Generation
