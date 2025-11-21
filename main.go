/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview This program prints a message based on the colour chosen for the sweater.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Declaring variables
	var colour string

	// Input
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please choose a sweater colour from the available choices: Blue, Black, Red, White. ")
	colour, _ = reader.ReadString('\n')
	colour = strings.TrimSpace(colour)

	// Output
	if colour == "Black" || colour == "White" {
		fmt.Println("You have enough sweaters in this colour.")
	} else if colour == "Red" {
		fmt.Println("This colour will look good on you.")
		fmt.Println("How about a pair of jeans to go with the sweater?")
	} else if colour == "Blue" {
		fmt.Println("This colour doesn't go well with your eyes.")
	} else {
		fmt.Println("Your colour choice is invalid.")
	}
}
