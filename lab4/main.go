package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// стоимость операций
var (
	del = 1
	ins = 1
	rep = 1
)

// функция управления процессом работы программы
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	taskFlag := flag.Int("n", 3, "Номер задания")
	visFlag := flag.Bool("vis", false, "Исопльзовать визуализацию")
	stepFlag := flag.Bool("steps", false, "Использовать пошаговую визуализацию")
	flag.Parse()

	if *taskFlag == 1 || *taskFlag == 2 {
		fmt.Fscan(reader, &rep, &ins, &del)
	}

	var firstStr, secondStr string
	fmt.Fscan(reader, &firstStr, &secondStr)

	// Инициализация матрицы
	dp := make([][]int, len(firstStr)+1)
	for i := range dp {
		dp[i] = make([]int, len(secondStr)+1)
	}

	// Заполнение матрицы
	for i := 0; i <= len(firstStr); i++ {
		for j := 0; j <= len(secondStr); j++ {

			if i == 0 && j == 0 {
				dp[i][j] = 0
			} else if j == 0 {
				dp[i][j] = i * del
			} else if i == 0 {
				dp[i][j] = j * ins
			} else if firstStr[i-1] == secondStr[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1]+ins, dp[i-1][j]+del, dp[i-1][j-1]+rep)
			}
			if *stepFlag {
				time.Sleep(500 * time.Millisecond)
				if i > 0 && j > 0 {
					current := [][2]int{{i - 1, j - 1}, {i - 1, j}, {i, j - 1}, {i, j}}
					printDPTable(dp, firstStr, secondStr, nil, current)
				} else {
					current := [][2]int{{i, j}}
					printDPTable(dp, firstStr, secondStr, nil, current)
				}
			}
		}
	}

	if *visFlag {
		printDPTable(dp, firstStr, secondStr, getPath(dp, firstStr, secondStr, false), nil)
		fmt.Fprintln(writer, "Вычисленное расстояние:")
	}

	if *taskFlag == 2 {
		getPath(dp, firstStr, secondStr, true)
		fmt.Fprintln(writer, firstStr)
		fmt.Fprintln(writer, secondStr)
	} else {
		fmt.Fprintln(writer, dp[len(firstStr)][len(secondStr)])
	}

	for *taskFlag == 4 {
		fmt.Fprintln(writer, "\nВыберите операцию")
		fmt.Fprintln(writer, "1. Расширить первую строку")
		fmt.Fprintln(writer, "2. Расширить вторую строку")
		fmt.Fprintln(writer, "3. Завершить")
		fmt.Fprint(writer, "Вы выбрали: ")
		writer.Flush()

		var choice int
		fmt.Fscan(reader, &choice)

		switch choice {
		case 1:
			fmt.Fprint(writer, "Введите продолжение строки:")
			writer.Flush()
			var extension string
			fmt.Fscan(reader, &extension)
			firstStr, dp = extendString(firstStr, secondStr, extension, dp, true)
		case 2:
			fmt.Fprint(writer, "Введите продолжение строки: ")
			writer.Flush()
			var extension string
			fmt.Fscan(reader, &extension)
			secondStr, dp = extendString(firstStr, secondStr, extension, dp, false)
		case 3:
			return
		default:
			fmt.Fprintln(writer, "Некорректная команда")
			return
		}

		fmt.Fprintln(writer, "Новое расстояние:", dp[len(firstStr)][len(secondStr)])
		fmt.Fprintln(writer, "Первая строка:", firstStr)
		fmt.Fprintln(writer, "Вторая строка:", secondStr)
		if *visFlag {
			printDPTable(dp, firstStr, secondStr, getPath(dp, firstStr, secondStr, false), nil)
		}
	}

}

// функция расширения выбранной строки набором символов, вводимым пользователем
func extendString(str1, str2, extension string, dp [][]int, extendFirst bool) (string, [][]int) {
	if extendFirst {
		// Расширяем первую строку
		oldLen := len(str1)
		newStr := str1 + extension
		newDp := make([][]int, len(newStr)+1)

		// Копируем старые значения
		for i := 0; i <= oldLen; i++ {
			newDp[i] = make([]int, len(str2)+1)
			copy(newDp[i], dp[i])
		}

		// Добавляем новые строки в матрицу
		for i := oldLen + 1; i <= len(newStr); i++ {
			newDp[i] = make([]int, len(str2)+1)
			newDp[i][0] = i * del // Заполняем первый столбец

			for j := 1; j <= len(str2); j++ {
				if newStr[i-1] == str2[j-1] {
					newDp[i][j] = newDp[i-1][j-1]
				} else {
					newDp[i][j] = min(
						newDp[i-1][j]+del,
						newDp[i][j-1]+ins,
						newDp[i-1][j-1]+rep,
					)
				}
			}
		}

		return newStr, newDp
	} else {
		// Расширяем вторую строку
		oldLen := len(str2)
		newStr := str2 + extension
		newDp := make([][]int, len(str1)+1)

		// Копируем старые значения и расширяем столбцы
		for i := range newDp {
			newDp[i] = make([]int, len(newStr)+1)
			copy(newDp[i], dp[i][:oldLen+1]) // Копируем старые значения

			// Заполняем новые столбцы
			for j := oldLen + 1; j <= len(newStr); j++ {
				if i == 0 {
					newDp[i][j] = j * ins // Заполняем первую строку
				} else {
					if str1[i-1] == newStr[j-1] {
						newDp[i][j] = newDp[i-1][j-1]
					} else {
						newDp[i][j] = min(
							newDp[i-1][j]+del,
							newDp[i][j-1]+ins,
							newDp[i-1][j-1]+rep,
						)
					}
				}
			}
		}

		return newStr, newDp
	}
}

