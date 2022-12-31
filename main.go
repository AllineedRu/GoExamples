/*
[RU] Основной модуль для запуска всех примеров кода на языке Go для сайта https://allineed.ru/
[EN] The main module for running all Go examples for https://allineed.ru/ site

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

package main

import (
	"allineed.ru/x/examples/ain_common"
	"allineed.ru/x/examples/ain_for_loop"
	"allineed.ru/x/examples/ain_if_statements"
	"allineed.ru/x/examples/ain_variables"
	"allineed.ru/x/examples/ain_writing_to_file"
)

func main() {
	// [RU] Установить язык для всех примеров (ru = Русский; en = Английский)
	// [EN] Set the language for all the examples (ru = Russian; en = English)
	ain_common.Language = "ru"

	// [RU] Запуск примера из статьи "Используем условный оператор if в Go"
	// [RU] Ссылка на статью: https://allineed.ru/development/go-development/60-go-using-if-statement
	// [EN] Running example from the article "Using conditional operator if in Go"
	// [EN] Link to the article: https://allineed.ru/development/go-development/60-go-using-if-statement
	ain_if_statements.RunSample()

	// [RU] Запуск примера из статьи "Создаём файл на языке Go и пишем в него данные"
	// [RU] Ссылка на статью: https://allineed.ru/development/go-development/57-go-create-file-and-write-data
	// [RU] Во входном аргументе для функции RunSample() укажите путь, где будет создан файл 'mygofile.txt'
	// [EN] Running example from the article "Creating file in Go language and writing data to it"
	// [EN] Link to the article: https://allineed.ru/development/go-development/57-go-create-file-and-write-data
	// [EN] In the input argument for the RunSample() function specify the full path where to 'mygofile.txt' should be created
	ain_writing_to_file.RunSample("C:\\Users\\1\\mygofile.txt")

	// [RU] Запуск примера из статьи "Цикл for в языке Go"
	// [RU] Ссылка на статью: https://allineed.ru/development/go-development/61-go-using-for-loop
	// [EN] Running example from the article "for loop in Go"
	// [EN] Link to the article: https://allineed.ru/development/go-development/61-go-using-for-loop
	ain_for_loop.RunSample()

	// [RU] Запуск примера из статьи "Переменные в языке Go"
	// [RU] Ссылка на статью: https://allineed.ru/development/go-development/68-go-variables
	// [EN] Running example from the article "Variables in Go"
	// [EN] Link to the article: https://allineed.ru/development/go-development/68-go-variables
	ain_variables.RunSample()
}
