package integers

import (
	"fmt"
	"testing"
)

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestAdder(t *testing.T) {

	t.Run("2 + 2 = 4", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertSummed(t, expected, sum)

	})

	t.Run("2 + 8 = 10", func(t *testing.T) {
		sum := Add(2, 8)
		expected := 10

		assertSummed(t, expected, sum)
	})
}

func assertSummed(t testing.TB, expected, actual int) {

	t.Helper()

	if expected != actual {
		t.Errorf("expected '%d' but got '%d", expected, actual)
	}
}
