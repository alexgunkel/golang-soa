package middle_test

import (
	"github.com/alexgunkel/golang_soa/middle"
	"github.com/alexgunkel/golang_soa/soa"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollector_Messages(t *testing.T) {
	ch1 := make(chan soa.Message, 10)
	ch2 := make(chan soa.Message, 10)
	coll := middle.NewCollector(ch1, ch2)

	ch1 <- "a"
	ch2 <- "b"
	close(ch2)
	ch1 <- "c"
	close(ch1)

	for i := 0; i < 3; i++ {
		res, open := <-coll.Messages()
		assert.True(t, open)
		assert.Contains(t, "abc", res)
	}

	res, open := <-coll.Messages()
	assert.False(t, open)
	assert.Empty(t, res)
}
