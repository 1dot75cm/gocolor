package main

import (
	"github.com/nrechn/gocolor"
)

func main() {
	gocolor.Color("\n\n  ", "bold", "sl")
	gocolor.Color("g", "bold", "sl", "27")
	gocolor.Color("o", "bold", "sl", "209")
	gocolor.Color("C", "bold", "sl", "231", "99")
	gocolor.Color("o", "bold", "sl", "231", "99")
	gocolor.Color("l", "bold", "sl", "231", "99")
	gocolor.Color("o", "bold", "sl", "231", "99")
	gocolor.Color("r", "bold", "sl", "231", "99")

	gocolor.Color(" ", "bold", "sl")
	gocolor.Color("l", "sl", "159", "ul")
	gocolor.Color("o", "sl", "223", "ul")
	gocolor.Color("v", "sl", "171", "ul")
	gocolor.Color("e", "sl", "203", "ul")

	gocolor.Color(" ", "bold", "sl")
	gocolor.Color("c", "bold", "sl", "light", "red")
	gocolor.Color("o", "bold", "sl", "light", "green")
	gocolor.Color("l", "bold", "sl", "light", "yellow")
	gocolor.Color("o", "bold", "sl", "light", "blue")
	gocolor.Color("r", "bold", "sl", "light", "magenta")
	gocolor.Color("\n\n ", "bold")
}
