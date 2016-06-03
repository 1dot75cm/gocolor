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
	colorName           = []string{"^[Rr]ed$", "^[Gg]reen$", "^[Yy]ellow$", "^[Bb]lue$", "^[Mm]agenta$", "^[Cc]yan$", "^[Ww]hite$"}
	colorNormal         = []string{"1", "2", "3", "4", "5", "6", "7"}
	colorLight          = []string{"9", "10", "11", "12", "13", "14", "15"}
	fgcolorNon256Normal = []string{"31", "32", "33", "34", "35", "36", "97"}
	fgcolorNon256Light  = []string{"91", "92", "93", "94", "95", "96", "97"}
	bgcolorNon256Normal = []string{"41", "42", "43", "44", "45", "46", "107"}
	bgcolorNon256Light  = []string{"101", "102", "103", "104", "105", "106", "107"}
)

type colorOption struct {
	fg, bg, bold, style, msg   string
	newLine, underLine, non256 bool
}

func colorCode(co colorOption, color string) string {
	var code string
	for i := 0; i < cap(colorName); i++ {
		if matched, _ := regexp.MatchString(colorName[i], color); matched {
			if co.style == "normal" {
				if co.non256 {
					if color == co.fg {
						code = fgcolorNon256Normal[i]
					} else {
						code = bgcolorNon256Normal[i]
					}
				} else {
					code = colorNormal[i]
				}
			} else if co.style == "light" {
				if co.non256 {
					if color == co.fg {
						code = fgcolorNon256Light[i]
					} else {
						code = fgcolorNon256Light[i]
					}
				} else {
					code = colorLight[i]
				}
			}
		}
	}
	if code == "" {
		return color
	}
	return code
}

func colorPrint(p colorOption) {
	if p.non256 {
		if p.fg != "" {
			p.fg = "\033[" + colorCode(p, p.fg) + "m"
		}
		if p.bg != "" {
			p.bg = "\033[" + colorCode(p, p.bg) + "m"
		}
	} else {
		if p.fg != "" {
			p.fg = "\033[38;5;" + colorCode(p, p.fg) + "m"
		}
		if p.bg != "" {
			p.bg = "\033[48;5;" + colorCode(p, p.bg) + "m"
		}
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

func checkColor(arg string) bool {
	for i := 0; i < cap(colorName); i++ {
		matchColor, _ := regexp.MatchString(colorName[i], arg)
		if matchColor {
			return true
		}
	}
	return false
}

func check(matchRegExp, arg string) bool {
	matchString, _ := regexp.MatchString(matchRegExp, arg)
	if matchString {
		return true
	}
	return false
}

func checkIf(option, arg string) bool {
	switch option {
	case "color":
		return checkColor(arg)
	case "newLine":
		return check("^[Ss][A-Za-z]{0,6}[Lle]$", arg)
	case "underLine":
		return check("^[Uu][A-Za-z]{0,7}[Lle]$", arg)
	case "colorNum":
		return check("^[1-9][0-9]{0,2}$", arg)
	case "bold":
		return check("^[Bb]old$", arg)
	case "light":
		return check("^[Ll]ight$", arg)
	case "-256":
		return check("-256", arg)
	}
	return false
}

func Color(args ...string) {
	var (
		newLineChanged, underLineChanged, msgChanged, non256Changed bool
	)
	c := colorOption{style: "normal", newLine: true, underLine: false, non256: false}
	for _, arg := range args {
		if checkIf("newLine", arg) && c.newLine && !newLineChanged {
			c.newLine = false
			newLineChanged = true
		} else if checkIf("underLine", arg) && !c.underLine && !underLineChanged {
			c.underLine = true
			underLineChanged = true
		} else if checkIf("-256", arg) && !c.non256 && !non256Changed {
			c.non256 = true
			non256Changed = true
		} else if (checkIf("color", arg) || checkIf("colorNum", arg)) && (c.fg == "" || c.bg == "") {
			if c.fg == "" {
				c.fg = arg
			} else if c.bg == "" {
				c.bg = arg
			}
		} else if checkIf("bold", arg) && c.bold == "" {
			c.bold = "bold"
		} else if checkIf("light", arg) && c.style == "normal" {
			c.style = "light"
		} else if !msgChanged {
			c.msg = arg
			msgChanged = true
		}
	}
	colorPrint(c)
}
