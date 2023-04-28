package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)

		expected := 15

		if got != expected {
			t.Errorf("expected %d, got %d, for input %v", expected, got, numbers)
		}

	})

}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{2, 3}, []int{1, 2, 3})
	expected := []int{5, 6}

	checkSum(t, got, expected)

}

func TestSumAllTails(t *testing.T) {

	t.Run("sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{2, 9}, []int{5, 2})
		want := []int{9, 2}

		checkSum(t, got, want)
	})

	t.Run("safely compute sum of empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{5, 2, 6})
		want := []int{0, 8}
		checkSum(t, got, want)

	})

}

func checkSum(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, expected %v", got, want)
	}
}
