package main

import (
	"reflect"
	"testing"
)

func TestSumArraySlices(t *testing.T) {
	t.Run("Array with 5 elements", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := SumArraySlices(numbers)
		want := 15

		assertIntegers(got, want, t)
	})
	t.Run("Array with any elements size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := SumArraySlices(numbers)
		want := 6

		assertIntegers(got, want, t)
	})
	t.Run("Multiple arrays ", func(t *testing.T) {
		firstArray := []int{1, 2}
		secondArray := []int{0, 9}
		got := SumAllArrays(firstArray, secondArray)
		want := []int{3, 9}

		assertIntegersArray(got, want, t)
	})

	t.Run("Multiple arrays tails ", func(t *testing.T) {
		firstArray := []int{1, 2}
		secondArray := []int{0, 9}
		got := SumAllTails(firstArray, secondArray)
		want := []int{2, 9}

		assertIntegersArray(got, want, t)
	})

	t.Run("Empty array  tails", func(t *testing.T) {
		firstArray := []int{}
		secondArray := []int{0, 9}
		got := SumAllTails(firstArray, secondArray)
		want := []int{0, 9}

		assertIntegersArray(got, want, t)
	})
}

func assertIntegers(got int, want int, t testing.TB) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
func assertIntegersArray(got []int, want []int, t testing.TB) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
