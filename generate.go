package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	return (float32(r.RequestNumber-r.FailureNumber) / float32(r.RequestNumber)) * 100
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

func OutputResult(res []*PressureTestResult, f string) {
	getFormater(f)(res, getFilename(f))
}

type Formater func([]*PressureTestResult, string)

func getFormater(formater string) Formater {
	switch formater {
	case "csv":
		return CsvFormater
	case "xml":
		return XmlFormater
	case "html":
		return HtmlFormater
	case "json":
		return JsonFormater
	default:
		return GolangFormater
	}
}

func XmlFormater(res []*PressureTestResult, filename string) {
	data, _ := xmlMarshal(res)
	ioutil.WriteFile(filename, data, 0666)
}

func CsvFormater(res []*PressureTestResult, filename string) {
	data, _ := csvMarshal(res)
	ioutil.WriteFile(filename, data, 0666)
}

func HtmlFormater(res []*PressureTestResult, filename string) {
	data, _ := htmlMarshal(res)
	ioutil.WriteFile(filename, data, 0666)
}

func JsonFormater(res []*PressureTestResult, filename string) {
	data, _ := jsonMarshal(res)
	ioutil.WriteFile(filename, data, 0666)
}

func GolangFormater(res []*PressureTestResult, filename string) {
	data := []byte(fmt.Sprint(res))
	ioutil.WriteFile(filename, data, 0666)
}

func xmlMarshal(res []*PressureTestResult) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString("<PressureTestResults>\n")
	for _, r := range res {
		buf.WriteString("\t<PressureTestResult>\n")
		buf.WriteString(fmt.Sprintf("\t\t<Id>%d</Id><RequestNumber>%d</RequestNumber><FailureNumber>%d</FailureNumber><SuccessRate>%.2f</SuccessRate><AverageTime>%s</AverageTime>\n", r.Id, r.RequestNumber, r.FailureNumber, r.success_rate(), r.average_time()))
		buf.WriteString("\t</PressureTestResult>\n")
	}
	buf.WriteString("</PressureTestResults>")
	return buf.Bytes(), nil
}

func csvMarshal(res []*PressureTestResult) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString("Id,RequestNumber,FailureNumber,SuccessRate,AverageTime\n")
	for _, r := range res {
		buf.WriteString(fmt.Sprintf("%d,%d,%d,%f,%s\n", r.Id, r.RequestNumber, r.FailureNumber, r.success_rate(), r.average_time()))
	}
	return buf.Bytes(), nil
}

func htmlMarshal(res []*PressureTestResult) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString("<html><head><meta chararset=\"utf-8\" /><style>*{margin:0;padding:0;font-family:consolas}tr:nth-of-type(odd){background:#e8edff}tr:nth-of-type(even){background:white}tr:nth-of-type(odd):hover{background:white}tr:nth-of-type(even):hover{background:#e8edff}table{width:100%;text-align:center;color:#669}tr:first-of-type,tr:last-of-type{font-size:18px;font-weight:900}td{padding:6px}</style></head><body><table><tr><td>任务</td><td>请求总数</td><td>失败总数</td><td>成功率</td><td>平均耗时</td></tr>")
	for _, r := range res {
		buf.WriteString(fmt.Sprintf("<tr><td>%d</td><td>%d</td><td>%d</td><td>%f</td><td>%s</td></tr>", r.Id, r.RequestNumber, r.FailureNumber, r.success_rate(), r.average_time()))
	}
	buf.WriteString("</table></body></html>")
	return buf.Bytes(), nil
}

func jsonMarshal(res []*PressureTestResult) ([]byte, error) {
	buf := bytes.Buffer{}
	buf.WriteString("[")
	for i, r := range res {
		buf.WriteString(fmt.Sprintf("{\"Id\":\"%d\",\"RequestNumber\":\"%d\",\"FailureNumber\":\"%d\",\"SuccessRate\":\"%.2f\",\"AverageTime\":\"%s\"}", r.Id, r.RequestNumber, r.FailureNumber, r.success_rate(), r.average_time()))
		if i < len(res) -1 {
			buf.WriteByte(',')
		}
	}
	buf.WriteString("]")
	return buf.Bytes(), nil
}
