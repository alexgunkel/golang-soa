package start

import "time"

type Start struct {
	close chan struct{}
	msg   chan string
}

func NewStart(name string, interval time.Duration) *Start {
	start := &Start{
		close: make(chan struct{}),
		msg:   make(chan string),
	}

	go func() {
		ticker := time.NewTicker(interval)

		for true {
			select {
			case <-ticker.C:
				start.msg <- name
			case <-start.close:
				ticker.Stop()
				close(start.msg)
			}
		}
	}()

	return start
}

func (s *Start) Stop() {
	select {
	case s.close <- struct{}{}:
	default:
	}
}

func (s *Start) Messages() <-chan string {
	return s.msg
}
