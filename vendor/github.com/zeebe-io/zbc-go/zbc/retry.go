package zbc

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

const maxInt64 = float64(math.MaxInt64 - 512)

// RetryDeadlineReached is error which occurs when request failed multiple times unsuccessfully.
var RetryDeadlineReached = errors.New("Message retry deadline reached")

// backoff is used to calculated backoffs for requests.
type backoff struct {
	attempt  float64
	Factor   float64
	Jitter   bool
	Min, Max time.Duration
}

// Duration returns the duration for the current attempt before incrementing.
func (b *backoff) Duration() time.Duration {
	d := b.ForAttempt(b.attempt)
	b.attempt++
	return d
}

// ForAttempt returns the duration for a specific attempt.
func (b *backoff) ForAttempt(attempt float64) time.Duration {
	min := b.Min
	if min <= 0 {
		min = 100 * time.Millisecond
	}
	max := b.Max
	if max <= 0 {
		max = 10 * time.Second
	}
	if min >= max {
		return max
	}
	factor := b.Factor
	if factor <= 0 {
		factor = 2
	}

	minf := float64(min)
	durf := minf * math.Pow(factor, attempt)
	if b.Jitter {
		durf = rand.Float64()*(durf-minf) + minf
	}
	if durf > maxInt64 {
		return max
	}
	dur := time.Duration(durf)
	if dur < min {
		return min
	} else if dur > max {
		return max
	}
	return dur
}

// Reset restarts the current attempt counter at zero.
func (b *backoff) Reset() {
	b.attempt = 0
}

// Attempt returns the current attempt counter value.
func (b *backoff) Attempt() float64 {
	return b.attempt
}

// Operation defines retry compatible operation.
type Operation func() (*Message, error)

// MessageRetry will try to execute operation and handle retrying under specified deadline.
func MessageRetry(op Operation) (*Message, error) {
	b := &backoff{
		Min:    BackoffMin,
		Max:    BackoffMax,
		Factor: 2,
		Jitter: true,
	}

	start := time.Now()
	for {
		msg, err := op()
		if err != nil && time.Since(start) < BackoffDeadline {
			time.Sleep(b.Duration())
			continue
		}
		if time.Since(start) > BackoffDeadline {
			return nil, RetryDeadlineReached
		}

		b.Reset()
		return msg, nil
	}
}
