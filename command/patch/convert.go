package patch

import (
	"encoding/json"

	"os"
	"strconv"

	"github.com/MainPoser/papa/pkg/printer"
	"github.com/MainPoser/papa/pkg/util"
	json_patch "github.com/evanphx/json-patch"
	"github.com/ghodss/yaml"
	"github.com/spf13/cobra"
)

var (
	patchCmd = &cobra.Command{
		Use:   "patch",
		Short: "patch data",
		Long:  "patch, like <Angle Luke> to <Like Skywalker>",
		Example: `read data from stdin and patch by json type
    echo '{"name":"Luke"}' | ` + os.Args[0] + ` patch -t json -p '[{"op":"replace","path":"/name","value":"change"}]' -o yaml
read data from file and patch by merge type
  ` + os.Args[0] + ` patch -p '{"name":"change"}' -o yaml -f jsonfile`,
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE:          run,
	}
)

func init() {
	patchCmd.Flags().StringP("patch-data", "p", "{}", "patch data")
	patchCmd.Flags().StringP("patch-type", "t", "merge", "patch type")
}

func RegisterPatchCommand(cmd *cobra.Command) {
	cmd.AddCommand(patchCmd)
}

func run(cmd *cobra.Command, args []string) error {
	// 获取flag
	output := cmd.Flag("output").Value.String()
	indent := cmd.Flag("indent").Value.String()
	patchData := cmd.Flag("patch-data").Value.String()
	patchType := cmd.Flag("patch-type").Value.String()
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

	// 根据不同的策略更新
	if patchType == "merge" {
		// merge更新
		patch, err := json_patch.MergePatch(srcData, []byte(patchData))
		if err != nil {
			return err
		}
		srcData = patch
	}
	if patchType == "json" {
		// json策略更新
		patch, err := json_patch.DecodePatch([]byte(patchData))
		if err != nil {
			return err
		}
		apply, err := patch.Apply(srcData)
		if err != nil {
			return err
		}
		srcData = apply
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
