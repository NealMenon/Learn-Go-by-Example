package iteration

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("Test Repeat5 for a", func(t *testing.T) {
		got := Repeat5("a")
		want := "aaaaa"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Test RepeatN for x, 7", func(t *testing.T) {
		got := RepeatN("x", 7)
		want := "xxxxxxx"
		assertCorrectMessage(t, got, want)
	})

}

func BenchmarkRepeat5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat5("a")
	}
}

func BenchmarkRepeatN(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RepeatN("x", 10)
	}
}
