package end

import "github.com/alexgunkel/golang_soa/soa"

type End struct {
	done <-chan struct{}
}

func NewEnd(name string, in <-chan soa.Message) soa.EndNode {
	doneChanel := make(chan struct{})
	go func() {
		for true {
			newMsg, running := <-in
			if !running {
				println("close " + name)
				close(doneChanel)
				return
			}

			println(name + " received " + string(newMsg))
		}
	}()
	return &End{
		done: doneChanel,
	}
}

func (e *End) Done() <-chan struct{} {
	return e.done
}
