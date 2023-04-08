package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// a type to represent Deck of cards
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "two", "three", "four"}

	// creation of deck of cards
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

// function to return the two seperate slices
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// returns a string from deck of cards
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

// save to hard-drive
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	// Error handling
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// processing the byte slice and returning the deck
	stringSlice := strings.Split(string(bs), ",")

	return deck(stringSlice)
}

// function to shuffle the deck of cards
func (d deck) shuffle() {

	source := rand.NewSource(time.Now().Unix())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
