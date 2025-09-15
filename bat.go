package colorit

// BatHighlighter implement syntax highlighting with bat cli tool.
type BatHighlighter struct{}

// Name returns "bat".
func (p BatHighlighter) Name() string {
	return "bat"
}

// Highlight returns text highlighted for syntax or "" if faled.
func (p BatHighlighter) Highlight(text, syntax string) string {
	return runWithStdin(text, "bat", "-pfl", syntax)
}
