package main

import (
	"fmt"
	"geometry/rectangle"
	"log"
)

var rectLength, rectWidth float64 = 6, 7

func init() {
	println("main package initialized")
	if rectLength < 0 {
		log.Fatal("length is less than zero")
	}
	if rectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func main() {
	fmt.Println("Geometrical shape properties")
	/*Area function of rectangle package used
	 */
	fmt.Printf("Area of rectangle %.2f\n", rectangle.Area(rectLength, rectWidth))
	/*Diagonal function of rectangle package used
	 */
	fmt.Printf("Diagonal of the rectangle %.2f", rectangle.Diagonal(rectLength, rectWidth))
}
