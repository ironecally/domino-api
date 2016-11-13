package domino

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck : struct for implementing cards in deck
type Deck struct {
	cards []Card
}

// DominoDeck : static array of one set complete domino deck
var DominoDeck = []CardValue{BlankBlank, BlankOne, BlankTwo, BlankThree, BlankFour, BlankFive, BlankSix, OneOne, OneTwo, OneThree, OneFour, OneFive, OneSix, TwoTwo, TwoThree, TwoFour, TwoFive, TwoSix, ThreeThree, ThreeFour, ThreeFive, ThreeSix, FourFour, FourFive, FourSix, FiveFive, FiveSix, SixSix}

// NewDeck : function for creating a new deck
func NewDeck(shuffle bool) Deck {
	deck := NewDeckSpecified(shuffle, DominoDeck)
	return deck
}

// NewDeckSpecified : create deck with array of card
func NewDeckSpecified(shuffle bool, cardValues []CardValue) Deck {
	cards := make([]Card, len(cardValues))
	for cindex, card := range cardValues {
		cards[cindex] = MakeCard(card)
	}
	deck := Deck{cards}
	return deck
}

// NumberOfCards : method to get the number of card that is still available
func (d *Deck) NumberOfCards() int {
	return len(d.cards)
}

// GetCard : method for get the card in the index
func (d *Deck) GetCard(index int) *Card {
	return &d.cards[index]
}

func (d Deck) String() string {
	str := ""
	for _, card := range d.cards {
		str += fmt.Sprint(card) + "\n"
	}
	return str
}

// MultiShuffle : method shuffling the card as x manytimes
func (d *Deck) MultiShuffle(iterations int) {
	for i := 0; i < iterations; i++ {
		d.Shuffle()
	}
}

// Shuffle : method for suffling the card one times
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	N := len(d.cards)
	for i := 0; i < N; i++ {
		r := i + rand.Intn(N-i)
		d.cards[r], d.cards[i] = d.cards[i], d.cards[r]
	}
}
