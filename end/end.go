package end

type End struct {
	done <-chan struct{}
}

func NewEnd(name string, in <-chan string) *End {
	go func() {
		for true {
			newMsg, running := <-in
			if !running {
				return
			}

			println(name + " received " + newMsg)
		}
	}()
	return &End{}
}

func (e *End) Done() <-chan struct{} {
	return e.done
}
