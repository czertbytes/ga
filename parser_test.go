package ga

import "testing"

func Test_ParseTimestampUserId(t *testing.T) {
	type Test struct {
		Timestamp int64
		UserId    int
		Value     string
		Err       error
	}

	tests := []Test{
		Test{
			Value: "",
			Err:   ErrParserContentWrongFormat,
		},
		Test{
			Value: "foobar",
			Err:   ErrParserContentWrongFormat,
		},
		Test{
			Value: `127.0.0.1 - hostname [-1371731248] "GET /123/foo/bar HTTP/1.1" 200 2326`,
			Err:   ErrParserContentWrongFormat,
		},
		Test{
			Value: `127.0.0.1 - hostname [1371731248] "GET /0.123/foo/bar HTTP/1.1" 200 2326`,
			Err:   ErrParserContentWrongFormat,
		},
		Test{
			Value: `127.0.0.1 - hostname HTTP/1.1" 200 2326`,
			Err:   ErrParserContentWrongFormat,
		},
		Test{
			Value: `127.0.0.1 - hostname [xxx] "GET /123/foo/bar HTTP/1.1" 200 2326`,
			Err:   ErrParserContentWrongFormat,
		},
		Test{
			Value: `127.0.0.1 - hostname [1371731248] "GET /xxx/foo/bar HTTP/1.1" 200 2326`,
			Err:   ErrParserContentWrongFormat,
		},
		Test{
			Value: `127.0.0.1 - hostname [1371731248] "GET / HTTP/1.1" 200 2326`,
			Err:   ErrParserContentWrongFormat,
		},
		Test{
			Timestamp: 1371731248,
			UserId:    123,
			Value:     `127.0.0.1 - hostname [1371731248] "GET /123/foo/bar HTTP/1.1" 200 2326`,
		},
	}

	for _, test := range tests {
		timestamp, userId, err := NewParser(test.Value).ParseTimestampUserId()

		if test.Err != nil {
			if err != test.Err {
				t.Fatalf("Bad result! Expected error %s but got %s", test.Err, err)
			}
		}

		if timestamp != test.Timestamp {
			t.Fatalf("Bad result! Expected timestamp %s but got %s", test.Timestamp, timestamp)
		}

		if userId != test.UserId {
			t.Fatalf("Bad result! Expected userId %s but got %s", test.UserId, userId)
		}
	}
}
