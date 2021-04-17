package soa

type Message string

type Node interface {
	Messages() <-chan Message
}

type Stoppable interface {
	Node
	Stop()
}

type End interface {
	Done() <-chan struct{}
}
