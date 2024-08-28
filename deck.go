package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create new type of 'deck' which consist a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spade", "Diamond", "Heart", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handValue int) (deck, deck) {
	return d[:handValue], d[handValue:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Log the error and quit the program
		fmt.Println("Error : ", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ", ")
	return deck(s)
}

func (d deck) saveDealtCardsToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDealtCardsFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		// Handle error
		panic(err)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)

		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
