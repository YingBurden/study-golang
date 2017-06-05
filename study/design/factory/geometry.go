package factory

const PI = 3.14

type Geometry interface {
	GetArea() float64
	GetPer() float64
}

type Circle struct {
	Radius float64
}

func (circleGeometry *Circle) GetArea() float64 {
	return circleGeometry.Radius * PI * PI
}

func (circleGeometry *Circle) GetPer() float64 {
	return 2 * circleGeometry.Radius * PI
}

type Rectangle struct {
	Hight  float64
	Weight float64
}

func (rectangleGeometry *Rectangle) GetArea() float64 {
	return rectangleGeometry.Hight * rectangleGeometry.Weight
}

func (rectangleGeometry *Rectangle) GetPer() float64 {
	return 2 * (rectangleGeometry.Hight + rectangleGeometry.Weight)
}

type OperationFactory struct {
}

func (this *OperationFactory) CreateGeometry(functype string) (geo Geometry) {

	switch functype {
	case "circle":
		geo = new(Circle)
	case "rectangle":
		geo = new(Rectangle)

	}
	return
}
