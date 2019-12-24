package cmd

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bradford-hamilton/cipher-bin-cli/internal/api"
	"github.com/spf13/cobra"
)

// Email will be hydrated with it's value if a user runs a cmd with flag --email (or -e)
var Email string

// ReferenceName will be hydrated with it's value if a user runs a cmd with flag --reference_name (or -r)
var ReferenceName string

// Password will be hydrated with it's value if a user runs a cmd with flag --password (or -p)
var Password string

// APIClient is the exported APIClient, it is set during init
var APIClient *api.Client

func init() {
	client := http.Client{Timeout: 15 * time.Second}

	browserBaseURL := "https://cipherb.in"
	apiBaseURL := "https://api.cipherb.in"

	api, err := api.NewClient(browserBaseURL, apiBaseURL, &client)
	if err != nil {
		fmt.Printf("Error creating API client. Err: %v", err)
		os.Exit(1)
		return
	}

	// Set the globally exported APIClient variable to the new client we created
	APIClient = api

	// Add the other commands to the base command
	rootCmd.AddCommand(createCmd)

	// Hydrate variables with any user provided flags
	createCmd.Flags().StringVarP(&Email, "email", "e", "", "when provided, a read receipt will be sent to this email upon read/destruction")
	createCmd.Flags().StringVarP(&ReferenceName, "reference_name", "r", "", "requires: email flag. This reference name will be quoted in the read receipt email")
	createCmd.Flags().StringVarP(&Password, "password", "p", "", "provide an additional password to read the message")
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
