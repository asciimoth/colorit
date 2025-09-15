package colorit

// RichHighlighter implement syntax highlighting with rich cli tool.
type RichHighlighter struct{}

// Name returns "rich".
func (p RichHighlighter) Name() string {
	return "rich"
}

// Highlight returns text highlighted for syntax or "" if faled.
func (p RichHighlighter) Highlight(text, syntax string) string {
	return runWithStdin(text, "rich", "-", "--force-terminal", "-x", syntax)
}
