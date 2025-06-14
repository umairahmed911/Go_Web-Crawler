package main

import (
	"context"
	"time"
)

type RateLimiter struct {
	ticker *time.Ticker
	ch     chan struct{}
}

func NewRateLimiter(rps int) *RateLimiter {
	interval := time.Second / time.Duration(rps)
	return &RateLimiter{
		ticker: time.NewTicker(interval),
		ch:     make(chan struct{}, 1),
	}
}

func (rl *RateLimiter) Wait(ctx context.Context) {
	select {
	case <-rl.ticker.C:
		return
	case <-ctx.Done():
		return
	}
}
