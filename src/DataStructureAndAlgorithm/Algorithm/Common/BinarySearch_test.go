package Common

import "testing"

func TestBinarySearch(t *testing.T) {
	array := []int{1, 5, 9, 15, 81, 123, 189, 333}

	type test struct {
		input int
		want  int
	}

	tests := []test{
		{
			input: 2,
			want:  -1,
		},
		{
			input: 123,
			want:  5,
		},
	}

	for _, tc := range tests {
		got := BinarySearch(array, tc.input, 0, len(array)-1)
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
