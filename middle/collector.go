package middle

import (
	"github.com/alexgunkel/golang_soa/soa"
	"sync"
)

type collector struct {
	out chan soa.Message
}

func NewCollector(nodes ...<-chan soa.Message) soa.Node {
	o := make(chan soa.Message)

	wg := sync.WaitGroup{}
	for _, n := range nodes {
		n := n
		wg.Add(1)
		go func() {
			defer wg.Done()
			for true {
				v, open := <-n
				if !open {
					return
				}

				o <- v
			}
		}()
	}

	go func() {
		wg.Wait()
		close(o)
	}()

	return &collector{
		out: o,
	}
}

func (c *collector) Messages() <-chan soa.Message {
	return c.out
}
