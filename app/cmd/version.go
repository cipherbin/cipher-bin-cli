package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Prints your currently installed version of cipherbin",
	Args:  cobra.MinimumNArgs(0),
	Run:   runVersionCmd,
}

// runVersionCmd defines the behavior thats executed when cipherbin read is ran.
// The version command simply prints a user's currently installed cipherbin version.
func runVersionCmd(cmd *cobra.Command, args []string) {
	fmt.Println(Version)
}
