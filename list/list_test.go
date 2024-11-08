package list_test

import (
	"math/rand"
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

type TestStruct struct {
	n int
}

func (t TestStruct) Compare(o *TestStruct) int {
	return t.n - o.n
}

func TestSortedList(t *testing.T) {
	t.Run("test ISortedList", func(t *testing.T) {
		t.Parallel()

		l := list.ISortedList[TestStruct]{}

		seed := rand.Int63()
		source := rand.NewSource(seed)
		r := rand.New(source)

		for j := 0; j < 50000; j++ {
			v := int(r.Int())
			l.Append(TestStruct{n: v})
		}

		before := l.Elem(0)
		for _, v := range l.Iter(1) {
			require.True(t, before.Compare(v) <= 0)
		}
		println("seed:", seed)
	})

	t.Run("append", func(t *testing.T) {
		t.Parallel()

		l := list.SortedList[int]{}

		seed := rand.Int63()
		source := rand.NewSource(seed)
		r := rand.New(source)

		for i := 0; i < 50000; i++ {
			v := int(r.Int())
			l.Append(v)
		}

		before := l.Elem(0)
		for _, v := range l.Iter(1) {
			require.LessOrEqual(t, before, v)
		}
		println("seed:", seed)
	})

	t.Run("search", func(t *testing.T) {
		t.Parallel()

		l := list.SortedList[int]{}
		l.Append(2)
		l.Append(1)
		l.Append(5)
		i, ok := l.Search(1)
		assert.True(t, ok)
		assert.Equal(t, 0, i)

		ll := list.ISortedList[TestStruct]{}
		ll.Append(TestStruct{n: -1})
		ll.Append(TestStruct{n: 1})
		ll.Append(TestStruct{n: 5})
		i, ok = ll.Search(&TestStruct{n: 1})
		assert.True(t, ok)
		assert.Equal(t, 1, i)
	})
}
