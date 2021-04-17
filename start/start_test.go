package start_test

import (
	"github.com/alexgunkel/golang_soa/start"
	"testing"
	"time"
)

func TestStartTicker(t *testing.T) {
	startModule := start.NewStart("start", time.Millisecond)
	go func() {
		time.Sleep(5 * time.Millisecond)
		startModule.Stop()
	}()

	count := 0
	end := time.After(time.Second)
loop:
	for true {
		select {
		case msg, running := <-startModule.Messages():
			if running && msg == "" {
				t.Error("empty string received")
				break loop
			} else if running {
				count++
			} else {
				break loop
			}
		case <-end:
			t.Errorf("timeout")
			break loop
		}
	}

	if count == 0 {
		t.Errorf("count is zero")
	}

	select {
	case _, running := <-startModule.Messages():
		if running {
			t.Error("still running")
		}
	case <-time.After(time.Millisecond):
		t.Errorf("did not stop")
	}
}
