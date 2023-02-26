# [![Mit License][mit-img]][mit] [![GitHub Release][release-img]][release] [![GoDoc][godoc-img]][godoc] [![Go Report Card][reportcard-img]][reportcard]
# toolman.org/terminal/decor

Package decor provides facilities for decorating a string of characters with
display attributes for the current (or specified) terminal type. This is
done using a notation inspired by (but slightly different from) Zsh prompt
formatting. Supported attributes are __bold__, _italic_, and/or <u>underlined</u>
characters as well as 256-color support for foreground and background colors.

## Terminal Attributes

At its basis, this package is about decorating short blurbs of text to display
in a terminal. As a simple example, consider the following program:

``` go
package main

import (
	"fmt"
	"os"

	"toolman.org/terminal/decor"
)

func main() {
	d, err := decor.New()
	if err != nil {
		panic(err)
	}

	input := "[@B@F{44}@Iuser@i@F{Orchid1}@@@F{Green3}host@f@b]"
	output, err := d.Format(input)
	if err != nil {
		panic(err)
	}

	fmt.Printf("\nTERM....: %q\n\n", os.Getenv("TERM"))
	fmt.Printf("Input...: %q\n\n", input)
	fmt.Printf("Quoted..: %q\n\n", output)
	fmt.Printf("RAW.....: %s\n\n", output)
}
```
When executed, this emits output similar to the following:

![Screenshot #1][ss1-img]

This shows I'm using a terminal of type `rxvt-256color`, the input I'd like to format,
a quoted representation of the formatted output for this terminal and then the output
as rendered by the terminal itself.

Below is a quick-n-dirty chart that walks through each of the attribute markings used above:


![Screenshot #2][ss2-img-dark]
![Screenshot #2][ss2-img-light]


[ss1-img]: img/image1a.png
[ss2-img-dark]: img/ss2c-dark.png#gh-dark-mode-only
[ss2-img-light]: img/ss2b-light.png#gh-light-mode-only

[mit-img]: http://img.shields.io/badge/License-MIT-c41e3a.svg
[mit]: https://github.com/tep/time-timetool/blob/master/LICENSE

[release-img]: https://img.shields.io/github/release/tep/terminal-decor/all.svg
[release]: https://github.com/tep/terminal-decor/releases

[godoc-img]: https://pkg.go.dev/badge/toolman.org/terminal/decor?utm_source=godoc
[godoc]: https://pkg.go.dev/toolman.org/terminal/decor

[reportcard-img]: https://goreportcard.com/badge/toolman.org/terminal/decor
[reportcard]: https://goreportcard.com/report/toolman.org/terminal/decor
