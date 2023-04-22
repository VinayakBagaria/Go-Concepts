package designpatterns

import "fmt"

func DoWork() {
	fmt.Println("-----------------")
	fmt.Println("Mediator pattern")
	usingMediator()

	fmt.Println("-----------------")
	fmt.Println("Memento pattern")
	usingMemento()
}
