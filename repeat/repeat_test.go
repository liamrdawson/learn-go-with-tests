package iteration

import "testing"

func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a", 6)
	}
}

func TestRepeat(t *testing.T) {

	t.Run("should repeat the provided number of times", func(t *testing.T) {

		repeated := Repeat("a", 6)
		expected := "aaaaaa"

		if expected != repeated {
			t.Errorf("\nExpected: %q\nReceived: %q", expected, repeated)
		}

	})
}
