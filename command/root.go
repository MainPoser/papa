package command

import (
	"github.com/MainPoser/papa/command/convert"
	"github.com/MainPoser/papa/command/patch"
	"github.com/spf13/cobra"
	"os"
)

var (
	rootCmd = &cobra.Command{
		Use:   os.Args[0],
		Short: "just use --help, you will know everything",
		Long:  "shiny shiny shiny, support some tiny tool, just use --help ",
	}
)

func init() {
	rootCmd.PersistentFlags().StringP("file", "f", "", "file for read")
	rootCmd.PersistentFlags().StringP("output", "o", "json", "output {json|yaml|jsonpath=...}")
	rootCmd.PersistentFlags().String("indent", "", "json pretty indent")
	rootCmd.PersistentFlags().Bool("allow-missing-keys", false, "when --output is jsonpath, use")
	convert.RegisterConvertCommand(rootCmd)
	patch.RegisterPatchCommand(rootCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
