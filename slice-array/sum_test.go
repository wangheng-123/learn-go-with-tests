package slice_array

import "testing"

func Sum(numbers [5]int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 15

	if want != got {
		t.Errorf("got %d want %d given,%v", got, want, numbers)
	}
}
