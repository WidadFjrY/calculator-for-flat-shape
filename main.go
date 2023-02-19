package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	ViewHome()
}

// HomePage

func ViewHome() {

	fmt.Println("\n+++++++++++++Calculator+++++++++++++")
	fmt.Println("1. Square")
	fmt.Println("2. Rectangle")
	fmt.Println("3. Circle")
	fmt.Println("4. Triangle")

	inputUser := SelectionUserHome()

	switch inputUser {
	case 1:
		ViewSquare()
	case 2:
		ViewRectangle()
	case 3:
		ViewCircle()
	case 4:
		ViewTriangle()
	}
}

func SelectionUserHome() int {
	fmt.Print("Your selection : ")
	input := getInputUser()

	if input >= 1 && input <= 4 {
		return input
	} else {
		fmt.Println("Enter the correct number!!")
		return SelectionUserHome()
	}
}

// Utility

func getInputUser() int {
	var inputUser int

	_, err := fmt.Scanln(&inputUser)
	if err != nil {
		fmt.Println("Please enter the correct number.")
		os.Exit(0)
	}

	return inputUser
}

func selectOptionsUser() int {
	fmt.Print("Your selection : ")
	input := getInputUser()

	if input >= 1 && input <= 2 {
		return input
	} else {
		fmt.Println("Enter the correct number!!")
		return selectOptionsUser()
	}
}

// Repository

type SideOfTheSquare struct {
	side int
}

type LengthAndWidthOfTheRectangle struct {
	length, width int
}

type DiameterOfTheCircle struct {
	diameter int
}

type SideOfTheTriangle struct {
	height, base int
}

// Service

type Square interface {
	AreaOfSquare() int
	PerimeterOfSquare() int
}

func (s SideOfTheSquare) AreaOfSquare() int {
	return int(math.Pow(float64(s.side), 2))
}

func (s SideOfTheSquare) PerimeterOfSquare() int {
	return s.side * 4
}

type Rectangle interface {
	AreaOfRectangle() int
	PerimeterOfRectangle() int
}

func (s LengthAndWidthOfTheRectangle) AreaOfRectangle() int {
	return s.length * s.width
}

func (s LengthAndWidthOfTheRectangle) PerimeterOfRectangle() int {
	return 2 * (s.length + s.width)
}

type Circle interface {
	AreaOfCircle() float64
	CircumferenceOfCircle() float64
	RadiusOfCircle() float64
}

func (s DiameterOfTheCircle) AreaOfCircle() float64 {
	const PI = 3.14
	radius := float64(s.diameter) / 2

	return PI * math.Pow(radius, 2)
}

func (s DiameterOfTheCircle) CircumferenceOfCircle() float64 {
	const PI = 3.14
	radius := float64(s.diameter) / 2

	return 2 * PI * radius
}

func (s DiameterOfTheCircle) RadiusOfCircle() float64 {
	return float64(s.diameter) / 2
}

type Triangle interface {
	AreaOfTriangle() int
}

func (s SideOfTheTriangle) AreaOfTriangle() int {
	return (s.base / 2) * s.height
}

func PerimeterOfTriangle(side ...int) int {
	sum := 0
	for _, num := range side {
		sum += num
	}

	return sum
}

// View

func ViewSquare() {
	var square Square

	fmt.Println("\n+++++++++++SQUARE++++++++++++")
	fmt.Print("Input side of square (Cm) : ")
	square = SideOfTheSquare{getInputUser()}

	fmt.Printf("\nArea Of Square : %d Cm^2", square.AreaOfSquare())
	fmt.Printf("\nPerimeter Of Square : %d Cm\n", square.PerimeterOfSquare())
}

func ViewRectangle() {
	var rectangle Rectangle

	fmt.Println("\n+++++++++++++Rectangle+++++++++++++")
	fmt.Println("Input length and width of rectangle (Cm) : ")

	fmt.Print("Length : ")
	length := getInputUser()
	fmt.Print("Width : ")
	width := getInputUser()

	rectangle = LengthAndWidthOfTheRectangle{
		length: length,
		width:  width,
	}

	fmt.Printf("\nRectangle Area : %d Cm^2", rectangle.AreaOfRectangle())
	fmt.Printf("\nRectangle Perimeter : %d Cm\n", rectangle.PerimeterOfRectangle())
}

func ViewCircle() {
	var circle Circle

	fmt.Println("\n+++++++++++++Circle+++++++++++++")
	fmt.Print("Input Diameter of Circle (Cm) : ")
	circle = DiameterOfTheCircle{diameter: getInputUser()}

	fmt.Printf("\nArea Of Circle : %.2f Cm^2", circle.AreaOfCircle())
	fmt.Printf("\nCircumference Of Circle : %.2f Cm", circle.CircumferenceOfCircle())
	fmt.Printf("\nRadius Of Circle : %.2f Cm\n", circle.RadiusOfCircle())

}

func ViewTriangle() {
	var triangle Triangle

	fmt.Println("\n+++++++++++++Triangle+++++++++++++")
	fmt.Println("1. Area Of Triangle")
	fmt.Println("2. Perimeter Of Triangle")
	input := selectOptionsUser()

	if input == 1 {
		fmt.Println("\nInput height and base of Triangle (Cm) : ")
		fmt.Print("Height : ")
		height := getInputUser()
		fmt.Print("Base : ")
		base := getInputUser()

		triangle = SideOfTheTriangle{
			height: height,
			base:   base,
		}

		fmt.Printf("\nArea Of Triangle : %d Cm^2\n", triangle.AreaOfTriangle())
	} else if input == 2 {
		var sideA, sideB, sideC int

		fmt.Println("\nInput three side of Triangle (Cm) : ")
		fmt.Print("Side A : ")
		sideA = getInputUser()
		fmt.Print("Side B : ")
		sideB = getInputUser()
		fmt.Print("Side C : ")
		sideC = getInputUser()

		perimeterOfTriangle := PerimeterOfTriangle(sideA, sideB, sideC)
		fmt.Printf("Perimeter Of Triangle : %d Cm\n", perimeterOfTriangle)
	}
}
