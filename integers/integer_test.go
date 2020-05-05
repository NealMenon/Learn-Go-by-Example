package integers

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestAdder(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	y := rand.Intn(100)
	sum := Add(x, y)
	expected := x + y

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	} else {
		fmt.Printf("x = %d and y = %d\n", x, y)
	}
}
