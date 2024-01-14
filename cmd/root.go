package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/cipherbin/cipher-bin-cli/pkg/api"
	"github.com/spf13/cobra"
)

const (
	// Version constant that represents the current build version
	Version = "v.0.6.0"

	// CLI usage description strings
	usageEmail         = "when provided, a read receipt will be sent to this email upon read/destruction"
	usageRefName       = "requires: email flag. This reference name will be quoted in the read receipt email"
	usagePassword      = "provide an additional password to read the message"
	usageOpenInBrowser = "open and view the message in the web app in your browser"
)

var (
	// APIClient is the exported APIClient, it is set during init and used within commands.
	APIClient *api.Client

	// Command flag variables
	email         string // -e, --email
	referenceName string // -r, --reference_name
	password      string // -p, --password
	openInBrowser bool   // -o, --open
)

func init() {
	client := http.Client{Timeout: 15 * time.Second}
	browserBaseURL := "https://cipherb.in"
	apiBaseURL := "https://api.cipherb.in"

	if os.Getenv("CIPHER_BIN_ENV") == "development" {
		browserBaseURL = "http://localhost:3000"
		apiBaseURL = "http://localhost:4000"
	}

	api, err := api.NewClient(browserBaseURL, apiBaseURL, &client)
	if err != nil {
		fmt.Printf("Error creating API client. Err: %v", err)
		os.Exit(1)
		return
	}

	// Set the globally exported APIClient variable to the new client we created
	APIClient = api

	// Add all cipherbin commands to the base command
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(readCmd)
	rootCmd.AddCommand(versionCmd)

	// Create command flags
	createCmd.Flags().StringVarP(&email, "email", "e", "", usageEmail)
	createCmd.Flags().StringVarP(&referenceName, "reference_name", "r", "", usageRefName)
	createCmd.Flags().StringVarP(&password, "password", "p", "", usagePassword)

	// Read command flags
	readCmd.Flags().BoolVarP(&openInBrowser, "open", "o", false, usageOpenInBrowser)
}

var rootCmd = &cobra.Command{
	Use:   "cipherbin",
	Short: "Cipherbin is a simple CLI tool for generating encrypted messages",
	Long: `
cipherbin is a free and open source service for sending encrypted messages (https://cipherb.in).
This CLI interacts with the api (https://api.cipherb.in) in order to create and read encrypted
messages from the command line
`,
	Args: cobra.MinimumNArgs(1),
	Run:  runRootCmd,
}

func runRootCmd(cmd *cobra.Command, args []string) {
	fmt.Println("Unknown command. Try `cipherbin help` for more information")
}

// Execute runs a user's command. On error, it will print the error and cause
// the program to exit with status code 1
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
