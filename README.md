# colorit
[![asciicast](https://asciinema.org/a/740664.svg)](https://asciinema.org/a/740664)

`colorit` is a tiny Go helper that adds optional syntax highlighting for your CLI tools output by calling existing external highlighter tools (e.g. `bat`, `chroma`, [etc](#supported-tools)). It keeps your tool small - if no highlighter is available in working env it simply prints plain text.

## Install
```sh
go get github.com/asciimoth/colorit
```

## Quick example
```go
package main

import (
	"os"

	"github.com/asciimoth/colorit"
)

const help = `...` // your help text

func main() {
	// text, syntax name, destination writer
	// If output is not a TTY or no highlighter is found, text is left unchanged.
	colorit.HighlightTo(help, "help", os.Stdout)
}
```

## Behavior
* Attempts to highlight using available external highlighter CLI tools.
* If no configured highlighter is present or the output is not a terminal, the original text is written unchanged.
* Keeps highlighting optional - users only get colored output if they have a supported highlighter installed.

## Configuration
Set `GO_COLORIT` env var to control which highlighters are tried and in which order. Use a semicolon-separated list, for example:
```
GO_COLORIT="chroma;rich;bat"
```

Set `GO_COLORIT=disable` to force no highlighting.

## Supported tools
- [bat](https://github.com/sharkdp/bat)
- [pygmentize](https://pygments.org/)
- [rich-cli](https://github.com/Textualize/rich-cli)
- [chroma](https://github.com/alecthomas/chroma)

## Contributing
If you know other standalone highlighter CLI tools that could be integrated, please open an issue or PR.

## License
This project is licensed under either of

- Apache License, Version 2.0, ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
- MIT license ([LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT)

at your option.

