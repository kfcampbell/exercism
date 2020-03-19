// Package gigasecond should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package gigasecond

import "time"

// AddGigasecond should have a comment documenting it.
// somehow this is actually slower than the string version???
/*
=== RUN   TestAddGigasecond
    TestAddGigasecond: gigasecond_test.go:28: PASS: date only specification of time
    TestAddGigasecond: gigasecond_test.go:28: PASS: second test for date only specification of time
    TestAddGigasecond: gigasecond_test.go:28: PASS: third test for date only specification of time
    TestAddGigasecond: gigasecond_test.go:28: PASS: full time specified
    TestAddGigasecond: gigasecond_test.go:28: PASS: full time with day roll-over
    TestAddGigasecond: gigasecond_test.go:30: Tested 5 cases.
--- PASS: TestAddGigasecond (0.00s)
goos: darwin
goarch: amd64
pkg: gigasecond
BenchmarkAddGigasecond
BenchmarkAddGigasecond-12    	204734395	         5.69 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	gigasecond	1.781s
*/
func AddGigasecond(t time.Time) time.Time {
	const duration = time.Duration(1000000000 * 1000000000)
	return t.Add(duration)
}

// StringAddGigasecond holds a first attempt
// result of `go test -v --bench . --benchmem`
/*
--- PASS: TestAddGigasecond (0.00s)
goos: darwin
goarch: amd64
pkg: gigasecond
BenchmarkAddGigasecond
BenchmarkAddGigasecond-12    	20259914	        50.7 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	gigasecond	1.110s
*/
func StringAddGigasecond(t time.Time) time.Time {
	duration, err := time.ParseDuration("1000000000s")
	if err != nil {
		panic(err)
	}

	t = t.Add(duration)
	return t
}
