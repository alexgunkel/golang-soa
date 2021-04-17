package start

import (
	"github.com/alexgunkel/golang_soa/soa"
	"strconv"
	"time"
)

type Start struct {
	msg chan soa.Message
}

func NewStart(name string, interval time.Duration) soa.Node {
	start := &Start{
		msg: make(chan soa.Message),
	}

	go func() {
		killer := soa.NewKiller()
		ticker := time.NewTicker(interval)
		cnt := uint64(0)

		for true {
			cnt++
			select {
			case <-ticker.C:
				start.msg <- soa.Message(name + " " + strconv.FormatUint(cnt, 10))
			case <-killer.Done():
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
