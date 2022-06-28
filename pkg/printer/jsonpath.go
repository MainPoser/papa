package printer

import (
	"bytes"
	"encoding/json"
	"io"

	"k8s.io/client-go/util/jsonpath"
)

type Jsonpath struct {
	Jp     *jsonpath.JSONPath
	Indent string
}

// NewJsonpathPrinter 创建jsonpathPrinter
func NewJsonpathPrinter(path string, indent string, allowMissingKeys bool) (*Jsonpath, error) {
	jp := jsonpath.New("default")
	jp.EnableJSONOutput(true)
	if err := jp.Parse(path); err != nil {
		return nil, err
	}
	jp.AllowMissingKeys(allowMissingKeys)
	return &Jsonpath{
		Jp:     jp,
		Indent: indent,
	}, nil

}

func (j *Jsonpath) Print(writer io.Writer, data []byte) error {
	var writeBytes []byte

	// 校验是否为合法json
	if valid := json.Valid(data); !valid {
		// 不是合法json，直接输出
		writeBytes = data
	} else {
		// 根据jsonpath获取数据
		var tmpObj = new(interface{})
		if err := json.Unmarshal(data, tmpObj); err != nil {
			return err
		}
		buffer := new(bytes.Buffer)
		if err := j.Jp.Execute(buffer, tmpObj); err != nil {
			return err
		}
		writeBytes = buffer.Bytes()

		buffer.Reset()
		if j.Indent != "" {
			if err := json.Indent(buffer, writeBytes, "", j.Indent); err != nil {
				return err
			}
		} else {
			if err := json.Compact(buffer, writeBytes); err != nil {
				return err
			}
		}
		writeBytes = buffer.Bytes()

		// 输出之前去掉两端的中括号
		writeBytes = writeBytes[1 : len(writeBytes)-1]
	}
	if _, err := writer.Write(writeBytes); err != nil {
		return err
	}
	return nil
}
