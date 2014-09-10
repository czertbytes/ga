package ga

import (
	"math/rand"
	"sort"
	"testing"
)

// Tests if after inserting new record storage is sorted by timestamp
func Test_Insert(t *testing.T) {
	s := NewStorage()

	for i := 0; i < 10000; i++ {
		s.Insert(NewRecord(rand.Int63(), 1))
	}

	// records should be sorted by timestamp
	if !sort.IsSorted(ByTimestamp(s.Records)) {
		t.Fatalf("Bad result! Records are not sorted! Records: %s", s.Records)
	}
}

// Tests if searching by timestamp frame window
func Test_Find(t *testing.T) {
	s := NewStorage()

	s.Insert(NewRecord(2, 1))
	s.Insert(NewRecord(5, 1))
	s.Insert(NewRecord(10, 2))
	s.Insert(NewRecord(2, 3))
	s.Insert(NewRecord(4, 4))
	s.Insert(NewRecord(1, 4))
	s.Insert(NewRecord(3, 5))

	type Test struct {
		Start      int64
		End        int64
		ResultsLen int
	}

	tests := []Test{
		Test{
			Start:      0,
			End:        0,
			ResultsLen: 0,
		},
		Test{
			Start:      1000,
			End:        2000,
			ResultsLen: 0,
		},
		Test{
			Start:      0,
			End:        3,
			ResultsLen: 3, // timestamps [1, 2, 2]
		},
		Test{
			Start:      3,
			End:        10,
			ResultsLen: 3, // timestamps [3, 4, 5]
		},
		Test{
			Start:      3,
			End:        11,
			ResultsLen: 4, // timestamps [3, 4, 5, 10]
		},
	}

	for _, test := range tests {
		records := s.Find(test.Start, test.End)

		if len(records) != test.ResultsLen {
			t.Fatalf("Bad result! Expected %d but got %d", test.ResultsLen, len(records))
		}
	}
}
