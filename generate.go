package main

import "time"

type PressureTestResult struct {
	id             int
	request_number int
	failure_number int
	duration_time  time.Duration
}

// get Success rate
func (r PressureTestResult) success_rate() float32 {
	return float32(r.request_number - r.failure_number) / float32(r.request_number) * 100
}

// get Average time consuming
func (r PressureTestResult) average_time() time.Duration {
	return r.duration_time / time.Duration(r.request_number)
}

// add data
func (r PressureTestResult) add(other *PressureTestResult) {
	r.request_number += other.request_number
	r.failure_number += other.failure_number
	r.duration_time += other.duration_time
}