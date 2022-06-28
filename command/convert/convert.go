package convert

import (
	"encoding/json"
	"github.com/MainPoser/papa/pkg/printer"
	"github.com/MainPoser/papa/pkg/util"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
	"os"
	"strconv"
)

var (
	convertCmd = &cobra.Command{
		Use:   "convert",
		Short: "convert data",
		Long:  "convert, like <Anakin Skywalker> to <Darth Vader>",
		Example: `read data from stdin adn convert
    echo '{"name":"Luke"}' | ` + os.Args[0] + ` convert -o yaml
read data from file and convert
  ` + os.Args[0] + ` convert -o yaml -f jsonfile`,
		SilenceErrors: true,
		RunE:          run,
	}
)

func RegisterConvertCommand(cmd *cobra.Command) {
	cmd.AddCommand(convertCmd)
}
func run(cmd *cobra.Command, args []string) error {
	// 获取flag
	output := cmd.Flag("output").Value.String()
	indent := cmd.Flag("indent").Value.String()
	file := cmd.Flag("file").Value.String()
	allowMissingKeysStr := cmd.Flag("allow-missing-keys").Value.String()
	parseBool, err := strconv.ParseBool(allowMissingKeysStr)
	if err != nil {
		return err
	}
	allowMissingKeys := parseBool

	// 标准输入或者文件读取数据
	srcData, err := util.ReadFromFileOrStdin(file)
	if err != nil {
		return err
	}

	if valid := json.Valid(srcData); !valid {
		// 输入的不是json，按照yaml读取
		toJSON, err := yaml.YAMLToJSON(srcData)
		if err != nil {
			return err
		}
		srcData = toJSON
	}

	// 打印
	getPrinter, err := printer.GetPrinter(output, indent, allowMissingKeys)
	if err != nil {
		return err
	}
	if err := getPrinter.Print(os.Stdout, srcData); err != nil {
		return err
	}
	return nil
}
