package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
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
	buf := bytes.Buffer{}
	buf.WriteString("Id,RequestNumber,FailureNumber,DurationTime\n")
	for _, r := range res  {
		buf.WriteString(fmt.Sprintf("%d,%d,%d,%s\n", r.Id, r.RequestNumber, r.FailureNumber, r.DurationTime))
	}
	return buf.Bytes(), nil
}

func htmlMarshal(res []*PressureTestResult) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString("<html><head><title>压力测试结果报告</title><meta chararset=\"utf-8\" /><style>*{margin:0;padding:0;font-family:consolas}tr:nth-of-type(odd){background:#e8edff}tr:nth-of-type(even){background:white}tr:nth-of-type(odd):hover{background:white}tr:nth-of-type(even):hover{background:#e8edff}table{width:100%;text-align:center;color:#669}tr:first-of-type,tr:last-of-type{font-size:18px;font-weight:900}tr:hover{cursor:pointer}td{padding:8px}</style></head><body><table><tr><td>任务</td><td>请求总数</td><td>失败总数</td><td>成功率</td><td>平均耗时(s)</td></tr>")
	for _, r := range res  {
		buf.WriteString(fmt.Sprintf("<tr><td>%d</td><td>%d</td><td>%d</td><td>%s</td></tr>", r.Id, r.RequestNumber, r.FailureNumber, r.DurationTime))
	}
	buf.WriteString("</table></body></html>")
	return buf.Bytes(), nil
}
