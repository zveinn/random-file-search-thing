package main

func main() {

	h := new(Honda)
	DoSomethingWithInterface(h)
}

func DoSomethingWithInterface(car Car) {
	car.Drive(3)
}

type Car interface {
	Drive(int) int
	// ChangeGear(int)
}

type Honda struct {
}

func (h *Honda) Drive(x int) int {
	return 1
}
