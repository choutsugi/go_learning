package Common

import "testing"

func TestFibonacci(t *testing.T) {

	c := make(chan int, 1)

	type test struct {
		input int
		want  int
	}

	tests := []test{
		{
			input: 5,
			want:  3,
		},
		{
			input: 10,
			want:  34,
		},
	}

	for _, tc := range tests {
		go Fibonacci(c, tc.input)
		got := <-c
		if got != tc.want {
			t.Errorf("expected: %v, got: %v", tc.want, got)
		}
	}
}
