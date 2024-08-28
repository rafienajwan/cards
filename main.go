package main

import (
	"fmt"
	"os"
)

func main() {
	var cards, dealtCards deck
	var choice int

	for {
		fmt.Println("===== Cards Shuffler =====")
		fmt.Println("1. Create new deck")
		fmt.Println("2. Shuffle the deck")
		fmt.Println("3. Save the deck")
		fmt.Println("4. Load the deck from file")
		fmt.Println("5. Deal the deck")
		fmt.Println("6. Print the deck")
		fmt.Println("7. Show dealt cards")
		fmt.Println("8. Save dealt cards")
		fmt.Println("9. Load dealt cards from file")
		fmt.Println("10. Exit")
		fmt.Println("==========================")
		fmt.Print("Choose an option : ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			cards = newDeck()
			dealtCards = nil
			fmt.Println("New deck created.")
		case 2:
			if len(cards) == 0 {
				fmt.Println("No deck available to shuffle. Please create a deck.")
			} else {
				cards.shuffle()
				fmt.Println("Deck shuffled.")
			}
		case 3:
			if len(cards) == 0 {
				fmt.Println("No deck available to save. Please create a deck.")
			} else {
				var filename string
				fmt.Print("Enter file name : ")
				fmt.Scan(&filename)
				cards.saveToFile(filename)
				fmt.Println("Deck saved as", filename)
			}
		case 4:
			var filename string
			fmt.Print("Enter file name to load : ")
			fmt.Scan(&filename)
			cards = newDeckFromFile(filename)
			dealtCards = nil
			fmt.Println("Deck loaded from ", filename)
		case 5:
			if len(cards) == 0 {
				fmt.Println("No deck available to deal. Please create a deck.")
			} else {
				var handSize int
				fmt.Print("Enter number of cards for hand : ")
				fmt.Scan(&handSize)
				hand, remainingDeck := deal(cards, handSize)
				fmt.Println("Hand dealt :")
				hand.print()
				cards = remainingDeck
				dealtCards = append(dealtCards, hand...)
			}
		case 6:
			if len(cards) == 0 {
				fmt.Println("No deck available to print. Please create a deck.")
			} else {
				fmt.Println("Current deck :")
				cards.print()
			}
		case 7:
			if len(dealtCards) == 0 {
				fmt.Println("No cards have been dealt.")
			} else {
				fmt.Println("Dealt cards :")
				dealtCards.print()
			}
		case 8:
			if len(dealtCards) == 0 {
				fmt.Println("No dealt cards to save.")
			} else {
				var filename string
				fmt.Print("Enter file name : ")
				fmt.Scan(&filename)
				dealtCards.saveDealtCardsToFile(filename)
				fmt.Println("Dealt cards saved as ", filename)
			}
		case 9:
			var filename string
			fmt.Print("Enter file name to load : ")
			fmt.Scan(&filename)
			dealtCards = newDealtCardsFromFile(filename)
			fmt.Println("Dealt cards loaded from ", filename)
		case 10:
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid option, please choose again.")
		}
		fmt.Println()
	}
}
