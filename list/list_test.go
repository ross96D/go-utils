package list_test

import (
	"testing"

	"github.com/ross96D/go-utils/list"
	"github.com/stretchr/testify/assert"
)

func TestListForEach(t *testing.T) {
	var l list.List[int]
	l.PushFront(1)
	l.PushFront(2)
	l.PushFront(3)
	l.PushFront(4)

	count := 1
	for num := range l.Each {
		assert.Equal(t, count, num)
		count++
	}
}
