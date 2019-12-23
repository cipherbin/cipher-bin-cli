// Package editor was sourced from this article and then updated:
// https://samrapdev.com/capturing-sensitive-input-with-editor-in-golang-from-the-cli
package editor

import (
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// DefaultEditor will always fall back to vim
const DefaultEditor = "vim"

// PreferredEditorResolver is a function that returns an editor that the user
// prefers to use, such as the configured `$EDITOR` environment variable.
type PreferredEditorResolver func() string

// PreferredEditor returns the user's editor as defined by the `$EDITOR`
// environment variable, or the `DefaultEditor` if it is not set.
func PreferredEditor() string {
	if editor := os.Getenv("EDITOR"); editor != "" {
		return editor
	}

	return DefaultEditor
}

// More editor support can be added here, for now only support vs code (and default to vim)
func resolveArguments(executable string, filename string) []string {
	args := []string{filename}

	if strings.Contains(executable, "Visual Studio Code.app") || strings.Contains(executable, "code") {
		args = append([]string{"--wait"}, args...)
	}

	return args
}

// OpenFile opens filename in a text editor
func OpenFile(filename string, resolveEditor PreferredEditorResolver) error {
	// Get the full executable path for the editor
	executable, err := exec.LookPath(resolveEditor())
	if err != nil {
		return err
	}

	command := exec.Command(executable, resolveArguments(executable, filename)...)
	command.Stdin = os.Stdin
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	return command.Run()
}

// CaptureInput opens a temporary file in a text editor and returns
// the written bytes on success or an error on failure. It handles deletion
// of the temporary file behind the scenes.
func CaptureInput(resolveEditor PreferredEditorResolver) ([]byte, error) {
	// Create a tmp file that we will use to capture user input
	file, err := ioutil.TempFile(os.TempDir(), "*")
	if err != nil {
		return []byte{}, err
	}

	filename := file.Name()

	// Defer removal of the tmp file
	defer os.Remove(filename)

	// Attempt to close the file and if there is an error, return it
	if err = file.Close(); err != nil {
		return []byte{}, err
	}

	// Attempt to OpenFile and if there is an error, return it
	if err = OpenFile(filename, resolveEditor); err != nil {
		return []byte{}, err
	}

	// Attempt to Readfile to bytes and if there is an error, return it
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
