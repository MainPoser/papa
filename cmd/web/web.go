package main

import (
	goflag "flag"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"k8s.io/klog/v2"
)

// Options contains everything necessary to create and run a proxy server.
type Options struct {
}

// NewOptions returns initialized Options
func NewOptions() *Options {
	return &Options{}
}

// Complete completes all the required options.
func (o *Options) Complete() error {
	return nil
}

// Validate validates all the required options.
func (o *Options) Validate() error {
	return nil
}

// Run runs the specified ProxyServer.
func (o *Options) Run() error {
	return nil
}

func (o *Options) AddFlags(fs *pflag.FlagSet) {

}

func NewWebCommand() *cobra.Command {
	opts := NewOptions()

	cmd := &cobra.Command{
		Use:  os.Args[0],
		Long: `start a web server.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.Complete(); err != nil {
				return fmt.Errorf("failed complete: %w", err)
			}

			if err := opts.Validate(); err != nil {
				return fmt.Errorf("failed validate: %w", err)
			}

			if err := opts.Run(); err != nil {
				klog.ErrorS(err, "Error running WebServer")
				return err
			}

			return nil
		},
		Args: func(cmd *cobra.Command, args []string) error {
			for _, arg := range args {
				if len(arg) > 0 {
					return fmt.Errorf("%q does not take any arguments, got %q", cmd.CommandPath(), args)
				}
			}
			return nil
		},
	}
	fs := cmd.Flags()
	opts.AddFlags(fs)
	fs.AddGoFlagSet(goflag.CommandLine) // for --boot-id-file and --machine-id-file

	// 可以给cmp添加子命令
	return cmd
}

func main() {
	// 初始化Klog
	klog.InitFlags(goflag.CommandLine)
	// 退出前刷新klog
	defer klog.Flush()
	// 新建command
	cmd := NewWebCommand()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
