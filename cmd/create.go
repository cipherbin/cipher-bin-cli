package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/briandowns/spinner"
	"github.com/cipherbin/cipher-bin-cli/pkg/aes256"
	"github.com/cipherbin/cipher-bin-cli/pkg/colors"
	"github.com/cipherbin/cipher-bin-cli/pkg/editor"
	"github.com/cipherbin/cipher-bin-cli/pkg/randstring"
	"github.com/cipherbin/cipher-bin-server/db"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new encrypted message",
	Long: `
Create opens up either your editor of choice if you have $EDITOR set, or it
will default to vi. Type/paste your message content into the editor, save,
and close. Your message will be encrypted and you will receive the one time
use link.
	`,
	Args: cobra.MinimumNArgs(0),
	Run:  runCreateCmd,
}

func runCreateCmd(cmd *cobra.Command, args []string) {
	// Capture user input (the message) within their preferred editor
	msgBytes, err := editor.CaptureInput(editor.PreferredEditor)
	if err != nil {
		colors.Println(err.Error(), colors.Red)
		os.Exit(1)
	}

	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	s.Color("blue")
	s.Start()

	// Create a v4 uuid for message identification and to eliminate
	// almost any chance of stumbling upon the url
	uuidv4 := uuid.NewV4().String()

	// Generate a random 32 byte string
	key, err := randstring.New(32)
	if err != nil {
		colors.Println(err.Error(), colors.Red)
		os.Exit(1)
	}

	// Encrypt the message using AES-256
	encryptedMsg, err := aes256.Encrypt(msgBytes, key)
	if err != nil {
		colors.Println(err.Error(), colors.Red)
		os.Exit(1)
	}

	// Create one time use URL with format {host}?bin={uuidv4};{ecryption_key}
	oneTimeURL := fmt.Sprintf("%s/msg?bin=%s;%s", apiClient.BrowserBaseURL, uuidv4, key)

	msg := db.Message{
		UUID:          uuidv4,
		Message:       encryptedMsg,
		Email:         email,
		ReferenceName: referenceName,
		Password:      password,
	}

	// Post message to the cipherbin api
	err = apiClient.PostMessage(&msg)
	if err != nil {
		colors.Println(err.Error(), colors.Red)
		os.Exit(1)
	}

	// Stop the spinner and create warning message
	s.Stop()
	warning := "\nWarning! This message will self destruct after reading it."

	// Copy the one time url to the user's clipboard. Using nice little package here
	// that does the work around ensuring this works on OSX, Windows, Linux/Unix
	if err := clipboard.WriteAll(oneTimeURL); err != nil {
		colors.Println(warning, colors.Yellow)
		colors.Println(oneTimeURL+"\n", colors.Green)
		colors.Println(err.Error(), colors.Red)
		os.Exit(1)
	}

	colors.Println(warning, colors.Yellow)
	colors.Println(oneTimeURL+"\n", colors.Green)
}
