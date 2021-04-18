package soa

type Message string

type Node interface {
	Messages() <-chan Message
}

type StartNode interface {
	Node
	Stop()
}

type EndNode interface {
	Done() <-chan struct{}
}
