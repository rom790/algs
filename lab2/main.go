package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// инициализация структур ввода-вывода
	reader := bufio.NewReaderSize(os.Stdin, 5_000_100)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	visFlag := flag.Bool("vis", false, "Использовать визуализацию")
	numFlag := flag.Int("n", 1, "Номер используемой функции")

	flag.Parse()

	if *numFlag != 1 && *numFlag != 2 {
		fmt.Fprintln(writer, "Неверный аргумент флага запуска")
		fmt.Fprintln(writer, "Для поиска подстроки используйте значение <<1>>")
		fmt.Fprintln(writer, "Для проверки на циклический сдвиг используйте значение <<2>>")
		return
	}
	var s, t string

	if *numFlag == 1 {
		fmt.Fscan(reader, &t)
		fmt.Fscan(reader, &s)
		positions := KMPSearchBasic(s, t, *visFlag)
		if len(positions) == 0 {

			fmt.Fprintln(writer, "-1")
			return
		}

		for i := 0; i < len(positions); i++ {
			if i == len(positions)-1 {
				fmt.Fprintf(writer, "%d", positions[i])
				break
			}
			fmt.Fprintf(writer, "%d,", positions[i])
		}
		return
		// fmt.Fprintln(writer, positions)
	} else {
		fmt.Fscan(reader, &s)
		fmt.Fscan(reader, &t)
		if len(s) != len(t) {
			fmt.Fprintln(writer, -1)
			return
		}
		position := KMPCyclicSearch(s+s, t, *visFlag)
		fmt.Fprintln(writer, position)

	}

}

// Подсчитывает значения префикс функции
func prefFunc(s string, visFlag bool) []int {
	vals := make([]int, len(s))

	if visFlag {
		fmt.Println("Расчёт префикс функции:")
	}

	for i := 1; i < len(s); i++ {
		// текущее количество совпавших символов
		curr := vals[i-1]

		if visFlag {
			// вывод при визуализации + задержка
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("Шаг %d:\n", i)
		}

		for s[i] != s[curr] {
			//	уменьшаем curr, пока текущий символ не станет равен символу из префикса или пока curr не станет 0
			if curr < 1 {
				break
			}
			curr = vals[curr-1]
		}

		if s[i] == s[curr] {
			// если символы совпали, увеличиваем curr
			if visFlag {
				// вывод при визуализации
				fmt.Printf("\t\033[90m%s\033[42m%s\033[0m\033[32m%s\033[90m%s\033[0m\n", s[0:i-curr], s[i-curr:i], string(s[i]), s[i+1:])
				fmt.Printf("\t\033[42m\033[90m%s\033[0m\033[32m%s\033[90m%s\033[0m\n", s[0:curr], string(s[curr]), s[curr+1:])
			}
			vals[i] = curr + 1
			if visFlag {
				// вывод при визуализации
				fmt.Printf("\t\033[90m%s\033[0m %d\n", convToStr(vals[:i]), vals[i])
			}
		} else {
			if visFlag {
				// вывод при визуализации
				fmt.Printf("\t\033[90m%s\033[0m\033[91m%s\033[90m%s\033[0m\n", s[0:i], string(s[i]), s[i+1:])
				fmt.Printf("\t\033[42m\033[90m%s\033[0m\033[91m%s\033[90m%s\033[0m\n", s[0:curr], string(s[curr]), s[curr+1:])
				fmt.Printf("\t\033[90m%s\033[0m %d\n", convToStr(vals[:i]), vals[i])
			}
		}
	}

	return vals
}

// Проверяет, является ли одна строка циклическим сдвигом другой
func KMPCyclicSearch(text, pattern string, visFlag bool) int {

	// если строки пустые, то они являются циклическим сдвигом друг друга
	if len(text) == 0 && len(pattern) == 0 {
		return 0
	}
	// значение префикс функции для проверяемой строки
	prefArr := prefFunc(pattern, visFlag)
	n := len(text)
	m := len(pattern)
	// текущее количество совпавших символов
	curr := 0

	if visFlag {
		//
		fmt.Println("Поиск вхождений:")
	}

	for i := 0; i < n; i++ {

		if visFlag {
			//  вывод при визуализации
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("Шаг %d:\n", i+1)
		}

		for pattern[curr] != text[i] && curr > 0 {
			if visFlag {
				fmt.Printf("\t\033[90m%s\033[0m\033[91m%s\033[90m%s\033[0m\n", text[0:i], string(text[i]), text[i+1:])
				fmt.Printf("\t\033[90m%s\033[0m\033[91m%s\033[90m%s\033[0m\n", pattern[0:curr], string(pattern[curr]), pattern[curr+1:])
			}
			// уменьшаем curr, пока оно больше нуля или пока выбранный символ из префикса не станет равен текущему
			curr = prefArr[curr-1]
		}

		if pattern[curr] == text[i] {
			// Если символы совпал - увеличиваем curr
			if visFlag {
				fmt.Printf("\t\033[90m%s\033[42m\033[90m%s\033[0m\033[32m%s\033[90m%s\033[0m\n", text[0:i-curr], text[i-curr:i], string(text[i]), text[i+1:])
				fmt.Printf("\t\033[42m\033[90m%s\033[0m\033[32m%s\033[90m%s\033[0m\n", pattern[0:curr], string(pattern[curr]), pattern[curr+1:])
			}
			curr++
		} else {
			if visFlag {
				fmt.Printf("\t\033[90m%s\033[0m\033[91m%s\033[90m%s\033[0m\n", text[0:i], string(text[i]), text[i+1:])
				fmt.Printf("\t\033[90m%s\033[0m\033[91m%s\033[90m%s\033[0m\n", pattern[0:curr], string(pattern[curr]), pattern[curr+1:])
			}
		}
		if curr == m {
			// Если получилось так, что количество совпавших символов равно длинее проверяемой строки, значит она является циклическим сдвигом
			if visFlag {
				fmt.Println("Подстрока найдена:")
				fmt.Printf("\t\033[90m%s\033[0m\033[43m%s\033[0m\033[90m%s\033[0m\n", text[0:i-m+1], text[i-m+1:i+1], text[i+1:])
				fmt.Printf("\t\033[43m%s\033[0m\n", pattern)
			}
			return i - m + 1
		}

	}
	return -1
}

