package middle

import "github.com/alexgunkel/golang_soa/soa"

type Middle struct {
	msg <-chan soa.Message
}

func NewMiddle(name string, incoming <-chan soa.Message) soa.Node {
	msg := make(chan soa.Message)
	m := &Middle{
		msg: msg,
	}

	go func() {
		for true {
			newIn, running := <- incoming
			if !running {
				println("close " + name)
				close(msg)
				return
			}
			msg <- soa.Message(name + " received ") + newIn
		}
	}()

	return m
}

func (m *Middle) Messages() <-chan soa.Message {
	return m.msg
}
