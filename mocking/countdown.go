package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	finalWord      = "Go!\n"
	countdownStart = 3
)

// Sleeper sleeps
type Sleeper interface {
	Sleep()
}

// SpySleeper pretends to sleep
type SpySleeper struct {
	Calls int
}

// DefaultSleeper is real, not for test
type DefaultSleeper struct{}

// Sleep is real  sleep function
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// Sleep is mock sleep function
func (s *SpySleeper) Sleep() {
	s.Calls++
}

// Countdown gives op to out
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}

	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
