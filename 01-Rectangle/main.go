/*
Create a rectangle abstraction using struct. Create two functions to calculate area
and circumference of given rectangle instance and set produced values on given
rectangle instance. Create another function to create an instance of rectangle
struct and return it.
But that function would be able to return an error in case of
passing invalid arguments such as negative length or height.
Choose one of strategies mentioned i the text book and exemplified in error.go of ch 05 in example
project.
*/
package main

import "fmt"

type rectangle struct {
	a, b int
}

func (r rectangle) display() {
	fmt.Println("Side lengths:", r.a, "and", r.b, "of the rectangle")
}

func (r rectangle) area() int {
	return r.a * r.b

}

func (r rectangle) circumference() int {
	return 2 * (r.a + r.b)

}

func main() {
	r1 := rectangle{5, 8}
	r1.display()
	fmt.Println("Area:", r1.area())
	fmt.Println("Circumference:", r1.circumference())
}
