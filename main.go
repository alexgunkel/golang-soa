package main

import (
	"github.com/alexgunkel/golang_soa/end"
	"github.com/alexgunkel/golang_soa/middle"
	"github.com/alexgunkel/golang_soa/soa"
	"github.com/alexgunkel/golang_soa/start"
	"os"
	"time"
)

func main() {
	startOne := start.NewStart("start micro one", time.Millisecond*10)
	startTwo := start.NewStart("start micro one ", time.Millisecond*15)

	soa.NewKiller(func() {
		startOne.Stop()
		startTwo.Stop()
	}, os.Kill, os.Interrupt)

	middleMicroOne := middle.NewMiddle("middle microOne", startOne.Messages())
	middleMicroTwo := middle.NewMiddle(
		"middle microTwo",
		middle.NewCollector(startOne.Messages(), startTwo.Messages()).Messages(),
	)
	middleMicroThree := middle.NewMiddle("middle microTwo", startTwo.Messages())

	coll := middle.NewCollector(middleMicroOne.Messages(), middleMicroTwo.Messages(), middleMicroThree.Messages())

	<-end.NewEnd("end one", coll.Messages()).Done()

	println("all done")
}