// Выполняет поиск всех вхождений подстроки в встроку
func KMPSearchBasic(text, pattern string, visFlag bool) []int {
	n := len(text)
	m := len(pattern)
	// подсчёт префикс функции
	prefArr := prefFunc(pattern, visFlag)
	curr := 0
	// массив с индексами начала вхождений
	matches := []int{}

	if visFlag {
		fmt.Println("Поиск вхождений:")
	}
	for i := 0; i < n; i++ {
		if visFlag {
			time.Sleep(1000 * time.Millisecond)
			fmt.Printf("Шаг %d:\n", i+1)
		}

		for curr > 0 && pattern[curr] != text[i] {
			if visFlag {
				fmt.Printf("\t\033[90m%s\033[0m\033[91m%s\033[90m%s\033[0m\n", text[0:i], string(text[i]), text[i+1:])
				fmt.Printf("\t\033[90m%s\033[0m\033[91m%s\033[90m%s\033[0m\n", pattern[0:curr], string(pattern[curr]), pattern[curr+1:])
			}
			curr = prefArr[curr-1]
		}
		if pattern[curr] == text[i] {
			if visFlag {
				fmt.Printf("\t\033[90m%s\033[42m\033[90m%s\033[0m\033[32m%s\033[90m%s\033[0m\n", text[0:i-curr], text[i-curr:i], string(text[i]), text[i+1:])
				fmt.Printf("\t\033[42m\033[90m%s\033[0m\033[32m%s\033[90m%s\033[0m\n", pattern[0:curr], string(pattern[curr]), pattern[curr+1:])
			}
			curr++
		} else {
			if visFlag {
				fmt.Printf("\t\033[90m%s\033[0m\033[91m%s\033[90m%s\033[0m\n", text[0:i], string(text[i]), text[i+1:])
				fmt.Printf("\t\033[90m%s\033[0m\033[91m%s\033[90m%s\033[0m\n", pattern[0:curr], string(pattern[curr]), pattern[curr+1:])
			}
		}
		if curr == m {
			// Если количество совпавших символов равно длине подстроки, значит она полность входит в строку, сохраняем индекс начала вхождения
			if visFlag {
				fmt.Println("Подстрока найдена:")
				fmt.Printf("\t\033[90m%s\033[0m\033[43m%s\033[0m\033[90m%s\033[0m\n", text[0:i-m+1], text[i-m+1:i+1], text[i+1:])
				fmt.Printf("\t\033[43m%s\033[0m\n", pattern)
			}
			matches = append(matches, i-m+1)
			curr = prefArr[curr-1]
		}
	}
	return matches
}

// Представляет массив в строки
func convToStr(nums []int) string {

	// Создаем срез строк для хранения строковых представлений чисел
	strNumbers := make([]string, len(nums))

	// Преобразуем числа в строки
	for i, num := range nums {
		strNumbers[i] = strconv.Itoa(num)
	}

	// Объединяем строки с разделителем " "
	result := strings.Join(strNumbers, " ")

	// Выводим результат
	return result
}

/*
package main

import (
	"bufio"
	"fmt"
	"os"
	// "strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 5_000_100)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var s, t string
	fmt.Fscan(reader, &t)
	fmt.Fscan(reader, &s)

	// var sb strings.Builder

	trimSpecSym(&t)
	trimSpecSym(&s)

	positions := kmpSearch(s, t)

	if len(positions) == 0 {
		fmt.Fprintln(writer, -1)
	}
	for i := 0; i < len(positions); i++ {
		if i == len(positions)-1 {
			fmt.Fprintf(writer, "%d", positions[i])
			break
		}
		fmt.Fprintf(writer, "%d,", positions[i])
	}

}

func prefFunc(s string) []int {
	vals := make([]int, len(s))

	for i := 1; i < len(s); i++ {
		curr := vals[i-1]

		for s[i] != s[curr] {
			if curr < 1 {
				break
			}
			curr = vals[curr-1]
		}

		if s[i] == s[curr] {
			vals[i] = curr + 1
		}
	}

	return vals
}

func trimSpecSym(s *string) {
	if len(*s) < 2 {
		return
	}

	for i := 0; i < 2; i++ {
		if (*s)[len((*s))-1] == '\n' || (*s)[len((*s))-1] == '\r' || (*s)[len((*s))-1] == ' ' || (*s)[len((*s))-1] == '\t' {
			(*s) = (*s)[:len((*s))-1]
		}
	}

}

func kmpSearch(text, pattern string) []int {
	n := len(text)
	m := len(pattern)
	prefArr := prefFunc(pattern)
	q := 0
	matches := []int{}

	for i := 0; i < n; i++ {
		for q > 0 && pattern[q] != text[i] {
			q = prefArr[q-1]
		}
		if pattern[q] == text[i] {
			q++
		}
		if q == m {
			matches = append(matches, i-m+1)
			q = prefArr[q-1]
		}
	}
	return matches
}*/
