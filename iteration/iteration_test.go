package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	got := Repeat("a", 10)
	expected := "aaaaaaaaaa"

	if got != expected {
		t.Errorf("expected: %q but got: %q", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 10)
	}
}

func BenchmarkStrings_Repeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 10)
	}
}

func ExampleRepeat() {
	result := Repeat("b", 5)
	fmt.Println(result)
	// Output: bbbbb
}
