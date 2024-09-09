package piscine

import "fmt"

func DealAPackOfCards(deck []int) {
	fmt.Printf("Player 1: %d, %d, %d\n", deck[0], deck[1], deck[2])
	fmt.Printf("Player 2: %d, %d, %d\n", deck[3], deck[4], deck[5])
	fmt.Printf("Player 3: %d, %d, %d\n", deck[6], deck[7], deck[8])
	fmt.Printf("Player 4: %d, %d, %d\n", deck[9], deck[10], deck[11])
}
