# goColor

goColor is a Go (golang) package for generating colorized output in terminals simply and easily. It supports 256 colors which should be the limit of modern terminal emulators. Other features such as "underline", "bold", and "highlight" are alse included.

<br>
# Insatllation

```shell
go get github.com/nrechn/gocolor
```

<br>
# Example

```go
package main

import (
	"github.com/nrechn/gocolor"
)

func main() {
        // Print in cyan color.
        gocolor.Color("This is cyan", "cyan")
        
        // Print red color text in bold .
        gocolor.Color("This is red in bold", "red", "bold")
        
        // Print in light yellow color.
        gocolor.Color("This is yellow in light", "yellow", "light")
        
        // Print without new line. Allow following texts displayed in the this (same) line.
        gocolor.Color("green", "This is green without new line following.", "sameline")
        
        // Add underline to text displayed.
        gocolor.Color("white", "underline", "This is white with underline")
        
        // Use 256 color number to indicate the color.
        gocolor.Color("123", "This is color 87ffff.")
        
        // "255" is foreground color, "19" is background color.
        gocolor.Color("255", "19", "This text in display in #eeeeee foreground color and #0000af background color")
}
```

<br>
# Feature Detail

- **Don't care arguments order**. You don't need to care arguments order. Just put string as arguments in whatever order you like.
- **Underline** support. goColor can also understand `UnderLine`, `Underline`. Or `ul`, `UL`, `Ul`, etc. if you wold like to use only two character to indicate `underline`.
- **Sameline** support. It will stop the next text display in the new line. Have similar alias as `underline`, you are welcome use `SL`, `sl`, `SameLine`, etc. instead.
- **Bold** support. It will display output in bold. Just pust `bold` or `Bold` as argument in whever you like.
- **Highlight** support. It is actually background. (Foreground is the text color). goColor automatically assumes the first color argument is for foreground color, and the second one is for background.
- **Builtin colors**. The builtin colors are "red", "green", "yellow", "blue", "magenta", "cyan", and "white". It means you can use english word to indicate you would like to use this color for output.
- **Light colors**. If you would like to display text in light color of builtin colors, simply add `light` argument when you call `gocolor.Color` function.
- **256 colors** support. Most modern terminal emulators support 256 colors, and goColor supports it as well. The color number can be found in the 256 color chart below.

![Xterm 256 color chart](https://upload.wikimedia.org/wikipedia/en/1/15/Xterm_256color_chart.svg)

<br>
# Contributing

If you have any suggestion, idea, or bug report; feel free to open an issue on this repository.
