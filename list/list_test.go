package list_test

import (
	"testing"

	"github.com/ross96D/go-utils/list"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListForEach(t *testing.T) {
	var l list.LinkedList[struct{ num int }]
	l.PushBack(struct{ num int }{num: 1})
	l.PushBack(struct{ num int }{num: 2})
	l.PushBack(struct{ num int }{num: 3})
	l.PushBack(struct{ num int }{num: 4})

	count := 1
	forFuncPassed := false
	for num := range l.Each {
		forFuncPassed = true
		assert.Equal(t, count, num.Value.num)
		count++
	}
	require.True(t, forFuncPassed)
}

func TestRemove(t *testing.T) {
	var l list.LinkedList[struct{ num int }]
	e := l.PushBack(struct{ num int }{num: 1})
	l.PushBack(struct{ num int }{num: 2})
	l.PushBack(struct{ num int }{num: 3})
	e2 := l.PushBack(struct{ num int }{num: 4})

	l.Remove(e)
	l.Remove(e2)

	count := 2
	forFuncPassed := false
	for num := range l.Each {
		forFuncPassed = true
		assert.Equal(t, count, num.Value.num)
		count++
	}
	require.True(t, forFuncPassed)
}

func TestRemoveInsideEach(t *testing.T) {
	var l list.LinkedList[struct{ num int }]
	l.PushBack(struct{ num int }{num: 1})

	require.Equal(t, 1, l.Len())

	for node := range l.Each {
		if node.Value.num == 1 {
			l.Remove(&node)
		}
	}

	require.Equal(t, 0, l.Len())
}
