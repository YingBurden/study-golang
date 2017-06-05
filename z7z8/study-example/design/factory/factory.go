package main

import (
	"fmt"

	"github.com/YingBurden/study-golang/study/design/factory"
)

func main() {
	var testFactory factory.OperationFactory
	circle := testFactory.CreateGeometry("circle")
	fmt.Println("cicle area:", circle.GetArea())
	rectangle := testFactory.CreateGeometry("rectangle")
	fmt.Println("rectangle area:", rectangle.GetArea())
}
