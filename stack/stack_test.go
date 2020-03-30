package stack_test

import (
	"data-structures/stack"
	"testing"
)

func TestStack(t *testing.T) {
	tests := []struct {
		name string
		push []interface{}
		pop  []interface{}
	}{
		{
			"int",
			[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			[]interface{}{9, 8, 7, 6, 5, 4},
		},
		{
			"float",
			[]interface{}{0.0, 1.1, 2.2, 3.3, 4.4},
			[]interface{}{4.4, 3.3},
		},
		{
			"string",
			[]interface{}{"foo", "bar", "baz", "foobar"},
			[]interface{}{"foobar", "baz"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s := stack.New()
			for _, item := range test.push {
				s.Push(item)
			}

			for _, expected := range test.pop {
				got := s.Pop()
				if got != expected {
					t.Logf("Got: %v, Expected: %v, Stack: %s", got, expected, s)
					t.Fail()
				}
			}
		})
	}
}
