package fansi

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

// https://github.com/mgutz/ansi/blob/master/ansi.go
type PromptOptions struct {
	in  io.Reader
	out io.Writer
}

// Prompt encapsulates the functionality necessary to prompt
// the user for various kinds of inputs in an interactive program
type Prompt struct {
	in  *bufio.Reader
	out *bufio.Writer
}

// NewPrompt creates an instance of Prompt
func NewPrompt(opts PromptOptions) *Prompt {
	return &Prompt{
		in:  bufio.NewReader(opts.in),
		out: bufio.NewWriter(opts.out),
	}
}

func NewStdPrompt() *Prompt {
	return NewPrompt(PromptOptions{
		in:  os.Stdin,
		out: os.Stdout,
	})
}

type nullWriter struct{}

func (dn nullWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func NewTestPrompt(input string) *Prompt {
	inbuf := bytes.NewBufferString(input)
	return NewPrompt(PromptOptions{
		in:  inbuf,
		out: nullWriter{},
	})
}

// GetInput prints the message and prompts the user for string input.
// It returns the input up until a newline, or an error if an interrupt
// signal was sent.
func (p *Prompt) GetInput(msg string) (string, error) {
	fmt.Fprintln(p.out, msg)
	fmt.Fprint(p.out, "-> ")
	p.out.Flush()
	text, err := p.in.ReadString('\n')
	return strings.TrimSpace(text), err
}

// GetSecretInput works like GetInput, but displays a configurable
// character instead of the current buffer, in the style of a password
// input.
// func (p *Prompt) GetSecretInput(msg string) (string, error) {
// 	ok := false
// 	bytePassword := []byte{}

// 	for !ok {
// 		fmt.Println(msg)
// 		fmt.Print("-> ")
// 		bytePassword, _ = terminal.ReadPassword(int(syscall.Stdin))
// 		fmt.Println()

// 		fmt.Println("Enter again to confirm")
// 		fmt.Print("-> ")
// 		bytePasswordConfirm, _ := terminal.ReadPassword(int(syscall.Stdin))
// 		fmt.Println()

// 		ok = bytes.Compare(bytePassword, bytePasswordConfirm) == 0
// 		if !ok {
// 			fmt.Println("Entries did not match!")
// 		}
// 	}

// 	return strings.TrimSpace(string(bytePassword)), nil
// }

func (p *Prompt) AskYesNo(msg string) bool {
	fmt.Fprintf(p.out, "%s [Y/n] ", msg)
	p.out.Flush()
	text, _ := p.in.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)
	return !strings.Contains(strings.ToLower(text), "n")
}
