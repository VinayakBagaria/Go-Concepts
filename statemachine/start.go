package statemachine

import (
	"errors"
	"fmt"
)

type State string

type Event string

type Node struct {
	State
	Transitions map[Event]*Transition
}

type Transition struct {
	*Node
	Action func() error
}

type StateMachine struct {
	InitialNode *Node
	CurrentNode *Node
}

func (m *StateMachine) GetCurrentNode() *Node {
	return m.CurrentNode
}

func (m *StateMachine) Transition(event Event) (*Node, error) {
	transition, ok := m.CurrentNode.Transitions[event]
	if !ok {
		return nil, errors.New(fmt.Sprintf("%s not found in %s", event, m.CurrentNode.State))
	}

	err := transition.Action()
	if err != nil {
		return nil, err
	}

	m.CurrentNode = transition.Node
	return m.CurrentNode, nil
}

func DoWork() {
	var (
		cartPageNode,
		checkoutProcessingNode,
		paymentProcessingNode,
		doneNode,
		failedNode Node
	)

	cartPageNode = Node{
		State: "cartPage",
		Transitions: map[Event]*Transition{
			"checkout_requested": &Transition{
				Node: &checkoutProcessingNode,
				Action: func() error {
					fmt.Println("cart page -> checkout_requested -> checkout processing")
					return nil
				}},
		},
	}

	checkoutProcessingNode = Node{
		State: "checkoutProcessing",
		Transitions: map[Event]*Transition{
			"payment_requested": &Transition{
				Node: &paymentProcessingNode,
				Action: func() error {
					fmt.Println("checkout processing -> payment_requested -> payment processing")
					return nil
				}},
		},
	}

	paymentProcessingNode = Node{
		State: "paymentProcessing",
		Transitions: map[Event]*Transition{
			"success": &Transition{
				Node: &doneNode,
				Action: func() error {
					fmt.Println("payment processing -> success -> done")
					return nil
				}},
			"timed_out": &Transition{
				Node: &cartPageNode,
				Action: func() error {
					fmt.Println("payment processing -> timed_out -> cart page")
					return nil
				}},
			"failed": &Transition{
				Node: &failedNode,
				Action: func() error {
					fmt.Println("payment processing -> failed -> failed")
					return nil
				},
			},
		},
	}

	machine := &StateMachine{InitialNode: &cartPageNode, CurrentNode: &cartPageNode}
	fmt.Printf("0. initial: %#v\n\n", machine.GetCurrentNode().State)

	nextNode, _ := machine.Transition("checkout_requested")
	fmt.Printf("1. next state for event checkout requested: %#v\n\n", nextNode.State)

	nextNode, err := machine.Transition("gibberish")
	fmt.Printf("2. next state for event gibberish: %#v, error: %#v\n\n", nextNode, err)

	nextNode, _ = machine.Transition("payment_requested")
	fmt.Printf("3. next state for event payment requested: %#v\n\n", nextNode.State)

	nextNode, _ = machine.Transition("timed_out")
	fmt.Printf("4. next state for event timed out: %#v\n\n", nextNode.State)

	nextNode, err = machine.Transition("success")
	fmt.Printf("5. next state for event success: %#v, error: %#v\n\n", nextNode, err)

	machine.Transition("checkout_requested")
	machine.Transition("payment_requested")
	nextNode, _ = machine.Transition("success")
	fmt.Printf("6. next state for event new success: %#v\n\n", nextNode)

	fmt.Printf("Initial node: %#v,\nCurrent node: %#v\n\n", machine.InitialNode.State, machine.CurrentNode.State)
}
