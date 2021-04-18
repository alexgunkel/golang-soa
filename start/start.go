package start

import (
	"github.com/alexgunkel/golang_soa/soa"
	"strconv"
	"time"
)

type Start struct {
	msg  chan soa.Message
	stop chan struct{}
}

func NewStart(name string, interval time.Duration) soa.Stoppable {
	start := &Start{
		msg:  make(chan soa.Message),
		stop: make(chan struct{}),
	}

	go func() {
		ticker := time.NewTicker(interval)
		cnt := uint64(0)

		for true {
			cnt++
			select {
			case <-ticker.C:
				start.msg <- soa.Message(name + " " + strconv.FormatUint(cnt, 10))
			case <-start.stop:
				println("close " + name)
				ticker.Stop()
				close(start.msg)
				return
			}
		}
	}()

	return start
}

func (s *Start) Messages() <-chan soa.Message {
	return s.msg
}

func (s *Start) Stop() {
	select {
	case s.stop <- struct{}{}:
	default:
	}
}
