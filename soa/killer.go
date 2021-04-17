package soa

import (
	"os"
	"os/signal"
)

type Killer struct {
	done <-chan struct{}
}

func NewKiller() *Killer {
	done := make(chan struct{})

	go func() {
		stop := make(chan os.Signal)
		signal.Notify(stop, os.Interrupt, os.Kill)
		<-stop
		close(done)
	}()

	return &Killer{
		done: done,
	}
}

func (k *Killer) Done() <-chan struct{} {
	return k.done
}
