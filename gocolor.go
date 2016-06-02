// Copyright © 2016 nrechn <nrechn@gmail.com>
//
// This file is part of gocolor.
//
// gocolor is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// gocolor is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with gocolor. If not, see <http://www.gnu.org/licenses/>.
//

package gocolor

import (
	"fmt"
	"regexp"
)

var (
	colorName   = []string{"^[Rr]ed$", "^[Gg]reen$", "^[Yy]ellow$", "^[Bb]lue$", "^[Mm]agenta$", "^[Cc]yan$", "^[Ww]hite$"}
	colorNormal = []string{"1", "2", "3", "4", "5", "6", "7"}
	colorLight  = []string{"9", "10", "11", "12", "13", "14", "15"}
)

type colorOption struct {
	fg, bg, bold, style, msg string
	newLine, underLine       bool
}

func colorCode(color, style string) string {
	var code string = ""
	for i := 0; i < cap(colorName); i++ {
		if matched, _ := regexp.MatchString(colorName[i], color); matched {
			if style == "normal" {
				code = colorNormal[i]
			} else if style == "light" {
				code = colorLight[i]
			}
		}
	}
	if code == "" {
		return color
	} else {
		return code
	}
}

func colorPrint(p colorOption) {
	if p.fg != "" {
		p.fg = "\033[38;5;" + colorCode(p.fg, p.style) + "m"
	}
	if p.bg != "" {
		p.bg = "\033[48;5;" + colorCode(p.bg, p.style) + "m"
	}
	if p.bold != "" {
		p.bold = "\033[1m"
	}
	if p.newLine {
		if p.underLine {
			fmt.Println("\033[4m" + p.fg + p.bg + p.bold + p.msg + "\033[0m")
		} else {
			fmt.Println(p.fg + p.bg + p.bold + p.msg + "\033[0m")
		}
	} else {
		if p.underLine {
			fmt.Print("\033[4m" + p.fg + p.bg + p.bold + p.msg + "\033[0m")
		} else {
			fmt.Print(p.fg + p.bg + p.bold + p.msg + "\033[0m")
		}
	}
}

func checkIfColor(color string) bool {
	for i := 0; i < cap(colorName); i++ {
		matchColor, _ := regexp.MatchString(colorName[i], color)
		if matchColor {
			return true
		}
	}
	return false
}

func checkIfNewLine(arg string) bool {
	matchNewLine, _ := regexp.MatchString("^[Ss][A-Za-z]{0,6}[Lle]$", arg)
	if matchNewLine {
		return true
	} else {
		return false
	}
}

func checkIfUnderLine(arg string) bool {
	matchUnderLine, _ := regexp.MatchString("^[Uu][A-Za-z]{0,7}[Lle]$", arg)
	if matchUnderLine {
		return true
	} else {
		return false
	}
}

func checkIfColorNum(arg string) bool {
	matchColorNum, _ := regexp.MatchString("^[1-9][0-9]{0,2}$", arg)
	if matchColorNum {
		return true
	} else {
		return false
	}
}

func checkIfBold(arg string) bool {
	matchBold, _ := regexp.MatchString("^[Bb]old$", arg)
	if matchBold {
		return true
	} else {
		return false
	}
}

func checkIfLight(arg string) bool {
	matchLight, _ := regexp.MatchString("^[Ll]ight$", arg)
	if matchLight {
		return true
	} else {
		return false
	}
}

func Color(args ...string) {
	var (
		newLineChanged   bool = false
		underLineChanged bool = false
		msgChanged       bool = false
	)
	c := colorOption{style: "normal", newLine: true, underLine: false}
	for _, arg := range args {
		if checkIfNewLine(arg) && c.newLine && !newLineChanged {
			c.newLine = false
			newLineChanged = true
		} else if checkIfUnderLine(arg) && !c.underLine && !underLineChanged {
			c.underLine = true
			underLineChanged = true
		} else if (checkIfColor(arg) || checkIfColorNum(arg)) && (c.fg == "" || c.bg == "") {
			if c.fg == "" {
				c.fg = arg
			} else if c.bg == "" {
				c.bg = arg
			}
		} else if checkIfBold(arg) && c.bold == "" {
			c.bold = "bold"
		} else if checkIfLight(arg) && c.style == "normal" {
			c.style = "light"
		} else if !msgChanged {
			c.msg = arg
			msgChanged = true
		}
	}
	colorPrint(c)
}
