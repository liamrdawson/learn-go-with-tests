package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) SetDurationSlept(duration time.Duration) {
	s.durationSlept = duration
}

type SpyCountDownOperations struct {
	Calls []string
}

func (s *SpyCountDownOperations) Sleep() {
	s.Calls = append(s.Calls, "sleep")
}

func (s *SpyCountDownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, "write")
	return
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.SetDurationSlept}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v seconds but slept for %v seconds", sleepTime, spyTime.durationSlept)
	}
}

func TestCountdown(t *testing.T) {

	sleep := "sleep"
	write := "write"

	t.Run("prints 3 to go", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		mockSleeper := &SpyCountDownOperations{}

		Countdown(buffer, mockSleeper)

		got := buffer.String()
		want := `3
2
1
Go!`

		if got != want {
			t.Errorf("got %q but wanter %q", got, want)
		}
	})

	t.Run("sleeps before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOperations{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})

}
