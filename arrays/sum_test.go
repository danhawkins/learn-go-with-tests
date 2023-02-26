package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func ExampleSum() {
	numbers := []int{1, 2, 3, 4, 5}
	result := Sum(numbers)

	fmt.Println(result)
	// Output: 15
}

func ExampleSumAll() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{2, 4, 6}

	result := SumAll(nums1, nums2)
	fmt.Println(result)

	//Output: [6 12]
}

func TestSum(t *testing.T) {

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		assertSum(t, got, want, numbers)
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertSum(t testing.TB, got int, want int, numbers []int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

// func assertSumAll(t testing.TB, got []int, want []int) {
// 	t.Helper()

// 	if !reflect.DeepEqual(got, want) {
// 		t.Errorf("got %v want %v", got, want)
// 	}
// }
