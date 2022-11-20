/*
[RU] Модуль содержит пример кода на языке Go из статьи
https://allineed.ru/development/go-development/60-go-using-if-statement
[EN] Module contains the sample Go code from the article
https://allineed.ru/development/go-development/60-go-using-if-statement

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

package ain_if_statements

import (
	"strconv"

	"allineed.ru/x/examples/ain_common"
)

func RunSample() {
	ain_common.PrintExampleTitle(
		"Пример работы с оператором 'if'",
		"Example of using 'if' statement", "https://allineed.ru/development/go-development/60-go-using-if-statement")

	ain_common.PrintlnLocalized(
		"1) Демонстрация константных выражений в операторе 'if':",
		"1) Demo of constant expressions in 'if' statement:")

	if 10%2 == 0 {
		ain_common.PrintlnLocalizedWithTab("10 - чётное число", "10 is even")
	} else {
		ain_common.PrintlnLocalizedWithTab("10 - нечётное число", "10 is odd")
	}

	ain_common.PrintlnLocalized(
		"2) Демонстрация инициализации переменной с последующей проверкой значения в операторе 'if':",
		"2) Demo of variable initialization with further check of its value in the 'if' statement:")

	if x := 1; x > 10 {
		ain_common.PrintlnLocalizedStringsWithTab(
			ain_common.NewLocalizedStrings(strconv.Itoa(x), " больше 10"),
			ain_common.NewLocalizedStrings(strconv.Itoa(x), " less than 10"))
	} else if x < 0 {
		ain_common.PrintlnLocalizedStringsWithTab(
			ain_common.NewLocalizedStrings(strconv.Itoa(x), " меньше 0"),
			ain_common.NewLocalizedStrings(strconv.Itoa(x), " less than 0"))
	} else {
		ain_common.PrintlnLocalizedStringsWithTab(
			ain_common.NewLocalizedStrings(strconv.Itoa(x), " больше или равен 0 и меньше или равен 10"),
			ain_common.NewLocalizedStrings(strconv.Itoa(x), " is greater or equal 0 and less or equal 10"))
	}

	ain_common.PrintlnLocalized(
		"3) Демонстрация присваивания значений нескольким переменным одновременно и затем использование переменных в операторе 'if'::",
		"3) Demo of assigning values for several variables and further usage of these variables in the 'if' statement:")

	y, z, t := 10, 20, 99
	if (y > 0 && z > 0) || t < -20 {
		ain_common.PrintlnLocalizedStringsWithTab(
			ain_common.NewLocalizedStrings(strconv.Itoa(y), " и ", strconv.Itoa(z), " оба больше 0 или ", strconv.Itoa(t), " < -20"),
			ain_common.NewLocalizedStrings(strconv.Itoa(y), " and ", strconv.Itoa(z), " are both greater than 0 or ", strconv.Itoa(t), " < -20"))
	} else {
		ain_common.PrintlnLocalizedStringsWithTab(
			ain_common.NewLocalizedStrings("Как минимум ", strconv.Itoa(y), " или ", strconv.Itoa(z), " не больше 0, либо ", strconv.Itoa(t), " > -20"),
			ain_common.NewLocalizedStrings("At least one of ", strconv.Itoa(y), " or ", strconv.Itoa(z), " are not greater than 0, or ", strconv.Itoa(t), " > -20"))
	}
}
