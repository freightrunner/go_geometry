package main

import (
	"fmt"
	"geometry/rectangle"
)

func main() {
	var rectLength, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used
	 */
	fmt.Printf("Area of rectangle %.2f\n", rectangle.Area(rectLength, rectWidth))
	/*Diagonal function of rectangle package used
	 */
	fmt.Printf("Diagonal of the rectangle %.2f", rectangle.Diagonal(rectLength, rectWidth))
}
