/*
 	Behavioural design pattern
	Restore an object to its previous state; Versioning

	Examples: Redux, Calculator, Editor undo feature
*/

package designpatterns

import (
	"fmt"
)

type Order struct {
	id     string
	client string
	amount float64
	status bool
}

type Memento struct {
	order *Order
}

type History struct {
	history []*Memento
}

func NewHistory() *History {
	return &History{}
}

func (h *History) Save(m *Memento) {
	h.history = append(h.history, m)
}

func (h *History) Get(index int) *Memento {
	return h.history[index]
}

type OrderWrapper struct {
	order *Order
}

func NewOrderWrapper() *OrderWrapper {
	return &OrderWrapper{}
}

func (o *OrderWrapper) Get() *Order {
	return o.order
}

func (o *OrderWrapper) Set(order *Order) {
	o.order = order
}

func (o *OrderWrapper) CreateMemento() *Memento {
	return &Memento{order: o.order}
}

func (o *OrderWrapper) Restore(m *Memento) {
	o.order = m.order
}

func usingMemento() {
	history := NewHistory()
	orderWrapper := NewOrderWrapper()

	order := Order{
		id:     "ORD001",
		client: "a6dba009adfe",
		amount: 4.99,
		status: false,
	}

	orderWrapper.Set(&order)
	history.Save(orderWrapper.CreateMemento())

	changedOrder := Order{
		id:     "ORD001",
		client: "a6dba009adfe",
		amount: 6.99,
		status: true,
	}
	orderWrapper.Set(&changedOrder)
	history.Save(orderWrapper.CreateMemento())

	fmt.Println("Latest:")
	fmt.Printf("%#v\n", orderWrapper.Get())

	fmt.Println("Printing History 1:")
	orderWrapper.Restore(history.Get(0))
	fmt.Printf("%#v\n", orderWrapper.Get())
	fmt.Println("Printing History 2:")
	orderWrapper.Restore(history.Get(1))
	fmt.Printf("%#v\n", orderWrapper.Get())
}
