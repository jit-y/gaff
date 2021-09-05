package gaff

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type RootOptions struct {
	Filepath string
}

func NewRootCmd() *cobra.Command {
	opts := &RootOptions{}
	rootCmd := &cobra.Command{
		Use: "gaff <command> [flags]",
		RunE: func(cmd *cobra.Command, args []string) error {
			filesys := os.DirFS(".")
			p := NewParser(filesys)
			strategy, diags := p.LoadHCLFile(opts.Filepath)

			if diags.HasErrors() {
				return fmt.Errorf("%s", diags.Error())
			}

			fmt.Printf("%+v", strategy)

			return nil
		},
	}

	rootCmd.PersistentFlags().StringVarP(&opts.Filepath, "file", "f", "", "")

	return rootCmd
}
