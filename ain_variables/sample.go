/*
[RU] Модуль содержит пример кода на языке Go из статьи
https://allineed.ru/development/go-development/68-go-variables
[EN] Module contains the sample Go code from the article
https://allineed.ru/development/go-development/68-go-variables

This program is free software: you can redistribute it and/or modify it under
the terms of the GNU General Public License as published by the Free Software
Foundation, either version 3 of the License, or (at your option) any later
version.
This program is distributed in the hope that it will be useful, but WITHOUT
ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS
FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
You should have received a copy of the GNU General Public License along with
this program. If not, see <http://www.gnu.org/licenses/>.
*/

package ain_variables

import (
	"fmt"

	"allineed.ru/x/examples/ain_common"
)

func RunSample() {
	ain_common.PrintExampleTitle(
		"Переменные в Go",
		"Variables in Go",
		"https://allineed.ru/development/go-development/68-go-variables")

	var s = "some value"
	fmt.Println("s =", s)

	var s0 string = "s0 value"
	fmt.Println("s0 =", s0)

	var f1 = 1.5
	fmt.Println("f1 =", f1)

	var f2 float64 = 456.7
	fmt.Println("f2 =", f2)

	var b1 = true
	fmt.Println("b1 =", b1)

	var b2 bool = false
	fmt.Println("b2 =", b2)

	var def_b bool
	var def_f float64
	var def_i int64
	var def_s string
	fmt.Println("def_b =", def_b)
	fmt.Println("def_f =", def_f)
	fmt.Println("def_i =", def_i)
	fmt.Println("def_s =", def_s)

	s1 := "a b c"
	fmt.Println("s1 =", s1)

	s2, s3 := "2", "3"
	fmt.Println("s2 =", s2)
	fmt.Println("s3 =", s3)

	var s4, s5 string = "4", "5"
	fmt.Println("s4 =", s4)
	fmt.Println("s5 =", s5)
}
