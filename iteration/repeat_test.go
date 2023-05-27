package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 8)
	expected := "aaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 8)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 3))
	// Output: aaa
}

func TestToLower(t *testing.T) {
	lower := strings.ToLower(("Gopher"))
	expected := "gopher"

	if lower != expected {
		t.Errorf("expected %q but got %q", expected, lower)
	}
}
