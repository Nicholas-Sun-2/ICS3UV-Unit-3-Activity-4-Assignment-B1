/**
 * @author Nicholas Sun
 * @version 1.0.0
 * @date 2025-11-21
 * @fileoverview This program prints a message based on the colour chosen for the sweater.
 */

// Input
const colour: string = prompt("Please choose a sweater colour from the available choices: Blue, Black, Red, White.") || "No colour entered!";

// Output
if (colour == "Black" || colour == "White") {
  console.log("You have enough sweaters in this colour.");
} else if (colour == "Red") {
  console.log("This colour will look good on you.");
  console.log("How about a pair of jeans to go with the sweater?");
} else if (colour == "Blue") {
  console.log("This colour doesn't go well with your eyes.");
} else {
  console.log("Your colour choice is invalid.");
}
