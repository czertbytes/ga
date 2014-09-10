package ga

import (
	"errors"
	"testing"
)

func Test_ParseInt64Param(t *testing.T) {
	type Test struct {
		Name  string
		Map   map[string][]string
		Value int64
		Err   error
	}

	tests := []Test{
		Test{
			Name: "test",
			Map: map[string][]string{
				"test": []string{"1234"},
			},
			Value: 1234,
		},
		Test{
			Name: "test",
			Map: map[string][]string{
				"test": []string{"1234", "2345"},
			},
			Err: errors.New("bad request"),
		},
		Test{
			Name: "test",
			Map:  map[string][]string{},
			Err:  errors.New("bad request"),
		},
		Test{
			Name: "test",
			Map: map[string][]string{
				"aaa": []string{"1234"},
			},
			Err: errors.New("bad request"),
		},
		Test{
			Name: "test",
			Map: map[string][]string{
				"test": []string{"XXX"},
			},
			Err: errors.New("bad request"),
		},
	}

	for _, test := range tests {
		value, err := parseInt64Param(test.Name, test.Map)

		if test.Err != nil {
			if err == nil {
				t.Fatalf("Bad result! Expected error!")
			}
		}

		if value != test.Value {
			t.Fatalf("Bad result! Expected value %s but got %s", test.Value, value)
		}
	}
}