// функция вывода таблицы
func printDPTable(dp [][]int, s1, s2 string, path [][2]int, selectCells [][2]int) {
	pathMap := make(map[[2]int]bool)
	selectCellsMap := make(map[[2]int]bool)
	var currentCell [2]int

	if path != nil {

		// Создаем карту для быстрой проверки принадлежности к пути
		for _, pair := range path {
			pathMap[[2]int{pair[0], pair[1]}] = true
		}
	}
	if selectCells != nil {
		for i := 0; i < len(selectCells)-1; i++ {
			selectCellsMap[[2]int{selectCells[i][0], selectCells[i][1]}] = true
		}

		currentCell = [2]int{selectCells[len(selectCells)-1][0], selectCells[len(selectCells)-1][1]}
	}
	// Определяем максимальную ширину чисел
	maxValWidth := 1
	for _, row := range dp {
		for _, val := range row {
			if w := len(fmt.Sprint(val)); w > maxValWidth {
				maxValWidth = w
			}
		}
	}
	cellWidth := maxValWidth + 2

	// ANSI коды для цветов
	const (
		reset   = "\033[0m"
		bold    = "\033[1m"
		redBg   = "\033[41m"
		whiteFg = "\033[37m"
		blueBg  = "\033[44m"
	)

	// Верхние заголовки
	fmt.Print("          ")
	for _, c := range s2 {
		fmt.Printf(" %*s  ", cellWidth, string(c))
	}
	fmt.Println()

	fmt.Print("          ")
	for j := range s2 {
		fmt.Printf(" %*d  ", cellWidth, j+1)
	}
	fmt.Println()

	// Верхняя граница
	fmt.Print("   +")
	for j := 0; j <= len(s2); j++ {
		fmt.Printf("%s+", strings.Repeat("-", cellWidth+2))
	}
	fmt.Println()

	// Тело таблицы
	for i := range dp {
		// Левая колонка
		if i == 0 {
			fmt.Print("   |")
		} else {
			fmt.Printf("%c%2d|", s1[i-1], i)
		}

		// Ячейки строки
		for j := range dp[i] {
			if path != nil {
				if pathMap[[2]int{i, j}] {
					fmt.Printf("  %s%s%*d%s  |", bold+whiteFg+redBg, "", maxValWidth, dp[i][j], reset)
				} else {
					fmt.Printf("  %*d  |", maxValWidth, dp[i][j])
				}
			} else if selectCells != nil {
				// fmt.Println(selectCellsMap)
				if selectCellsMap[[2]int{i, j}] {
					fmt.Printf("%s%s  %*d  %s|", bold+whiteFg+redBg, "", maxValWidth, dp[i][j], reset)
				} else if i == currentCell[0] && j == currentCell[1] {
					fmt.Printf("%s%s  %*d  %s|", bold+whiteFg+blueBg, "", maxValWidth, dp[i][j], reset)
				} else {
					fmt.Printf("  %*d  |", maxValWidth, dp[i][j])
				}
			} else {
				fmt.Printf("  %*d  |", maxValWidth, dp[i][j])
			}
		}
		fmt.Println()

		// Разделители
		if i < len(dp)-1 {
			fmt.Print("   |")
			for j := 0; j <= len(s2); j++ {
				fmt.Printf("%s+", strings.Repeat("-", cellWidth+2))
			}
			fmt.Println()
		}
	}

	// Нижняя граница
	fmt.Print("   +")
	for j := 0; j <= len(s2); j++ {
		fmt.Printf("%s+", strings.Repeat("-", cellWidth+2))
	}
	fmt.Println()
}

// функция, осуществляющая обратный ход по таблицы для вычисления последовательности операций
func getPath(dp [][]int, firstStr, secondStr string, show bool) [][2]int {
	operations := make([]byte, 0)
	path := make([][2]int, 0)
	i, j := len(firstStr), len(secondStr)
	path = append(path, [2]int{i, j})

	for !(i == 0 && j == 0) {
		if i == 0 {
			for k := 0; k < j; k++ {
				operations = append(operations, 'I')
				path = append(path, [2]int{i, k})
			}
			break
		} else if j == 0 {
			for k := 0; k < i; k++ {
				operations = append(operations, 'D')
				path = append(path, [2]int{i, k})
			}
			break
		}

		sFlag := 0
		if firstStr[i-1] != secondStr[j-1] {
			sFlag = 1
		}

		minCurrent := min(
			dp[i-1][j]+1,
			dp[i][j-1]+1,
			dp[i-1][j-1]+sFlag*1,
		)

		if minCurrent == dp[i-1][j]+1 {
			path = append(path, [2]int{i, j})
			operations = append(operations, 'D')
			i--
		} else if minCurrent == dp[i][j-1]+1 {
			path = append(path, [2]int{i, j})
			operations = append(operations, 'I')
			j--
		} else if minCurrent == dp[i-1][j-1]+sFlag*1 {
			if sFlag == 1 {
				path = append(path, [2]int{i, j})
				operations = append(operations, 'R')
			} else {
				path = append(path, [2]int{i, j})
				operations = append(operations, 'M')
			}
			i--
			j--
		}
	}
	for k := 0; k < len(operations)/2; k++ {
		operations[k], operations[len(operations)-1-k] = operations[len(operations)-1-k], operations[k]
	}
	if show {
		fmt.Println(string(operations))
	}
	return path
}

// находит минимальное число из заданного набора
func min(nums ...int) int {
	minVal := nums[0]
	for _, num := range nums {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}
