// Package colorit ... TODO: pckg comment
package colorit

import (
	"io"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/term"
)

// runWithStdin runs the command `name` with `args`, writes `input` to stdin,
// and returns the command's stdout as a string.
// If the command fails or returns a non-zero exit code, it returns "".
func runWithStdin(input, name string, args ...string) string {
	cmd := exec.Command(name, args...) //nolint:noctx
	cmd.Stdin = strings.NewReader(input)
	out, err := cmd.Output()
	if err != nil {
		return ""
	}
	return string(out)
}

// Highlighter implementation provides logic for syntax highlighting with
// ANSI Escape codes.
type Highlighter interface {
	// Name should not return ""
	Name() string
	// If Highlight returns same text or "" it means that
	// it cannot highlight it for given syntax
	Highlight(text, syntax string) string
}

// FilterHighliters filters and reorders highliters against filter string
// that contains names of highlighters separated by ";": "a;b;c".
// If filter string is "", returns all highlighters from src.
// It returns new slice.
func FilterHighliters(filter string, src []Highlighter) []Highlighter {
	filter = strings.ToLower(strings.TrimSpace(filter))
	if filter == "disable" {
		return []Highlighter{}
	}
	filters := strings.Split(filter, ";")
	for i := range filters {
		filters[i] = strings.TrimSpace(filters[i])
	}
	filtered := make([]Highlighter, 0, len(src))
	for _, f := range filters {
		if f == "" {
			continue
		}
		for _, prov := range src {
			if prov.Name() == f {
				filtered = append(filtered, prov)
				continue
			}
		}
	}
	return filtered
}

// Highlight tries each of provided highlighters until one of them
// successfully, highlights provided text with provided syntax or just
// returns original text if there is no suitable highlighters.
func Highlight(text, syntax string, highlighters []Highlighter) string {
	syntax = strings.Trim(strings.ToLower(strings.TrimSpace(syntax)), "-")
	for _, highlighter := range highlighters {
		result := highlighter.Highlight(text, syntax)
		trimmedResult := strings.TrimSpace(result)
		// Continue if highlighter returns nothing
		if trimmedResult == "" {
			continue
		}
		// Continue if highlighter returns same text
		if trimmedResult == strings.TrimSpace(text) {
			continue
		}
		return result
	}
	return text
}

var defaultHighlighters = []Highlighter{
	BatHighlighter{},
	PygmentsHighlighter{},
	RichHighlighter{},
	ChromaHighlighter{},
}

// DefaultHighlighters returns slice of all builtin highlighters filtered
// with filter string from GO_COLORIT env var.
func DefaultHighlighters() []Highlighter {
	selected := os.Getenv("GO_COLORIT")
	return FilterHighliters(selected, defaultHighlighters)
}

// HighlightStr highlights text for syntax using default highlighters slice.
func HighlightStr(text, syntax string) string {
	return Highlight(text, syntax, DefaultHighlighters())
}

// IsTTY returns true if w refers to a terminal (tty).
func isTTY(w io.Writer) bool {
	if w == nil {
		return false
	}
	// any type that exposes Fd() uintptr (os.File and many file-like types)
	type fdWriter interface{ Fd() uintptr }
	if fw, ok := w.(fdWriter); ok {
		return term.IsTerminal(int(fw.Fd()))
	}
	return false
}

// HighlightTo highlights text for syntax using default highlighters slice
// and writes it to out if out is a TTY else writes original text.
func HighlightTo(text, syntax string, out io.Writer) error {
	if isTTY(out) {
		text = HighlightStr(text, syntax)
	}
	_, err := io.WriteString(out, text)
	return err //nolint:wrapcheck
}
