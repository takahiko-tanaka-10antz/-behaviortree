package behaviortree

import (
	"errors"
)

const (
	Success Status = 2
	Failure Status = 3
)

type (
	Node   func() (Tick, []Node)
	Tick   func(children []Node) (Status, error)
	Status int
)

func NewBranch(tick Tick, children ...Node) Node {
	return NewNode(tick, children)
}

func NewNode(tick Tick, children []Node) Node {
	return func() (Tick, []Node) {
		return tick, children
	}
}

func (n Node) Tick() (Status, error) {
	if n == nil {
		return Failure, errors.New("behaviortree.Tick :: cannot tick a nil node")
	}
	tick, children := n()
	if tick == nil {
		return Failure, errors.New("behaviortree.Tick :: cannot tick a node with a nil tick")
	}
	return tick(children)
}
