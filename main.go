package main

import (
	"github.com/alexgunkel/golang_soa/end"
	"github.com/alexgunkel/golang_soa/middle"
	"github.com/alexgunkel/golang_soa/start"
	"os"
	"os/signal"
	"time"
)

func main() {
	startOne := start.NewStart("micro", time.Millisecond*100)

	middleMicroOne := middle.NewMiddle("microOne", startOne.Messages())
	middleMicroTwo := middle.NewMiddle("microTwo", startOne.Messages())

	e1 := end.NewEnd("micro one", middleMicroOne.Messages())
	e2 := end.NewEnd("micro two", middleMicroOne.Messages())
	e3 := end.NewEnd("micro three", middleMicroTwo.Messages())
	e4 := end.NewEnd("micro four", middleMicroTwo.Messages())

	time.Sleep(time.Millisecond)

	go func() {
		stop := make(chan os.Signal)
		signal.Notify(stop, os.Interrupt, os.Kill)
		<-stop
		startOne.Stop()
	}()

	<-e1.Done()
	<-e2.Done()
	<-e3.Done()
	<-e4.Done()
	<-e4.Done()

	println("all done")
}
