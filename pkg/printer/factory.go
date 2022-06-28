package printer

import (
	"errors"
	"strings"
)

const (
	PrefixJsonPath = "jsonpath="
	OutputJson     = "json"
	OutputYaml     = "yaml"
)

const (
	SplitCharEqual = "="
)

var (
	notSupportOutputErr = errors.New("output not suppotr")
)

// GetPrinter 获取printer
func GetPrinter(output, indent string, allowMissingKeys bool) (Printer, error) {
	if strings.HasPrefix(output, PrefixJsonPath) {
		split := strings.Split(output, SplitCharEqual)
		return NewJsonpathPrinter(split[1], indent, allowMissingKeys)
	} else if output == OutputYaml {
		return NewYamlPrinter()
	} else if output == OutputJson {
		return NewJsonPrinter(indent)
	} else {
		return nil, notSupportOutputErr
	}

}
