package singlylinkedlist

import "testing"

func TestAppend(t *testing.T) {
	tests := []struct {
		name     string
		elements []interface{}
		expected []interface{}
	}{
		{
			"int",
			[]interface{}{0, 1, 2, 3, 4},
			[]interface{}{0, 1, 2, 3, 4},
		},
		{
			"float",
			[]interface{}{-99.123, -98.345, -200.12654, 1241412.0, 55555.88, 0.0},
			[]interface{}{-99.123, -98.345, -200.12654, 1241412.0, 55555.88, 0.0},
		},
		{
			"string",
			[]interface{}{"foo", "bar", "baz"},
			[]interface{}{"foo", "bar", "baz"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := New(false)
			for _, element := range test.elements {
				l.Append(element)
			}

			i := 0
			for elem := l.Iter(); elem != nil; elem = elem.Next() {
				if elem.Data() != test.expected[i] {
					t.Logf("Expected: %v, Got: %v", test.expected[i], elem)
					t.FailNow()
				}
				i++
			}
		})
	}
}

func TestPrepend(t *testing.T) {
	tests := []struct {
		name     string
		elements []interface{}
		prepend  []interface{}
		expected []interface{}
	}{
		{
			"int",
			[]interface{}{5, 6, 7, 8, 9},
			[]interface{}{4, 3, 2, 1, 0},
			[]interface{}{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"float",
			[]interface{}{5.5, 6.6, 7.7, 8.8, 9.9},
			[]interface{}{4.4, 3.3, 2.2, 1.1, 0.0},
			[]interface{}{0.0, 1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7, 8.8, 9.9},
		},
		{
			"string",
			[]interface{}{"foo", "bar", "baz"},
			[]interface{}{"foobar"},
			[]interface{}{"foobar", "foo", "bar", "baz"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := New(false)
			for _, element := range test.elements {
				l.Append(element)
			}

			for _, element := range test.prepend {
				l.Prepend(element)
			}

			i := 0
			for elem := l.Iter(); elem != nil; elem = elem.Next() {
				if elem.Data() != test.expected[i] {
					t.Logf("Expected: %v, Got: %v", test.expected[i], elem)
					t.FailNow()
				}
				i++
			}
		})
	}
}

func TestInsertAt(t *testing.T) {
	tests := []struct {
		name     string
		elements []interface{}
		item     interface{}
		index    int
		expected []interface{}
	}{
		{
			"int-head",
			[]interface{}{1, 2, 3, 4},
			0,
			0,
			[]interface{}{0, 1, 2, 3, 4},
		},
		{
			"int-tail",
			[]interface{}{0, 1, 2, 3},
			4,
			4,
			[]interface{}{0, 1, 2, 3, 4},
		},
		{
			"int",
			[]interface{}{0, 1, 3, 4},
			2,
			2,
			[]interface{}{0, 1, 2, 3, 4},
		},
		{
			"float",
			[]interface{}{0.0, 1.1, 2.2, 4.4},
			3.3,
			3,
			[]interface{}{0.0, 1.1, 2.2, 3.3, 4.4},
		},
		{
			"string",
			[]interface{}{"foo", "bar", "baz"},
			"foobar",
			2,
			[]interface{}{"foo", "bar", "foobar", "baz"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := New(false)
			for _, element := range test.elements {
				l.Append(element)
			}

			if err := l.InsertAt(test.index, test.item); err != nil {
				t.Log(l)
				t.Errorf("%v", err)
			}

			i := 0
			for elem := l.Iter(); elem != nil; elem = elem.Next() {
				if elem.Data() != test.expected[i] {
					t.Logf("Expected: %v, Got: %v, list: %s", test.expected[i], elem, l)
					t.FailNow()
				}
				i++
			}
		})
	}
}

func TestDeleteAt(t *testing.T) {
	tests := []struct {
		name     string
		elements []interface{}
		index    int
		expected []interface{}
	}{
		{
			"int-head",
			[]interface{}{0, 1, 2, 3, 4},
			0,
			[]interface{}{1, 2, 3, 4},
		},
		{
			"int-tail",
			[]interface{}{0, 1, 2, 3, 4},
			4,
			[]interface{}{0, 1, 2, 3},
		},
		{
			"int",
			[]interface{}{0, 1, 2, 3, 4},
			2,
			[]interface{}{0, 1, 3, 4},
		},
		{
			"float",
			[]interface{}{0.0, 1.1, 2.2, 3.3, 4.4},
			3,
			[]interface{}{0.0, 1.1, 2.2, 4.4},
		},
		{
			"string",
			[]interface{}{"foo", "bar", "foobar", "baz"},
			2,
			[]interface{}{"foo", "bar", "baz"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := New(false)
			for _, element := range test.elements {
				l.Append(element)
			}

			if err := l.DeleteAt(test.index); err != nil {
				t.Log(l)
				t.Errorf("%v", err)
			}

			i := 0
			for elem := l.Iter(); elem != nil; elem = elem.Next() {
				if elem.Data() != test.expected[i] {
					t.Logf("Expected: %v, Got: %v, list: %s", test.expected[i], elem, l)
					t.FailNow()
				}
				i++
			}
		})
	}
}
