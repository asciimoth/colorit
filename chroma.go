package colorit

// ChromaHighlighter implement syntax highlighting with chroma cli tool.
type ChromaHighlighter struct{}

// Name returns "chroma".
func (p ChromaHighlighter) Name() string {
	return "chroma"
}

// Highlight returns text highlighted for syntax or "" if faled.
func (p ChromaHighlighter) Highlight(text, syntax string) string {
	return runWithStdin(text, "chroma", "-l", syntax)
}
