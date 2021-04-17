package end

type End struct {
	done <-chan struct{}
}

func NewEnd(name string, in <-chan string) *End {
	doneChanel := make(chan struct{})
	go func() {
		for true {
			newMsg, running := <-in
			if !running {
				close(doneChanel)
				return
			}

			println(name + " received " + newMsg)
		}
	}()
	return &End{
		done: doneChanel,
	}
}

func (e *End) Done() <-chan struct{} {
	return e.done
}
