package Common

import "testing"

func TestFactrial(t *testing.T) {

	c := make(chan int64, 1)

	type test struct {
		input int
		want  int
	}

	tests := []test{
		{
			input: 5,
			want:  120,
		},
		{
			input: 10,
			want:  3628800,
		},
	}

	for _, tc := range tests {
		go Factrial(c, tc.input)
		got := <-c
		if got != int64(tc.want) {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
