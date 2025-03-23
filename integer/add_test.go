package integer

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// ExampleAdd is a test function that prints the sum of 10 and 20.
// We have to add the comment "Output: <expected output>" to make it work.
// This is a good way to document the function.
func ExampleAdd() {
	sum := Add(10, 20)
	fmt.Println(sum)
	// Output: 30
}
