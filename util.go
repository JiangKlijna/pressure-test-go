package main

import (
	"os"
	"bytes"
	"fmt"
)

// File if exist
func FileExists(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return !stat.IsDir()
	}
	return false
}

// Dir if exist
func DirExists(path string) bool {
	if stat, err := os.Stat(path); err == nil {
		return stat.IsDir()
	}
	return false
}

// params to string
func params_string(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}
	buf := &bytes.Buffer{}
	for k, v := range params {
		buf.WriteString(k)
		buf.WriteByte('=')
		switch v.(type) {
		case string:
			buf.WriteString(v.(string))
		default:
			buf.WriteString(fmt.Sprint(v))
		}
		buf.WriteByte('&')
	}
	buf.Truncate(buf.Len() - 1)
	return buf.String()
}
