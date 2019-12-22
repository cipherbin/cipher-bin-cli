package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Email will be hydrated with it's value if a user runs a cmd with flag --email (or -e)
var Email string

// ReferenceName will be hydrated with it's value if a user runs a cmd with flag --reference_name (or -r)
var ReferenceName string

// Password will be hydrated with it's value if a user runs a cmd with flag --password (or -p)
var Password string

func init() {
	rootCmd.Flags().StringVarP(&Email, "email", "e", "", "when provided, a read receipt will be sent to this email upon read/destruction")
	rootCmd.Flags().StringVarP(&ReferenceName, "reference_name", "r", "", "requires: email flag. This reference name will be quoted in the read receipt email")
	rootCmd.Flags().StringVarP(&Password, "password", "p", "", "provide an additional password to read the message")
}

var rootCmd = &cobra.Command{
	Use:   "cipherbin",
	Short: "Cipherbin is a simple CLI tool for generating encrypted messages",
	Long:  `TODO: Long description`,
	Args:  zeroArgs,
	Run:   runRootCmd,
}

func runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Println(Email)
	fmt.Println(ReferenceName)
	fmt.Println(Password)

	// Open a tmp file (or a user editable file at a known location like `~/.cipherbin`?
	// Either read the file on close for tmp file scenario or if there is known location,
	// always read from there when running create cmd

	// Create a uuidv4

	// Encrypt the message useing AES-256

	// Create URL with format {host}?bin={uuidv4};{ecryption_key}

	// Send ecrypted message to server in same shape as the front end does for hopeful
	// plug and play
}

func zeroArgs(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		return errors.New("Requires 0 arguments")
	}

	return nil
}

// Execute runs a user's command. On error, it will print the error and cause
// the program to exit with status code 1
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
