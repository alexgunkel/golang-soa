package soa

import (
	"os"
	"os/signal"
)

type Killer interface {
	Done() <-chan struct{}
}

type sigKiller struct {
	done <-chan struct{}
}

func NewKiller(cb func(), signal2 ...os.Signal) Killer {
	done := make(chan struct{})

	go func() {
		stop := make(chan os.Signal)
		signal.Notify(stop, signal2...)
		<-stop
		close(done)
		cb()
	}()

	return &sigKiller{
		done: done,
	}
}

func (k *sigKiller) Done() <-chan struct{} {
	return k.done
}
