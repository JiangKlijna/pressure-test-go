package main

import "time"

type PressureTestResult struct {
	Id            int
	RequestNumber int
	FailureNumber int
	DurationTime  time.Duration
}

// get Success rate
func (r *PressureTestResult) success_rate() float32 {
	return float32(r.RequestNumber-r.FailureNumber) / float32(r.RequestNumber) * 100
}

// get Average time consuming
func (r *PressureTestResult) average_time() time.Duration {
	return r.DurationTime / time.Duration(r.RequestNumber)
}

// add data
func (r *PressureTestResult) add(other *PressureTestResult) {
	r.RequestNumber += other.RequestNumber
	r.FailureNumber += other.FailureNumber
	r.DurationTime += other.DurationTime
}

// mark PressureTestResult
func (r *PressureTestResult) mark(isFailure bool, start time.Time) {
	r.RequestNumber++
	if isFailure {
		r.FailureNumber++
	}
	r.DurationTime += time.Since(start)
}
