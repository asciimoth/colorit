// Package colorit ... TODO: pckg comment
package colorit

import (
	"os/exec"
	"strings"
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
