/*
[RU] Модуль содержит пример кода на языке Go из статьи
https://allineed.ru/development/go-development/57-go-create-file-and-write-data

[EN] Module contains the sample Go code from the article
https://allineed.ru/development/go-development/57-go-create-file-and-write-data

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

package ain_writing_to_file

import (
	"fmt"
	"os"

	"allineed.ru/x/examples/ain_common"
)

func RunSample(path_to_file string) {
	ain_common.PrintExampleTitle(
		"Создаём файл на языке Go и пишем в него данные",
		"Creating file in Go language and writing data to it", "https://allineed.ru/development/go-development/57-go-create-file-and-write-data")

	// [RU] указываем путь к новому файлу, который мы хотим создать
	// [EN] specify the path to the new file that we are willing to create
	file, err := os.Create(path_to_file)

	// [RU] если при создании файла произошла ошибка, выдать детали по ней
	// [EN] in case error happens during the file creation, print the error details
	if err != nil {
		panic(err)
	}

	// [RU] записываем нашу строку в созданный файл
	// [EN] write our string into the newly created file
	written, err := file.WriteString(
		ain_common.GetLocalString("Это моя первая строка, которую я запишу в файл с помощью языка Go!",
			"This is my first string that I'm writing to the file using the Go language!"))

	// [RU] подготовим сообщение с количеством записанных в файл байт
	// [EN] prepare the message with the number of bytes actually written to the file
	message := ain_common.GetLocalString("Записано %#v байт в новый файл!\n", "Written %#v bytes into the new file!\n")

	// [RU] если при записи в файл произошла ошибка, выдать детали по ней
	// [EN] in case error happens when writing to the file, print the error details
	if err != nil {
		panic(err)
	}

	// [RU] выведем информацию о количестве байт, записанных в файл
	// [EN] print the info about the number of bytes written to the file
	fmt.Printf(message, written)
}
