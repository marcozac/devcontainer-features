package devcontainerfeatures

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "devcontainer-features",
}

func init() {
	rootCmd.AddCommand(newCmd, generateReadmeCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
