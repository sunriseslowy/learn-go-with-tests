package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func BenchmarkStringRepeat(b *testing.B) {

	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 5)
	}
}

func ExampleRepeat() {
	output := Repeat("a", 5)
	fmt.Println(output)
	// Output: aaaaa
}
