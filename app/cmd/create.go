package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/atotto/clipboard"
	"github.com/bradford-hamilton/cipher-bin-cli/pkg/aes256"
	"github.com/bradford-hamilton/cipher-bin-cli/pkg/colors"
	"github.com/bradford-hamilton/cipher-bin-cli/pkg/editor"
	"github.com/bradford-hamilton/cipher-bin-cli/pkg/randstring"
	"github.com/bradford-hamilton/cipher-bin-server/db"
	"github.com/briandowns/spinner"
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
		fmt.Println(err)
		os.Exit(1)
	}

	s := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	s.Color("blue")
	s.Start()

	// Create a v4 uuid for message identification and to eliminate
	// almost any chance of stumbling upon the url
	uuidv4 := uuid.NewV4().String()

	// Generate a random 32 byte string
	key := randstring.New(32)

	// Encrypt the message using AES-256
	encryptedMsg, err := aes256.Encrypt(msgBytes, key)

	// Create one time use URL with format {host}?bin={uuidv4};{ecryption_key}
	oneTimeURL := fmt.Sprintf("%s/msg?bin=%s;%s", APIClient.BrowserBaseURL, uuidv4, key)

	msg := db.Message{
		UUID:          uuidv4,
		Message:       encryptedMsg,
		Email:         Email,
		ReferenceName: ReferenceName,
		Password:      Password,
	}

	// Post message to the cipherbin api
	err = APIClient.PostMessage(&msg)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Stop the spinner and create warning message
	s.Stop()
	w := fmt.Sprintf("\nWarning! This message will self destruct after reading it.")

	// Copy the one time url to the user's clipboard. Using nice little package here
	// that does the work around ensuring this works on OSX, Windows 7, Linux/Unix
	if err := clipboard.WriteAll(oneTimeURL); err != nil {
		colors.Println(w, "yellow")
		colors.Println(oneTimeURL+"\n", "green")
		fmt.Println(err)
		os.Exit(1)
	}

	colors.Println(w, "yellow")
	colors.Println(oneTimeURL+"\n", "green")
}
