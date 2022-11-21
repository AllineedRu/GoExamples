/*
[RU] Модуль содержит пример кода на языке Go из статьи
https://allineed.ru/development/go-development/61-go-using-for-loop
[EN] Module contains the sample Go code from the article
https://allineed.ru/development/go-development/61-go-using-for-loop

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

package ain_for_loop

import (
	"strconv"

	"allineed.ru/x/examples/ain_common"
)

func RunSample() {
	ain_common.PrintExampleTitle(
		"Цикл for в языке Go",
		"for loop in Go",
		"https://allineed.ru/development/go-development/61-go-using-for-loop")

	ain_common.PrintlnLocalized(
		"1) Демонстрация \"традиционного\" цикла 'for':",
		"1) Demo of \"traditional\" 'for' cycle:")

	for i := 1; i <= 20; i++ {
		ain_common.PrintlnLocalizedStringsWithTab(
			ain_common.NewLocalizedStrings("значение счётчика: i = ", strconv.Itoa(i)),
			ain_common.NewLocalizedStrings("counter value: i = ", strconv.Itoa(i)))
	}

	ain_common.PrintlnLocalized(
		"2) Демонстрация цикла 'for', аналогичного циклу 'while':",
		"2) Demo of 'for' cycle that is similar to 'while' loop:")

	t := 1
	for t <= 5 {
		ain_common.PrintlnLocalizedStringsWithTab(
			ain_common.NewLocalizedStrings("значение счётчика: t = ", strconv.Itoa(t)),
			ain_common.NewLocalizedStrings("counter value: t = ", strconv.Itoa(t)))
		t++
	}

	ain_common.PrintlnLocalized(
		"3) Демонстрация цикла 'for' с использованием 'break' и 'continue':",
		"3) Demo of 'for' cycle with 'break' and 'continue':")

	for x, y := 0, 0; x <= 10 && y <= 30; x, y = x+1, y+2 {
		if x < 3 || y < 3 {
			continue
		}
		if x == 9 {
			break
		}
		ain_common.PrintlnLocalizedStringsWithTab(
			ain_common.NewLocalizedStrings("значение x = ", strconv.Itoa(x), "; значение y = ", strconv.Itoa(y)),
			ain_common.NewLocalizedStrings("value of x is: ", strconv.Itoa(x), "; value of y is: ", strconv.Itoa(y)))
	}

	ain_common.PrintlnLocalized(
		"4) Демонстрация цикла 'for' с условием выхода из цикла внутри самого цикла:",
		"4) Demo of 'for' cycle with a condition of exiting from loop inside the cycle itself:")

	v := 10
	for {
		v++
		if v > 20 {
			break
		}

		ain_common.PrintlnLocalizedStringsWithTab(
			ain_common.NewLocalizedStrings("значение v = ", strconv.Itoa(v)),
			ain_common.NewLocalizedStrings("value of v is: ", strconv.Itoa(v)))
	}
}
