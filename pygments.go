package colorit

// PygmentsHighlighter implement syntax highlighting with pygmentize cli tool.
type PygmentsHighlighter struct{}

// Name returns "pygments".
func (p PygmentsHighlighter) Name() string {
	return "pygments"
}

// Highlight returns text highlighted for syntax or "" if faled.
func (p PygmentsHighlighter) Highlight(text, syntax string) string {
	return runWithStdin(text, "pygmentize", "-l", syntax)
}
