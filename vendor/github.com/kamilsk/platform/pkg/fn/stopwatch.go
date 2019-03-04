package fn

import "time"

// Stopwatch calculates the fn execution time.
//
//  var result interface{}
//
//  duration := Stopwatch(func() { result = do.Some("heavy") })
//
func Stopwatch(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Now().Sub(start)
}
