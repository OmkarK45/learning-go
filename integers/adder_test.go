package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("adding two integers", func(t *testing.T) {
		sum := Add(1, 2)

		if sum != 3 {
			t.Errorf("got %d, want %d", sum, 3)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 2)
	fmt.Println(sum)
	// Output: 3
}
