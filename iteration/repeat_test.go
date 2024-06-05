package iteration

import (
	"fmt"
	"testing"
)

// Let's write a test for a function that repeats a character 5 times
func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected '%q' but got '%q'", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
		// Repeat("a")
	}
}

func ExampleRepeat() {
	repeat := Repeat("x", 5)
	fmt.Println(repeat)
}
