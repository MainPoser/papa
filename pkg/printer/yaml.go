package printer

import (
	"encoding/json"
	"io"

	"github.com/ghodss/yaml"
)

type Yaml struct {
}

func (y *Yaml) Print(writer io.Writer, data []byte) error {
	var writeBytes []byte

	// 校验是否为合法json
	if valid := json.Valid(data); !valid {
		// 不是合法json，直接输出
		writeBytes = data
	} else {
		// 输出yaml
		toYAML, err := yaml.JSONToYAML(data)
		if err != nil {
			return err
		}
		writeBytes = toYAML
	}
	if _, err := writer.Write(writeBytes); err != nil {
		return err
	}
	return nil
}

// NewYamlPrinter 创建yamlPrinter
func NewYamlPrinter() (*Yaml, error) {
	return &Yaml{}, nil
}
