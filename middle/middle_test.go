package middle_test

import (
	"github.com/alexgunkel/golang_soa/middle"
	"github.com/alexgunkel/golang_soa/soa"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewMiddle(t *testing.T) {
	in := make(chan soa.Message)
	m := middle.NewMiddle("middle", in)
	assert.NotNil(t, m)

	in <- "test"

	select {
	case out, running := <-m.Messages():
		assert.True(t, running)
		assert.Contains(t, out, "test")
		assert.Contains(t, out, "middle")
	case <-time.After(time.Millisecond):
		assert.Fail(t, "timeout")
	}

	close(in)

	select {
	case out, running := <-m.Messages():
		assert.Empty(t, out)
		assert.False(t, running)
	}
}
