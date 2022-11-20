/*
[RU] Общий модуль с разлиными полезными функциями, которые помогают организовать
структуру и выполнение всех примеров на Go с сайта https://allineed.ru
[EN] Common module with different useful functions that help to organize the
structure and execution for all the Go samples from https://allineed.ru site

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
package ain_common

import "fmt"

// [RU] Глобально задаёт язык вывода в консоль для всех примеров;
// [EN] Globally specifies the language of console output for all the examples
var Language string

// [RU] Функция выводит название примера из статьи сайта allineed.ru на консоль;
// [EN] This func prints to the console the name of the example from allineed.ru site article
func PrintExampleTitle(ru_title string, en_title string, url string) {

	ru_title = "\"" + ru_title + "\""
	en_title = "\"" + en_title + "\""

	fmt.Println("---------------------------------------------------------------------")
	if Language == "ru" {
		fmt.Println(" >>> [RU] Запуск примера из статьи с сайта allineed.ru:", ru_title)
		fmt.Println(" >>> [RU] Ссылка на статью:", url)
	} else if Language == "en" {
		fmt.Println(" >>> [EN] Running the sample from the allineed.ru site article:", en_title)
		fmt.Println(" >>> [EN] Link to the article:", url)
	}
	fmt.Println("---------------------------------------------------------------------")
}

// [RU] Возвращает true если текущий язык примеров - русский;
// [EN] Returns true if the current language for the examples is Russian
func IsRuLocale() bool {
	return Language == "ru"
}

// [RU] Возвращает true если текущий язык примеров - английский;
// [EN] Returns true if the current language for the examples is English
func IsEnLocale() bool {
	return Language == "en"
}

// [RU] Структура содержит массив строк в определённой локали;
// [EN] This struct contains an array of strings in a specific locale
type LocalizedStrings struct {
	args []string
}

// [RU] Создаёт новую структуру LocalizedStrings и возвращает на неё указатель;
// [EN] Creates new struct LocalizedStrings and returns the pointer to it
func NewLocalizedStrings(arguments ...string) *LocalizedStrings {
	li := LocalizedStrings{args: arguments}
	return &li
}

// [RU] Вызывает fmt.Println для одного из переданных аргументов, в зависимости от установленной локали;
// [EN] Calls fmt.Println for one of the passed arguments, depending on the predefined locale
func PrintlnLocalized(ru_message string, en_message string) {
	if Language == "ru" {
		fmt.Println(ru_message)
	} else if Language == "en" {
		fmt.Println(en_message)
	}
}

// [RU] Создаёт и возвращает одну строку из переданного массива строк;
// [EN] Creates and returns a single string from the given input array of strings
func GetStringFromArgs(args []string) string {
	result := ""
	for i := 0; i < len(args); i++ {
		result += args[i]
	}
	return result
}

// [RU] Вернёт одно из поданных входных значений, в зависимости от текущего языка примеров;
// [EN] Returns one of the input values depending on the current examples locale
func GetLocalString(ru_message string, en_message string) string {
	if IsRuLocale() {
		return ru_message
	} else {
		return en_message
	}
}

// [RU] Вызывает fmt.Println для одного из переданных аргументов (предварительно получив из структуры одну строку), в зависимости от установленной локали;
// [EN] Calls fmt.Println for one of the passed arguments (and gets from the struct one single string before this), depending on the predefined locale
func PrintlnLocalizedStrings(ru *LocalizedStrings, en *LocalizedStrings) {
	PrintlnLocalized(GetStringFromArgs(ru.args), GetStringFromArgs(en.args))
}

// [RU] Работа аналична PrintlnLocalizedStrings, но предваряет вывод символом табуляции;
// [EN] Works similarly as PrintlnLocalizedStrings, but prints tab symbol before printing the main output
func PrintlnLocalizedStringsWithTab(ru *LocalizedStrings, en *LocalizedStrings) {
	PrintlnLocalized("\t"+GetStringFromArgs(ru.args), "\t"+GetStringFromArgs(en.args))
}

// [RU] Работа аналична PrintlnLocalized, но предваряет вывод символом табуляции;
// [EN] Works similarly as PrintlnLocalized, but prints tab symbol before printing the main output
func PrintlnLocalizedWithTab(ru_message string, en_message string) {
	PrintlnLocalized("\t"+ru_message, "\t"+en_message)
}
