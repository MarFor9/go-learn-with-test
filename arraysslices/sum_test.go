package arraysslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		checkSum := func(t testing.TB, got, expected int) {
			if got != expected {
				t.Errorf("expected: '%d' but got: '%d'", expected, got)
			}
		}
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expected := 15
		checkSum(t, got, expected)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("sum two slices", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		expected := []int{6, 15}
		checkSums(t, got, expected)
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum tails of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 9})
		expected := []int{5, 9}
		checkSums(t, got, expected)
	})
	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		expected := []int{0, 9}
		checkSums(t, got, expected)
	})
}

func checkSums(t testing.TB, got, expected []int) {
	t.Helper()
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("expected: %v got %v", expected, got)
	}
}
