package ga

import "testing"

// Tests result of unique userIds action
func Test_UniqUserIds(t *testing.T) {
	type Test struct {
		Records Records
		UniqIds int
	}

	tests := []Test{
		Test{
			Records: Records{
				NewRecord(1, 1),
				NewRecord(2, 1),
				NewRecord(3, 2),
			},
			UniqIds: 2,
		},
		Test{
			Records: Records{
				NewRecord(1, 1),
				NewRecord(2, 2),
				NewRecord(3, 3),
			},
			UniqIds: 3,
		},
		Test{
			Records: Records{},
			UniqIds: 0,
		},
	}

	for _, test := range tests {
		uniqIds := test.Records.UniqueUserIds()

		if uniqIds != test.UniqIds {
			t.Fatalf("Bad result! Expected %d but got %d", test.UniqIds, uniqIds)
		}
	}
}
