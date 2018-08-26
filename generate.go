package main

type PressureTestResult struct {
	id             int
	request_number int
	failure_number int
	duration_time  float32
}

func (r PressureTestResult) success_rate() int {
	return (r.request_number - r.failure_number) / r.request_number * 100
}

func (r PressureTestResult) average_time() float32 {
	return float32(r.request_number) / r.duration_time
}
