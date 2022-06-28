package printer

import (
	"bytes"
	"encoding/json"
	"io"
)

type Json struct {
	Indent string
}

func (j *Json) Print(writer io.Writer, data []byte) error {
	var writeBytes []byte
	// 校验是否为合法json
	if valid := json.Valid(data); !valid {
		// 不是合法json，直接输出
		writeBytes = data
	} else {
		// 输出json
		buf := new(bytes.Buffer)
		if j.Indent != "" {
			if err := json.Indent(buf, data, "", j.Indent); err != nil {
				return err
			}
		} else {
			if err := json.Compact(buf, data); err != nil {
				return err
			}
		}
		writeBytes = buf.Bytes()
	}
	if _, err := writer.Write(writeBytes); err != nil {
		return err
	}
	return nil
}

// NewJsonPrinter 创建jsonPrinter
func NewJsonPrinter(indent string) (*Json, error) {
	return &Json{Indent: indent}, nil
}
