package queue

import "testing"

func TestEnqueue(t *testing.T) {
	tests := []struct {
		name     string
		elements []interface{}
	}{
		{
			"int",
			[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"string",
			[]interface{}{"foo", "bar", "foobar"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			q := New()

			for _, element := range test.elements {
				q.Enqueue(element)
			}

			for _, expected := range test.elements {
				if got := q.Dequeue(); got != expected {
					t.Logf("Got: %v, Expected: %v,  queue: %s", got, expected, q)
					t.Fail()
				}
			}
		})
	}
}
