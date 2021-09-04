package gaff

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "gaff <command> [flags]",
	}

	return rootCmd
}
