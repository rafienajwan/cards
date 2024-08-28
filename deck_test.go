package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spade" {
		t.Errorf("Expected first card of Ace of Spade, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Club" {
		t.Errorf("Expected last card of King of Club, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}

func TestShuffleDeck(t *testing.T) {
	deck := newDeck()
	deck.shuffle()

	if deck[0] == "Ace of Spades" && deck[len(deck)-1] == "King of Clubs" {
		t.Errorf("Deck does not appear to be shuffled")
	}
}

func TestDeal(t *testing.T) {
	deck := newDeck()

	hand, remainingDeck := deal(deck, 5)

	if len(hand) != 5 {
		t.Errorf("Expected 5 cards in hand, got %v", len(hand))
	}

	if len(remainingDeck) != 47 {
		t.Errorf("Expected 47 cards in remaining deck, got %v", len(remainingDeck))
	}
}

func TestSaveDealtCardsToFileAndNewDealtCardsFromFile(t *testing.T) {
	os.Remove("_dealtCardTesting")

	dealtCards := deck{"Ace of Spades", "Two of Hearts", "Three of Diamonds"}
	dealtCards.saveDealtCardsToFile("_dealtCardTesting")

	loadedDealtCards := newDealtCardsFromFile("_dealtCardTesting")

	if len(loadedDealtCards) != 3 {
		t.Errorf("Expected 3 dealt cards, got %v", len(loadedDealtCards))
	}

	if loadedDealtCards[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", loadedDealtCards[0])
	}

	os.Remove("_dealtCardTesting")
}
