package main

import (
	"time"
	"os"
)

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

type Formater interface {
	out([]*PressureTestResult, os.File)
}

type XmlFormater struct {
}

type CsvFormater struct {
}

type HtmlFormater struct {
}

type JsonFormater struct {
}

func (f XmlFormater) out([]*PressureTestResult, os.File) {

}

func (f CsvFormater) out([]*PressureTestResult, os.File) {

}

func (f HtmlFormater) out([]*PressureTestResult, os.File) {

}

func (f JsonFormater) out([]*PressureTestResult, os.File) {

}
