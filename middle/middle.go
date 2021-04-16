package middle

type Middle struct {
	msg <-chan string
}

func NewMiddle(name string, incoming <-chan string) *Middle {
	msg := make(chan string)
	m := &Middle{
		msg: msg,
	}

	go func() {
		for true {
			newIn, running := <- incoming
			if !running {
				close(msg)
				return
			}
			msg <- name + " received " + newIn
		}
	}()

	return m
}

func (m *Middle) Messages() <-chan string {
	return m.msg
}
