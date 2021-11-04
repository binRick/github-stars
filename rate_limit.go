package main

import (
	"time"

	"go.uber.org/ratelimit"
)

var github_request_rate_limiter = ratelimit.New(1)
var prev = time.Now()
