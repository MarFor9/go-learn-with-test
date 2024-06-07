package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWorld     = "Go!"
)

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{
		duration: 1 * time.Second,
		sleep:    time.Sleep,
	})
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(writer, finalWorld)
}

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(duration time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}
