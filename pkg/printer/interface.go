package printer

import "io"

type Printer interface {
	//Print 打印，data需要传入合法的json字符串，如果不是，则直接输出，不格式化
	Print(writer io.Writer, data []byte) error
}
