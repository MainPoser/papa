package util

import (
	"io/ioutil"
	"os"
)

// ReadFromFileOrStdin 从文件或者标准输出入中读取文件，文件定义为空，从标准输入读取
func ReadFromFileOrStdin(file string) ([]byte, error) {
	var result []byte
	if file != "" {

		readFile, err := ioutil.ReadFile(file)
		if err != nil {
			return nil, err
		}
		result = readFile
	} else {

		readAll, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return nil, err
		}
		result = readAll
		return nil, nil
	}
	return result, nil
}
