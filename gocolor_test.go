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

// This is for test only.

package gocolor

import (
	"testing"
)

var colorTestName = []string{"red", "green", "yellow", "blue", "magenta", "cyan", "white", "105", "227", "92", "200"}

func Test256Color(t *testing.T) {
	Color("Test 256 color: ", "white", "Bold", "light")
	for i := 0; i < cap(colorTestName); i++ {
		Color("Test "+colorTestName[i]+" color: ", "white", "light")
		Color("Light", " "+colorTestName[i]+" ", colorTestName[i], "sameline", "underline", "bold")
		Color("white", " "+colorTestName[i]+" ", colorTestName[i], "bold")
	}
}

func TestNon256Color(t *testing.T) {
	Color("\nTest non-256 color: ", "white", "Bold", "-256", "light")
	for i := 0; i < cap(colorTestName)-4; i++ {
		Color("Test "+colorTestName[i]+" color: ", "white", "light", "-256")
		Color(" "+colorTestName[i]+" ", colorTestName[i], "sameline", "-256", "underline", "bold")
		Color("white", " "+colorTestName[i]+" ", colorTestName[i], "bold", "-256")
	}
}
