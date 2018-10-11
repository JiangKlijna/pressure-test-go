package main

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"time"
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

type Formater func([]*PressureTestResult, os.File)

func XmlFormater(res []*PressureTestResult, f os.File) {
	defer f.Close()
	data, _ := xml.Marshal(res)
	f.Write(data)
}

func CsvFormater(res []*PressureTestResult, f os.File) {
	defer f.Close()
	data, _ := csvMarshal(res)
	f.Write(data)
}

func HtmlFormater(res []*PressureTestResult, f os.File) {
	defer f.Close()
	data, _ := htmlMarshal(res)
	f.Write(data)
}

func JsonFormater(res []*PressureTestResult, f os.File) {
	defer f.Close()
	data, _ := json.Marshal(res)
	f.Write(data)
}

func csvMarshal(res []*PressureTestResult) ([]byte, error) {
	return nil, nil
}

func htmlMarshal(res []*PressureTestResult) ([]byte, error) {
	return nil, nil
}
