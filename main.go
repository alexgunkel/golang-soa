package main

import (
	"github.com/alexgunkel/golang_soa/end"
	"github.com/alexgunkel/golang_soa/middle"
	"github.com/alexgunkel/golang_soa/start"
	"time"
)

func main() {
	startOne := start.NewStart("start micro one", time.Millisecond*10)
	startTwo := start.NewStart("start micro one ", time.Millisecond*15)

	middleMicroOne := middle.NewMiddle("middle microOne", startOne.Messages())
	middleMicroTwo := middle.NewMiddle("middle microTwo", startOne.Messages())
	middleMicroThree := middle.NewMiddle("middle microTwo", startTwo.Messages())

	e1 := end.NewEnd("end one", middleMicroOne.Messages())
	e2 := end.NewEnd("end two", middleMicroOne.Messages())
	e3 := end.NewEnd("end three", middleMicroTwo.Messages())
	e4 := end.NewEnd("end four", middleMicroTwo.Messages())
	e5 := end.NewEnd("end four", middleMicroThree.Messages())
	e6 := end.NewEnd("end four", middleMicroThree.Messages())

	<-e1.Done()
	<-e2.Done()
	<-e3.Done()
	<-e4.Done()
	<-e5.Done()
	<-e6.Done()

	println("all done")
}
