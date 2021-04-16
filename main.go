package golang_soa

import (
	"github.com/alexgunkel/golang-soa/end"
	"github.com/alexgunkel/golang-soa/middle"
	"github.com/alexgunkel/golang-soa/start"
	"time"
)

func main() {
	startOne := start.NewStart("micro", time.Microsecond)

	middleMicroOne := middle.NewMiddle("microOne", startOne.Messages())
	middleMicroTwo := middle.NewMiddle("microTwo", startOne.Messages())

	end.NewEnd("micro one", middleMicroOne.Messages())
	end.NewEnd("micro two", middleMicroOne.Messages())
	end.NewEnd("micro three", middleMicroTwo.Messages())
	end.NewEnd("micro four", middleMicroTwo.Messages())

	time.Sleep(time.Minute)
}
