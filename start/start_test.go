package start_test

import (
	"github.com/alexgunkel/golang_soa/start"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStartTicker(t *testing.T) {
	startModule := start.NewStart("start", time.Millisecond)

	for index := 0; index < 10; index++ {
		v, open := <-startModule.Messages()
		assert.True(t, open)
		assert.Contains(t, v, "start")
	}

	startModule.Stop()

	v, open := <-startModule.Messages()
	assert.False(t, open)
	assert.Empty(t, v)
}
