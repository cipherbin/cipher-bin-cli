<p align="center">
  <img src="cipher_bin_logo_black.png" alt="cipher bin logo" />
  <h1 align="center">Cipher Bin CLI</h1>
</p>

Source code for the cli, if you are looking for the client side React app [go here](https://github.com/bradford-hamilton/cipher-bin-client). If you are looking for the Golang server [go here](https://github.com/bradford-hamilton/cipher-bin-server)

## Installation
_**Option A:**_

Homebrew:
```
brew tap bradford-hamilton/cipherbin
brew install bradford-hamilton/cipherbin/cipherbin
```

Updating:
```
brew upgrade cipherbin
```

Uninstalling:
```
brew uninstall cipherbin
```

_**Option B:**_

If you mosey on over to [releases](https://github.com/bradford-hamilton/cipher-bin-cli/releases), you'll find binaries for darwin, linux, and amd64. You can download directly from there.

_**Option C:**_

If you have Go installed on your machine, use `go install`:
```
go install github.com/bradford-hamilton/cipher-bin-cli
```

This will place the binary in your `go/bin` and is ready to use, however the binary will be named `cipher-bin-cli` with this option.

The alternative solution here is to run `go build -o $GOPATH/bin/cipherbin`. This will essentially act like a `go install`, but you can name the binary what it's intended to be named.

## Using the CLI
___
**Creating a new message:**

The `create` command will open either your editor of choice (if you have $EDITOR env var set), or default to vi. As of now there is only specific support for VS Code. Other editors _may_ work, but it's not guaranteed. Within the editor, type or paste your secret content. When you save and quit your message will be encrypted and posted to the cipherbin api. The one time use URL will be automatically copied to your clipboard and printed in your teminal. It works similarly to a `git commit --amend` work-flow. The URL can now either be visited in a browser or the message can be read with the `read` command.
```
cipherbin create
```

**Flags:**

Email to send the notification to when your message is read and destroyed
```
--email, -e
```

Reference name for the message (Ex. "prod env vars"). You must be using the email flag for this to have any effect.
```
--reference_name, -r
```

<!-- Add a password to be able to read the message
```
--password, -p
``` -->

___
**Reading an encrypted message:**

Instead of visiting the URL in your browser you can use the `read` command. It takes one argument, which is the URL.
```
cipherbin read https://cipherb.in/msg?bin=some_uuid;some_key
```
___
# Development
## Running the application
Build
```
go build -o cipherbin main.go
```

Run
```
./cipherbin [commands...] [flags...]
```

Or for quicker iterations:
```
go run main.go [commands...] [flags...]
```
