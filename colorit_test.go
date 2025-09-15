package colorit_test

import (
	"slices"
	"testing"

	"github.com/asciimoth/colorit"
)

type mockHighlighter struct {
	name string
}

func (h mockHighlighter) Name() string {
	return h.name
}

func (h mockHighlighter) Highlight(text, syntax string) string {
	return h.name + ":" + text + ":" + syntax
}

func TestFilterHighliters(t *testing.T) {
	t.Parallel()
	filter := "c;b;;;d;;;"
	highlighters := []colorit.Highlighter{
		mockHighlighter{"a"},
		mockHighlighter{"b"},
		mockHighlighter{"c"},
		mockHighlighter{"d"},
		mockHighlighter{"e"},
	}
	estimated := []colorit.Highlighter{
		mockHighlighter{"c"},
		mockHighlighter{"b"},
		mockHighlighter{"d"},
	}

	filtered := colorit.FilterHighliters(filter, highlighters)

	if !slices.Equal(filtered, estimated) {
		t.Error(estimated, filtered)
	}
}

func TestHighlit(t *testing.T) {
	t.Parallel()
	highlighters := []colorit.Highlighter{
		mockHighlighter{"e"},
		mockHighlighter{"c"},
		mockHighlighter{"d"},
	}
	estimated := "e:txt:syn"

	result := colorit.Highlight("txt", "syn", highlighters)
	if result != estimated {
		t.Error(estimated, result)
	}
}
