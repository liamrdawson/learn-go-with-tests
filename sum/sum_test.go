package main

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("collection of any size", func(t *testing.T) {

		numbers := []int{1, 2, 3}

		received := Sum(numbers)
		expected := 6

		if received != expected {
			t.Errorf("\nExpected: %d\nReceived: %d\nGiven: %v", expected, received, numbers)
		}

	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{2, 4, 6})
	want := []int{6, 12}

	if !slices.Equal(got, want) {
		t.Errorf("\nExpected: %v\nReceived: %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {

		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("\nExpected: %v\nReceived: %v", got, want)
		}

	}

	t.Run("Sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{3, 4, 5})
		want := []int{2, 9}

		checkSums(t, got, want)

	})

	t.Run("Safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2}, []int{3, 4, 5})
		want := []int{0, 2, 9}

		checkSums(t, got, want)

	})

}
