package main

import "fmt"

func main() {

	d := Deck{}
	d.shuffle()
	fmt.Print(d)
}

type cardSuit string

const (
	spades   cardSuit = "spades"
	hearts   cardSuit = "hearts"
	clubs    cardSuit = "clubs"
	diamonds cardSuit = "diamonds"
)

type cardValue string

const (
	Six   cardValue = "Six"
	Seven cardValue = "Seven"
	Eight cardValue = "Eight"
	Nine  cardValue = "Nine"
	Ten   cardValue = "Ten"
	Jack  cardValue = "Jack"
	Queen cardValue = "Queen"
	King  cardValue = "King"
	Ace   cardValue = "Ace"
)

var mapToValue = map[cardValue]int{
	Six: 6, Seven: 7, Eight: 8, Nine: 9, Ten: 10, Jack: 10, Queen: 10, King: 10, Ace: 11,
}

type Card struct {
	Suit  cardSuit
	Value cardValue
}

type Deck []Card

func (d *Deck) shuffle() {
	*d = Deck{}

	for _, suit := range Suits {

		for _, val := range Values {
			*d = append(*d, Card{suit, val})
		}
	}
}

var Suits []cardSuit = []cardSuit{spades, hearts, diamonds, clubs}
var Values []cardValue = []cardValue{Six, Seven, Eight, Nine, Ten, Jack, Queen, King, Ace}
