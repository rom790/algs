package main

// // // // import (
// // // // 	// "bufio"

// // // // 	// "bufio"
// // // // 	"flag"
// // // // 	"fmt"
// // // // 	libs "lab1/libs"
// // // // 	// "os"
// // // // 	// "strconv"

// // // // 	"time"
// // // // 	// "time"
// // // // 	// "sort"
// // // // 	// "os"
// // // // 	// "strconv"
// // // // )

// // // // type State struct {
// // // // 	field   [][]int
// // // // 	squares [][]int
// // // // 	// lastPlacedInd int
// // // // 	lastSquare int
// // // // 	freeArea   int
// // // // }
// // // // type ExtraState struct {
// // // // 	field      [][]int
// // // // 	squares    [][]int
// // // // 	usedCoords [][]int
// // // // 	lastSquare int
// // // // 	freeArea   int
// // // // }

// // // // func main() {

// // // // 	// field := make([][]int, 7)
// // // // 	// for i := 0; i < len(field); i++ {
// // // // 	// 	field[i] = make([]int, 7)
// // // // 	// }

// // // // 	// placeSquare(field, 0, 0, 2)
// // // // 	// placeSquare(field, 3, 3, 4)
// // // // 	// placeSquare(field, 0, 2, 3)
// // // // 	// placeSquare(field, 0, 5, 2)

// // // // 	// // placeSquare(field, 2, 0, 2)
// // // // 	// for i := range len(field) {
// // // // 	// 	fmt.Println(field[i])
// // // // 	// }
// // // // 	// deleteSquare(field, 1, 2, 4)
// // // // 	// fmt.Println()
// // // // 	// for i := range len(field) {
// // // // 	// 	fmt.Println(field[i])
// // // // 	// }
// // // // 	// fmt.Println(findEmptySpace(field, 2))

// // // // 	// placeSquare(field, 0, 0, 2)

// // // // 	// fieldSize := 3
// // // // 	// squareSizes := []int{2, 1} // Квадраты 2x2 и 1x1
// // // // 	// placements := iterativeBacktrack(field, squareSizes, fieldSize)

// // // // 	// fmt.Printf("Всего размещений: %d\n", len(placements))
// // // // 	// fmt.Println("Примеры:")
// // // // 	// for i, p := range placements {
// // // // 	// 	if i >= 5 {
// // // // 	// 		break
// // // // 	// 	}
// // // // 	// 	fmt.Println(p)
// // // // 	// }

// // // // 	ioContr := libs.NewIoController()

// // // // 	field := make([][]int, 10)
// // // // 	for i := 0; i < len(field); i++ {
// // // // 		field[i] = make([]int, 10)
// // // // 	}

// // // // 	useIndividualization := flag.Bool("ind", false, "Использовать ввод для индивидуального задания")
// // // // 	useUnsafeMode := flag.Bool("unsafe", false, "Включить небезопасный режим")
// // // // 	flag.Parse()

// // // // 	var n int
// // // // 	var m = 0

// // // // 	var squares [][]int
// // // // 	if !(*useIndividualization) {
// // // // 		n = ioContr.ReadNum()
// // // // 	} else {
// // // // 		ioContr.WriteStrln("Введите число - размер стороны поля")
// // // // 		n = ioContr.ReadNum()

// // // // 		ioContr.WriteStrln("Введите число - количество используемых квадратов")
// // // // 		m = ioContr.ReadNum()
// // // // 		squares = make([][]int, m)

// // // // 		ioContr.WriteStrln("Введите квадраты в фомате <x y size>, разделённые символами переноса строки, если квадрат может быть размещён на любых координатах, то x и y ввести как -1")
// // // // 		for i := 0; i < m; i++ {
// // // // 			squares[i] = ioContr.ReadLineOfInts()
// // // // 		}
// // // // 	}
// // // // 	sqs, _ := backtracking(squares, n, *useUnsafeMode)
// // // // 	// for i := range len(sqs) {
// // // // 	// 	fmt.Println()
// // // // 	// }
// // // // 	fmt.Println(sqs, len(sqs))

// // // // 	// fmt.Println(checkPlace(field, 9, 9, 1))
// // // // 	// placeSquare(field, 9, 9, 1)
// // // // 	// for i := range 10 {
// // // // 	// 	fmt.Println(field[i])
// // // // 	// }
// // // // 	// fmt.Println(checkPlace(&field, 7, 7, 2))
// // // // 	// sort.Slice()
// // // // }

// // // // // Проверяет, можно ли разместить квадрат размером size*size на поле field, расположив верхний левый угол квадрата в координатах x, y поля
// // // // func checkPlace(field [][]int, x, y, squareSize int) bool {

// // // // 	if x+squareSize > len(field) || y+squareSize > len(field) {
// // // // 		// Если квадрат при наложении выходит за пределы поля, то его нельзя разместить
// // // // 		return false
// // // // 	}

// // // // 	for i := x; i < x+squareSize; i++ {
// // // // 		for j := y; j < y+squareSize; j++ {
// // // // 			// Если хотя бы одна клетка поля в рассматриваемой области занята, то размещение невозможно
// // // // 			if (field)[i][j] != 0 {
// // // // 				return false
// // // // 			}
// // // // 		}
// // // // 	}
// // // // 	return true
// // // // }

// // // // // Помечает клетки области поля, как "занятые", иметируя размещения квадрата на поле
// // // // func placeSquare(field [][]int, x, y, squareSize int) {
// // // // 	for i := x; i < x+squareSize; i++ {
// // // // 		for j := y; j < y+squareSize; j++ {
// // // // 			field[i][j] = 1
// // // // 		}
// // // // 	}
// // // // }

// // // // // Копирует двумерный массив и возвращает его копию
// // // // func copyArr(original [][]int) [][]int {
// // // // 	// Создаем новый срез с тем же количеством строк
// // // // 	copyGrid := make([][]int, len(original))

// // // // 	// Копируем каждую строку
// // // // 	for i := range original {
// // // // 		copyGrid[i] = make([]int, len(original[i])) // Создаем новый срез для каждой строки
// // // // 		copy(copyGrid[i], original[i])              // Копируем данные с помощью copy()
// // // // 	}

// // // // 	return copyGrid
// // // // }

// // // // func isFill(field [][]int) bool {
// // // // 	for i := 0; i < len(field); i++ {
// // // // 		for j := 0; j < len(field[i]); j++ {
// // // // 			if field[i][j] == 0 {
// // // // 				return false
// // // // 			}
// // // // 		}
// // // // 	}
// // // // 	return true
// // // // }

// // // // func findEmptySpace(field [][]int, sizeOfSquare int) []int {

// // // // 	for x := 0; x < len(field)-sizeOfSquare+1; x++ {

// // // // 		for y := 0; y < len(field[x])-sizeOfSquare+1; y++ {
// // // // 			if field[x][y] == 0 {
// // // // 				// 	return []int{x, y}
// // // // 				if checkPlace(field, x, y, sizeOfSquare) {
// // // // 					return []int{x, y}
// // // // 				}
// // // // 			}

// // // // 		}
// // // // 	}
// // // // 	return nil
// // // // }

// // // // // deleteSquare(state.field, state.lastSquare, state.usedCoords[len(state.usedCoords)-1])
// // // // func deleteSquare(field [][]int, x, y, squareSize int) {
// // // // 	for i := x; i < x+squareSize; i++ {
// // // // 		for j := y; j < y+squareSize; j++ {
// // // // 			field[i][j] = 0
// // // // 		}
// // // // 	}
// // // // }

// // // // /*ПРОХОДИТ ВСЕ ВЕТКИ, НО ОЧЕНЬ ДОЛГО ((((*/
// // // // // Замощаяет поле размером size*size, обязательно используя квадраты из массива squares
// // // // func backtracking(squares [][]int, size int, mode bool) ([][]int, error) {

// // // // 	field := make([][]int, size)
// // // // 	for i := 0; i < len(field); i++ {
// // // // 		field[i] = make([]int, size)
// // // // 	}

// // // // 	var stack []ExtraState
// // // // 	var bestPlacement [][]int
// // // // 	var placedSquares [][]int
// // // // 	var unplacedSquaresSizes []int
// // // // 	spinup := true

// // // // 	freeArea := size * size
// // // // 	// Размещаем квадраты, позиции которых заранее заданы
// // // // 	for _, sq := range squares {
// // // // 		if sq[0] != -1 {
// // // // 			if checkPlace(field, sq[0], sq[1], sq[2]) {
// // // // 				placeSquare(field, sq[0], sq[1], sq[2])
// // // // 				placedSquares = append(placedSquares, sq)
// // // // 				freeArea -= sq[2] * sq[2]

// // // // 			} else {
// // // // 				return nil, fmt.Errorf("%s", "Невозможно разместить квадраты таким образом")
// // // // 			}
// // // // 		} else {
// // // // 			unplacedSquaresSizes = append(unplacedSquaresSizes, sq[2])
// // // // 		}
// // // // 	}
// // // // 	fmt.Println(mode, "mode")
// // // // 	stack = append(stack, ExtraState{field, placedSquares, [][]int{}, (size + 1) / 2, freeArea})

// // // // 	fmt.Println(time.Now())
// // // // 	for len(stack) > 0 {
// // // // 		// if len(stack) == 2 && !spinup {
// // // // 		// 	return bestPlacement, nil
// // // // 		// }

// // // // 		// time.Sleep(800 * time.Millisecond)
// // // // 		state := stack[len(stack)-1]

// // // // 		// for i := range len(state.field) {
// // // // 		// 	fmt.Println(state.field[i])
// // // // 		// }
// // // // 		// // // // fmt.Println(len(stack), stack, state.freeArea, spinup, len(bestPlacement))
// // // // 		// fmt.Println()
// // // // 		// fmt.Println(len(stack))
// // // // 		// bufio.NewWriter(os.Stdin).WriteString(strconv.Atoi((len(stack))))

// // // // 		if state.freeArea == 0 {
// // // // 			if len(bestPlacement) > len(state.squares) || len(bestPlacement) == 0 {
// // // // 				bestPlacement = copyArr(state.squares)
// // // // 			}
// // // // 			// fmt.Println(bestPlacement)
// // // // 			// раскрутка закончилась
// // // // 			spinup = false
// // // // 			// удаляем все квадраты размер 1, потому перебор их перестановок не изменит результата
// // // // 			for state.lastSquare == 1 {
// // // // 				// !Нужно включить случай, когда расстановка полностью состоит из единичных квадратов (n = 2)!
// // // // 				if len(stack) == 0 {
// // // // 					break
// // // // 				}
// // // // 				stack = stack[:len(stack)-1]
// // // // 				state = stack[len(stack)-1]
// // // // 			}
// // // // 			for o := range size {
// // // // 				fmt.Println(state.field[o])
// // // // 			}
// // // // 			// fmt.Println(stack)
// // // // 			fmt.Println(state, len(bestPlacement), spinup, "-----------------------------------")
// // // // 		}
// // // // 		// если идёт скрутка, то нужно удалить поставленный на текущем шаге квадрат, чтобы подобрать для него другое место, имея возможность использовать часть занимаемых им ранее клеток
// // // // 		if !spinup {

// // // // 			// fmt.Println(state.usedCoords, state)
// // // // 			// если количество использованных координат на данном шаге - 0, значит, это самый первый шаг, у которого ни одного квадрата на поле выставлено не было
// // // // 			if len(state.usedCoords) == 0 {
// // // // 				break
// // // // 			}
// // // // 			deleteSquare(state.field, state.usedCoords[len(state.usedCoords)-1][0], state.usedCoords[len(state.usedCoords)-1][1], state.lastSquare)
// // // // 			// fmt.Println(stack)
// // // // 			// bufio.NewWriter(os.Stdin).WriteString(strconv.Itoa(len(stack)))
// // // // 			// fmt.Fprintf(bufio.NewWriter(os.Stdin), "%s ", strconv.Itoa(len(stack)))

// // // // 		}

// // // // 		/* Теперь, когда либо мы раскручиваем рекурсию, либо, когда мы ищем новую позицию для уже некогда поставленного квадрата,
// // // // 		перебираем все доступные размеры квадратов, начиная с поставленного ранее (брать больше нет смысла, потому что поставленный в прошлы раз квадрат занимал не менее четверти всей площади)
// // // // 		*/
// // // // 		stopFlag := false
// // // // 		// startStackLen := len(stack)
// // // // 		for k := state.lastSquare; k > 0; k-- {

// // // // 			if stopFlag {
// // // // 				// если стоит остановка, то выходим для удаления
// // // // 				break
// // // // 			}
// // // // 			/* Если рекурсия скручивается и мы пытаемся поставить квадратик с размером 1, то это бессмысленно, так как
// // // // 			это точно не уменьшит количество используемых квадратов, поэтому эти состояния можно отбросить, пока не будет найден больший квадрат*/
// // // // 			if !spinup && k == 1 {
// // // // 				stopFlag = true
// // // // 				continue
// // // // 			}
// // // // 			// Перебираем все возможные координаты для поиска подходящего места для квадрата
// // // // 			for x := 0; x < size-k+1; x++ {
// // // // 				if stopFlag {
// // // // 					// если стоит остановка, то переходим на цикл выше
// // // // 					break
// // // // 				}
// // // // 				for y := 0; y < size-k+1; y++ {
// // // // 					if stopFlag {
// // // // 						// если стоит остановка, то переходим на цикл выше
// // // // 						break
// // // // 					}
// // // // 					if checkPlace(state.field, x, y, k) {
// // // // 						// Просматриваем, что текущие координаты ранее не использовались на этом шаге
// // // // 						used := false
// // // // 						for _, coords := range state.usedCoords {
// // // // 							// Если текущие координаты были исопозованы, то переходим к следующим
// // // // 							if x == coords[0] && y == coords[1] && k == state.lastSquare {
// // // // 								used = true
// // // // 								break
// // // // 							}
// // // // 						}
// // // // 						if used {
// // // // 							continue
// // // // 						}
// // // // 						/*
// // // // 							Проверяем, количество квадратов текущего замощения не превышает количество квадратов в наилучшем замощении,
// // // // 							иначе нет смысла далее выстраивать замощение, т.к. оно будет гарантировано менее удачно
// // // // 						*/
// // // // 						if len(state.squares)+1 > len(bestPlacement) && len(bestPlacement) != 0 {
// // // // 							spinup = false
// // // // 							// После того, как стало ясно, что дальнейшее замощение бессмысленно, необходимо выйти из текущей цикла и начать процесс скрутки
// // // // 							// После остановки текущий шаг удалится из стек, потому что в него ничего не было добавлено
// // // // 							stopFlag = true
// // // // 							break
// // // // 						}
// // // // 						// Если текущими квадратами нельзя заполнить площадь более эффективно, то можно не продолжать перебор
// // // // 						if len(bestPlacement) != 0 && (state.freeArea/(k*k)+len(state.squares)) >= len(bestPlacement) {
// // // // 							spinup = false
// // // // 							stopFlag = true
// // // // 							break
// // // // 							// continue // Не может улучшить результат
// // // // 						}

// // // // 						// // Для текущего квадрата становиться на одну доступную позицию меньше
// // // // 						// state.usedCoords = append(state.usedCoords, []int{x, y})
// // // // 						// fmt.Println(state)
// // // // 						/*
// // // // 							Если координаты ранее не использовались, то нужно разместить квадрат,
// // // // 							пополнить список размещённых квадратов,
// // // // 							изменить поле (поле нужно перекопировать, потому что состояния поля на разных шагах различны)

// // // // 						*/
// // // // 						// Испоьльзованные координаты должны быть сохранены для того элемента, который был поставлен на этом шаге
// // // // 						// copyUsedCoords := copyArr(state.usedCoords)
// // // // 						// copyUsedCoords = append(copyUsedCoords, []int{x, y})

// // // // 						/*
// // // // 							На данный момент есть такой вариант: передавать только что использованные координаты, как used на следующий шаг,
// // // // 							тогда при возврате на текущем шаге будут точно храниться именно те координаты, которые были использованы.
// // // // 							При раскрутке можно класть в следующий, не обновляя у текущего, а при скрутке нужно будет не класть в следующий, но обновлять у текущего.
// // // // 							Потому что при раскрутке следующий шаг создаётся и помечается первая позиция, куда был поставлен квадрат, а при скрутке будет увеличиваться количество использованных позиций для текущего квадрата,
// // // // 							Нет, при скрутке нужно обновлять у текущего...

// // // // 							Нужно передавать на следующий. При скрутке нужно менять состояние текущего шага (изменять used), и оставить это состояние в стеке, продолжив замощение с него же

// // // // 							if spinup == false
// // // // 							При скрутке не нужно добавлять в стек новый шаг, нужно изменять либо удалять текущий.
// // // // 								// Нужно разместить новый квадрат не поле,
// // // // 								   Если k изменился, то нужно заменить последний элемент в массиве использованных, иначе не трогать этот массив
// // // // 								   Нужно пересчитать использованную площадь: если k уменьшилось нужно добавить квадрат lastSq и вычесть k^2, инече изменять не нужно

// // // // 							При переходе к новому k нужно очищать использованные координаты? - Да, похоже, что нужно,
// // // // 							потому что больше, чем k квадраты уже рассматриваться не будут, поэтому нет риска зацикливания, при этом позиции, занимаемые ранее бОльшим квадратом будут освобеждены
// // // // 						*/

// // // // 						// fmt.Println("here")
// // // // 						// Если идёт раскрутка
// // 	if spinup {
// // 		// Получаем копии объектов хранения, чтобы иметь возможность хранить разные их состояния на разных шагах
// // 		copyField := copyArr(state.field)
// // 		copyPlacedSquares := copyArr(state.squares)

// // 		placeSquare(copyField, x, y, k)
// // 		copyPlacedSquares = append(copyPlacedSquares, []int{x, y, k})
// // 		/*
// // 			Добавляем: скопированное поле, содержащее только что поставленный квадрат,
// // 					скопированный массив с использованными квадратами, в который был добавлен новый квадрат,
// // 					использованные координаты - пустой массив, потому что следующий квадрат ещё не был поставлен,
// // 					размер последнего поставленного квадрата,
// // 					размер пересчитанной свободной площади

// // 		*/
// // 		stack = append(stack, ExtraState{copyField, copyPlacedSquares, [][]int{{x, y}}, k, state.freeArea - (k * k)})
// // 		// После добавления квадрата необходимо продолжить заглублять рекурсию, поэтому из текущего цикла нужно выйти

// // 	} else { // Если скрутка
// // 		// Если k изменилось
// // 		if k != state.lastSquare {
// // 			// copyField := copyArr(field)
// // 			state.usedCoords = [][]int{{x, y}} // использованные координаты больше не актуальны, т.к. рассматривается новый квадрат, заменяем на массив с актуальными для текущего квалрата значениями
// // 			// Пересчитывается площадь, потому что ставиться квадрат другого размера
// // 			state.freeArea += state.lastSquare * state.lastSquare
// // 			state.freeArea -= k * k

// // 			// Мменяем последний квадрат на новый
// // 			state.lastSquare = k
// // 		} else {
// // 			// Если k не изменилось, то исопльзованные координаты пополняются новой позицией
// // 			state.usedCoords = append(state.usedCoords, []int{x, y})
// // 			// fmt.Println(state, " here ----------------------------------")
// // 		}
// // 		state.squares[len(state.squares)-1] = []int{x, y, k} // Убираем последний использованный квадрат, потому что теперь там будет другой

// // 		// квадрат размещается независимо от k
// // 		placeSquare(state.field, x, y, k) // размещаем квадрат (старый квадрат был удалён до входа в цикл)

// // 		stack[len(stack)-1] = state
// // 	}
// // 	// После остановки продолжится раскрутка, потому что размер стека изменился (устар)
// // 	stopFlag = true
// // 	// продолжаем раскрутку, потому началась новая ветвь рекурсии
// // 	spinup = true
// // }
// // // // 				}
// // // // 			}
// // // // 		}
// // // // 		/*
// // // // 			Если после всех попыток поставить квадрат на поле, это не произошло (размер стека не изменился),
// // // // 			то нужно возвращаться на шаг назад, напротив, если в стэк что-то было добавлено, то нужно продолжать заглублять рекурсию,
// // // // 			а удалять нельзя, потому что будет удалён только что добавленный элемент
// // // // 		*/
// // // // 		if !spinup && stopFlag {
// // // // 			// fmt.Println(len(stack))
// // // // 			stack = stack[:len(stack)-1]
// // // // 			// } else if startStackLen == len(stack) {
// // // // 		}
// // // // 	}
// // // // 	// if mode {
// // // // 	// 	/*
// // // // 	// 		При использовании небезопасного режима будут генерироваться все возможные перестановки квадартов с незаданными координатами.
// // // // 	// 		Это может Значительно увеличить время выполнения программы и занять Очень большой объём памяти при выполнении
// // // // 	// 	*/

// // // // 	// 	// Получение всех перестановок квадратов с незаданными координатами
// // // // 	// 	placements := iterativeBacktrack(field, unplacedSquaresSizes, size)

// // // // 	// 	if len(placements) == 0 {
// // // // 	// 		// Сохраняем состояния поля, при котором  квадраты уже расставлены
// // // // 	// 		stack = append(stack, State{copyArr(field), placedSquares, -1, freeArea})
// // // // 	// 	}
// // // // 	// 	// Размещение всех вариантов перестановок и сохранение их в стеке
// // // // 	// 	for _, plment := range placements {
// // // // 	// 		copyField := copyArr(field)
// // // // 	// 		copyPlacedSquares := copyArr(placedSquares)
// // // // 	// 		copyFreeArea := freeArea
// // // // 	// 		for _, sq := range plment {
// // // // 	// 			placeSquare(copyField, sq[0], sq[1], sq[2])
// // // // 	// 			copyPlacedSquares = append(copyPlacedSquares, sq)
// // // // 	// 			copyFreeArea -= sq[2] * sq[2]
// // // // 	// 		}
// // // // 	// 		stack = append(stack, State{copyField, copyPlacedSquares, -1, copyFreeArea})
// // // // 	// 	}
// // // // 	// } else {
// // // // 	// 	// Размещаем квадраты с неопределёнными позициями жадным образом
// // // // 	// 	for _, sq := range squares {
// // // // 	// 		if sq[0] == -1 {
// // // // 	// 			coords := findEmptySpace(field, sq[2])
// // // // 	// 			if coords != nil {
// // // // 	// 				placeSquare(field, coords[0], coords[1], sq[2])
// // // // 	// 				placedSquares = append(placedSquares, []int{coords[0], coords[1], sq[2]})
// // // // 	// 				freeArea -= sq[2] * sq[2]
// // // // 	// 			} else {
// // // // 	// 				return nil, fmt.Errorf("%s", "Такая расстановка невозможна")
// // // // 	// 			}
// // // // 	// 		}
// // // // 	// 	}
// // // // 	// 	// Сохраняем состояния поля, при котором  квадраты уже расставлены
// // // // 	// 	stack = append(stack, State{copyArr(field), placedSquares, -1, freeArea})
// // // // 	// }
// // // // 	// fmt.Println(stack, len(stack))

// // // // 	// Сохраняем состояния поля, при котором  квадраты уже расставлены
// // // // 	// stack = append(stack, State{copyArr(field), placedSquares, -1, freeArea})
// // // // 	// for i := range len(field) {
// // // // 	// 	fmt.Println(field[i])
// // // // 	// }
// // // // 	// fmt.Println(freeArea)
// // // // 	// for _, st := range stack {
// // // // 	// 	fmt.Println(st)
// // // // 	// }

// // // // 	// for len(stack) > 0 {
// // // // 	// 	time.Sleep(500 * time.Millisecond)
// // // // 	// 	currState := stack[len(stack)-1]
// // // // 	// 	stack = stack[:len(stack)-1]

// // // // 	// 	if (len(currState.squares) >= len(bestPlacement)) && (len(bestPlacement) != 0) {
// // // // 	// 		continue
// // // // 	// 	}

// // // // 	// 	// for i := range len(currState.field) {
// // // // 	// 	// 	fmt.Println(currState.field[i])
// // // // 	// 	// }
// // // // 	// 	// fmt.Println(currState.freeArea)

// // // // 	// 	if currState.freeArea == 0 {
// // // // 	// 		// fmt.Println(currState.squares, len(currState.squares), currState.freeArea)
// // // // 	// 		if len(bestPlacement) == 0 || len(bestPlacement) > len(currState.squares) {
// // // // 	// 			bestPlacement = currState.squares
// // // // 	// 		}
// // // // 	// 		continue
// // // // 	// 	}

// // // // 	// 	lastSquare := currState.lastSquare
// // // // 	// 	if lastSquare == -1 {
// // // // 	// 		lastSquare = (size + 1) / 2
// // // // 	// 	}

// // // // 	// startLen := len(stack)
// // // // 	// // for k := (size + 1) / 2; k > 0; k-- {
// // // // 	// for k := lastSquare; k > 0; k-- {
// // // // 	// 	if k == 1 {
// // // // 	// 		continue
// // // // 	// 	}
// // // // 	// 	for x := 0; x < size-k+1; x++ {
// // // // 	// 		for y := 0; y < size-k+1; y++ {
// // // // 	// 			if checkPlace(currState.field, x, y, k) {
// // // // 	// 				// fmt.Printf(" %v %v %v\n", x, y, k)
// // // // 	// 				copyFreeArea := currState.freeArea
// // // // 	// 				copyField := copyArr(currState.field)
// // // // 	// 				copyPlacedSquares := copyArr(currState.squares)

// // // // 	// 				placeSquare(copyField, x, y, k)
// // // // 	// 				copyFreeArea -= k * k
// // // // 	// 				copyPlacedSquares = append(copyPlacedSquares, []int{x, y, k})

// // // // 	// 				// if copyFreeArea == 0 && k == 1 {
// // // // 	// 				// 	return copyPlacedSquares, nil
// // // // 	// 				// }

// // // // 	// 				if copyFreeArea == 0 {
// // // // 	// 					// fmt.Println(currState.squares, len(currState.squares), currState.freeArea)
// // // // 	// 					if len(bestPlacement) == 0 || len(bestPlacement) > len(copyPlacedSquares) {
// // // // 	// 						bestPlacement = copyPlacedSquares
// // // // 	// 					}
// // // // 	// 					break
// // // // 	// 				}

// // // // 	// 				stack = append(stack, State{copyField, copyPlacedSquares, k, copyFreeArea})
// // // // 	// 			}
// // // // 	// 		}
// // // // 	// 		if k == 1 {
// // // // 	// 			break
// // // // 	// 		}
// // // // 	// 	}
// // // // 	// 	fmt.Println(len(stack), k, len(bestPlacement))
// // // // 	// 	if len(stack) != startLen {
// // // // 	// 		break
// // // // 	// 	}
// // // // 	// break
// // // // 	// for i := range len(currState.field) {
// // // // 	// 	fmt.Println(currState.field[i])
// // // // 	// }
// // // // 	// fmt.Println(currState.squares)
// // // // 	// fmt.Println(currState.freeArea)
// // // // 	// for x := 0; x < size-k+1; x++ {
// // // // 	// 	for y := 0; y <
// // // // 	// }
// // // // 	// coords := findEmptySpace(copyField, k)
// // // // 	// for coords != nil {
// // // // 	// 	if len(copyPlacedSquares) >= len(bestPlacement) && len(bestPlacement) != 0 {
// // // // 	// 		break
// // // // 	// 	}

// // // // 	// 	placeSquare(copyField, coords[0], coords[1], k)
// // // // 	// 	copyPlacedSquares = append(copyPlacedSquares, []int{coords[0], coords[1], k})
// // // // 	// 	copyFreeArea -= (k * k)
// // // // 	// 	// for i := range len(copyField) {
// // // // 	// 	// 	fmt.Println(copyField[i])
// // // // 	// 	// }
// // // // 	// 	// fmt.Println(copyFreeArea)
// // // // 	// 	// fmt.Println(k)

// // // // 	// 	stack = append(stack, State{copyField, copyPlacedSquares, k, copyFreeArea})
// // // // 	// 	fmt.Println(len(stack), k)
// // // // 	// 	coords = findEmptySpace(copyField, k)
// // // // 	// }
// // // // 	// }
// // // // 	// }

// // // // 	return bestPlacement, nil
// // // // }

// // // // // type State struct {
// // // // // 	Index      int
// // // // // 	Placements [][2]int
// // // // // 	Field      [][]bool
// // // // // }

// // // // // func canPlace(field [][]bool, x, y, size, fieldSize int) bool {
// // // // // 	for i := x; i < x+size; i++ {
// // // // // 		for j := y; j < y+size; j++ {
// // // // // 			if field[i][j] {
// // // // // 				return false
// // // // // 			}
// // // // // 		}
// // // // // 	}
// // // // // 	return true
// // // // // }

// // // // // // func placeSquare(field [][]bool, x, y, size int, state bool) {
// // // // // // 	for i := x; i < x+size; i++ {
// // // // // // 		for j := y; j < y+size; j++ {
// // // // // // 			field[i][j] = state
// // // // // // 		}
// // // // // // 	}
// // // // // // }

// // // // // Heavy Backtrack
// // // // // func iterativeBacktrack(field [][]int, squareSizes []int, fieldSize int) [][][]int {

// // // // // 	stack := []State{{lastSquare: 0, squares: [][]int{}, field: field}}
// // // // // 	validPlacements := [][][]int{}

// // // // // 	for len(stack) > 0 {
// // // // // 		// Достаем последний элемент стека
// // // // // 		state := stack[len(stack)-1]
// // // // // 		stack = stack[:len(stack)-1]

// // // // // 		if state.lastSquare == len(squareSizes) {
// // // // // 			validPlacements = append(validPlacements, append([][]int{}, state.squares...))
// // // // // 			continue
// // // // // 		}

// // // // // 		sqSize := squareSizes[state.lastSquare]

// // // // // 		for x := 0; x <= fieldSize-sqSize; x++ {
// // // // // 			for y := 0; y <= fieldSize-sqSize; y++ {
// // // // // 				if checkPlace(state.field, x, y, sqSize) {
// // // // // 					// Создаем копию текущего поля
// // // // // 					newField := copyArr(state.field)

// // // // // 					// Размещаем квадрат
// // // // // 					placeSquare(newField, x, y, sqSize)
// // // // // 					newPlacements := append([][]int{}, state.squares...)
// // // // // 					newPlacements = append(newPlacements, []int{x, y, sqSize})

// // // // // 					// Добавляем новое состояние в стек
// // // // // 					stack = append(stack, State{lastSquare: state.lastSquare + 1, squares: newPlacements, field: newField})
// // // // // 				}
// // // // // 			}
// // // // // 		}
// // // // // 	}

// // // // // 	return validPlacements
// // // // // }

// // // package main

// // // import (
// // // 	"bufio"
// // // 	"fmt"
// // // 	"os"
// // // 	// "strings"
// // // )

// // // // type ExtraState struct {
// // // // 	field      [][]int
// // // // 	squares    [][]int
// // // // 	usedCoords [][]int
// // // // 	lastSquare int
// // // // 	freeArea   int
// // // // }

// // // // Проверяет, можно ли разместить квадрат размером size*size на поле field, расположив верхний левый угол квадрата в координатах x, y поля
// // // // func checkPlace(field [][]int, x, y, squareSize int) bool {

// // // // 	if x+squareSize > len(field) || y+squareSize > len(field) {
// // // // 		// Если квадрат при наложении выходит за пределы поля, то его нельзя разместить
// // // // 		return false
// // // // 	}

// // // // 	for i := x; i < x+squareSize; i++ {
// // // // 		for j := y; j < y+squareSize; j++ {
// // // // 			// Если хотя бы одна клетка поля в рассматриваемой области занята, то размещение невозможно
// // // // 			if (field)[i][j] != 0 {
// // // // 				return false
// // // // 			}
// // // // 		}
// // // // 	}
// // // // 	return true
// // // // }

// // // // // Помечает клетки области поля, как "занятые", иметируя размещения квадрата на поле
// // // // func placeSquare(field [][]int, x, y, squareSize int) {
// // // // 	for i := x; i < x+squareSize; i++ {
// // // // 		for j := y; j < y+squareSize; j++ {
// // // // 			field[i][j] = 1
// // // // 		}
// // // // 	}
// // // // }

// // // // // Копирует двумерный массив и возвращает его копию
// // // // func copyArr(original [][]int) [][]int {
// // // // 	// Создаем новый срез с тем же количеством строк
// // // // 	copyGrid := make([][]int, len(original))

// // // // 	// Копируем каждую строку
// // // // 	for i := range original {
// // // // 		copyGrid[i] = make([]int, len(original[i])) // Создаем новый срез для каждой строки
// // // // 		copy(copyGrid[i], original[i])              // Копируем данные с помощью copy()
// // // // 	}

// // // // 	return copyGrid
// // // // }

// // // // // deleteSquare(state.field, state.lastSquare, *state.usedCoords[len(*state.usedCoords)-1])
// // // // func deleteSquare(field [][]int, x, y, squareSize int) {
// // // // 	for i := x; i < x+squareSize; i++ {
// // // // 		for j := y; j < y+squareSize; j++ {
// // // // 			field[i][j] = 0
// // // // 		}
// // // // 	}
// // // // }

// // // // // Замощаяет поле размером size*size, обязательно используя квадраты из массива squares
// // // // func backtracking(squares [][]int, size int) ([][]int, error) {

// // // // 	field := make([][]int, size)
// // // // 	for i := 0; i < len(field); i++ {
// // // // 		field[i] = make([]int, size)
// // // // 	}

// // // // 	var stack []ExtraState
// // // // 	var bestPlacement [][]int
// // // // 	var placedSquares [][]int
// // // // 	var unplacedSquaresSizes []int
// // // // 	spinup := true

// // // // 	freeArea := size * size
// // // // 	// Размещаем квадраты, позиции которых заранее заданы
// // // // 	for _, sq := range squares {
// // // // 		if sq[0] != -1 {
// // // // 			if checkPlace(field, sq[0], sq[1], sq[2]) {
// // // // 				placeSquare(field, sq[0], sq[1], sq[2])
// // // // 				placedSquares = append(placedSquares, sq)
// // // // 				freeArea -= sq[2] * sq[2]

// // // // 			} else {
// // // // 				return nil, fmt.Errorf("%s", "Невозможно разместить квадраты таким образом")
// // // // 			}
// // // // 		} else {
// // // // 			unplacedSquaresSizes = append(unplacedSquaresSizes, sq[2])
// // // // 		}
// // // // 	}

// // // // 	stack = append(stack, ExtraState{field, placedSquares, [][]int{}, -1, freeArea})

// // // // 	for len(stack) > 0 {
// // // // 		if len(stack) == 2 && !spinup {
// // // // 			return bestPlacement, nil
// // // // 		}

// // // // 		state := stack[len(stack)-1]

// // // // 		if state.freeArea == 0 {
// // // // 			if len(bestPlacement) > len(state.squares) || len(bestPlacement) == 0 {
// // // // 				bestPlacement = copyArr(state.squares)
// // // // 			}
// // // // 			// fmt.Println(bestPlacement)
// // // // 			// раскрутка закончилась
// // // // 			spinup = false
// // // // 			// удаляем все квадраты размер 1, потому перебор их перестановок не изменит результата
// // // // 			for state.lastSquare == 1 {
// // // // 				// !Нужно включить случай, когда расстановка полностью состоит из единичных квадратов (n = 2)!
// // // // 				if len(stack) == 0 {
// // // // 					break
// // // // 				}
// // // // 				stack = stack[:len(stack)-1]
// // // // 				state = stack[len(stack)-1]
// // // // 			}
// // // // 		}
// // // // 		// если идёт скрутка, то нужно удалить поставленный на текущем шаге квадрат, чтобы подобрать для него другое место, имея возможность использовать часть занимаемых им ранее клеток
// // // // 		if !spinup {

// // // // 			// если количество использованных координат на данном шаге - 0, значит, это самый первый шаг, у которого ни одного квадрата на поле выставлено не было
// // // // 			if len(state.usedCoords) == 0 {
// // // // 				break
// // // // 			}
// // // // 			deleteSquare(state.field, (state.usedCoords)[len(state.usedCoords)-1][0], (state.usedCoords)[len(state.usedCoords)-1][1], state.lastSquare)

// // // // 			// После удаления квадрата, эти координаты должны быть перемещены в начало списка, чтобы в его конце гарантированно находился квадрат, который точно размещён на поле
// // // // 			last := (state.usedCoords)[len(state.usedCoords)-1]                        // Сохраняем последний элемент
// // // // 			copy((state.usedCoords)[1:], (state.usedCoords)[:len(state.usedCoords)-1]) // Сдвигаем всё вправо
// // // // 			(state.usedCoords)[0] = last

// // // // 		}

// // // // 		/* Теперь, когда либо мы раскручиваем рекурсию, либо, когда мы ищем новую позицию для уже некогда поставленного квадрата,
// // // // 		перебираем все доступные размеры квадратов, начиная с поставленного ранее (брать больше нет смысла, потому что поставленный в прошлы раз квадрат занимал не менее четверти всей площади)
// // // // 		*/
// // // // 		stopFlag := false
// // // // 		// startStackLen := len(stack)
// // // // 		startSize := state.lastSquare
// // // // 		if state.lastSquare == -1 {
// // // // 			startSize = (size + 1) / 2
// // // // 			// startSize = (size - 1)
// // // // 		}
// // // // 		for k := startSize; k > 0; k-- {

// // // // 			if stopFlag {
// // // // 				// если стоит остановка, то выходим для удаления
// // // // 				break
// // // // 			}

// // // // 			if !spinup && k == 1 {
// // // // 				stopFlag = true
// // // // 				continue
// // // // 			}
// // // // 			// Перебираем все возможные координаты для поиска подходящего места для квадрата

// // // // 			for s := 0; s <= 2*(size-1); s++ {
// // // // 				if stopFlag {
// // // // 					// если стоит остановка, то переходим на цикл выше
// // // // 					break
// // // // 				}
// // // // 				for x := 0; x < size; x++ {
// // // // 					y := s - x
// // // // 					// Если j находится в допустимых пределах [0, size)
// // // // 					if y >= 0 && y < size {

// // // // 						if stopFlag {
// // // // 							// если стоит остановка, то переходим на цикл выше
// // // // 							break
// // // // 						}
// // // // 						if checkPlace(state.field, x, y, k) {
// // // // 							// Просматриваем, что текущие координаты ранее не использовались на этом шаге
// // // // 							used := false
// // // // 							for _, coords := range state.usedCoords {
// // // // 								// Если текущие координаты были исопозованы, то переходим к следующим
// // // // 								if x == coords[0] && y == coords[1] && k == state.lastSquare {
// // // // 									used = true
// // // // 									break
// // // // 								}
// // // // 							}
// // // // 							if used {
// // // // 								continue
// // // // 							}

// // // // 							if len(state.squares)+1 > len(bestPlacement) && len(bestPlacement) != 0 {
// // // // 								spinup = false
// // // // 								// После того, как стало ясно, что дальнейшее замощение бессмысленно, необходимо выйти из текущей цикла и начать процесс скрутки
// // // // 								// После остановки текущий шаг удалится из стек, потому что в него ничего не было добавлено
// // // // 								stopFlag = true
// // // // 								break
// // // // 							}
// // // // 							// Если текущими квадратами нельзя заполнить площадь более эффективно, то можно не продолжать перебор
// // // // 							if len(bestPlacement) != 0 && (state.freeArea/(k*k)+len(state.squares)) >= len(bestPlacement) {
// // // // 								spinup = false
// // // // 								stopFlag = true
// // // // 								break
// // // // 							}

// // if spinup {
// // 	// Получаем копии объектов хранения, чтобы иметь возможность хранить разные их состояния на разных шагах
// // 	copyField := copyArr(state.field)
// // 	copyPlacedSquares := copyArr(state.squares)

// // 	placeSquare(copyField, x, y, k)
// // 	copyPlacedSquares = append(copyPlacedSquares, []int{x, y, k})

// // 	if k == state.lastSquare {

// // 		state.usedCoords = append(state.usedCoords, []int{x, y})
// // 		stack = append(stack, ExtraState{copyField, copyPlacedSquares, state.usedCoords, k, state.freeArea - (k * k)})
// // 	} else {
// // 		stack = append(stack, ExtraState{copyField, copyPlacedSquares, [][]int{{x, y}}, k, state.freeArea - (k * k)})

// // 	}
// // 	// После добавления квадрата необходимо продолжить заглублять рекурсию, поэтому из текущего цикла нужно выйти

// // } else { // Если скрутка
// // 	// Если k изменилось
// // 	if k != state.lastSquare {
// // 		state.usedCoords = [][]int{{x, y}} // использованные координаты больше не актуальны, т.к. рассматривается новый квадрат, заменяем на массив с актуальными для текущего квалрата значениями
// // 		// Пересчитывается площадь, потому что ставиться квадрат другого размера
// // 		state.freeArea += state.lastSquare * state.lastSquare
// // 		state.freeArea -= k * k

// // 		// Мменяем последний квадрат на новый
// // 		state.lastSquare = k
// // 	} else {
// // 		// Если k не изменилось, то использованные координаты пополняются новой позицией
// // 		state.usedCoords = append(state.usedCoords, []int{x, y})
// // 	}
// // 	state.squares[len(state.squares)-1] = []int{x, y, k} // Убираем последний использованный квадрат, потому что теперь там будет другой

// // 	// квадрат размещается независимо от k
// // 	placeSquare(state.field, x, y, k) // размещаем квадрат (старый квадрат был удалён до входа в цикл)

// // 	stack[len(stack)-1] = state
// // }
// // // // 							// После остановки продолжится раскрутка, потому что размер стека изменился (устар)
// // // // 							stopFlag = true
// // // // 							// продолжаем раскрутку, потому началась новая ветвь рекурсии
// // // // 							spinup = true
// // // // 						}
// // // // 					}
// // // // 				}
// // // // 			}
// // // // 		}

// // // // 		if !spinup && stopFlag {
// // // // 			stack = stack[:len(stack)-1]

// // // // 		}
// // // // 	}
// // // // 	return bestPlacement, nil
// // // // }

// // // // import (
// // // // 	"bufio"
// // // // 	"fmt"
// // // // 	"os"

// // // // 	"strings"
// // // // )

// // // type ExtraState struct {
// // // 	field      [][]int
// // // 	squares    [][]int
// // // 	usedCoords *[][]int
// // // 	lastSquare int
// // // 	freeArea   int
// // // }

// // // // func main() {

// // // // 	reader := bufio.NewReader(os.Stdin)
// // // // 	// ioContr := libs.NewIoController()
// // // // 	var n int
// // // // 	fmt.Fscan(reader, &n)

// // // // 	field := make([][]int, n)
// // // // 	for i := 0; i < len(field); i++ {
// // // // 		field[i] = make([]int, n)
// // // // 	}

// // // // 	var squares [][]int

// // // // 	sqs, _ := backtracking(squares, n)

// // // // 	// fmt.Println(sqs, len(sqs))
// // // // 	fmt.Println(len(sqs))
// // // // 	for i := 0; i < len(sqs); i++ {
// // // // 		// str := strings.ReplaceAll(fmt.Sprint(sqs[i]), " ", "") // Убираем пробелы
// // // // 		fmt.Println(strings.Trim(fmt.Sprint(sqs[i]), "[]")) // Убираем квадратные скобки
// // // // 	}

// // // // }

// // // // Проверяет, можно ли разместить квадрат размером size*size на поле field, расположив верхний левый угол квадрата в координатах x, y поля
// // // func checkPlace(field [][]int, x, y, squareSize int) bool {

// // // 	if x+squareSize > len(field) || y+squareSize > len(field) {
// // // 		// Если квадрат при наложении выходит за пределы поля, то его нельзя разместить
// // // 		return false
// // // 	}

// // // 	for i := x; i < x+squareSize; i++ {
// // // 		for j := y; j < y+squareSize; j++ {
// // // 			// Если хотя бы одна клетка поля в рассматриваемой области занята, то размещение невозможно
// // // 			if (field)[i][j] != 0 {
// // // 				return false
// // // 			}
// // // 		}
// // // 	}
// // // 	return true
// // // }

// // // // Помечает клетки области поля, как "занятые", иметируя размещения квадрата на поле
// // // func placeSquare(field [][]int, x, y, squareSize int) {
// // // 	for i := x; i < x+squareSize; i++ {
// // // 		for j := y; j < y+squareSize; j++ {
// // // 			field[i][j] = 1
// // // 		}
// // // 	}
// // // }

// // // // Копирует двумерный массив и возвращает его копию
// // // func copyArr(original [][]int) [][]int {
// // // 	// Создаем новый срез с тем же количеством строк
// // // 	copyGrid := make([][]int, len(original))

// // // 	// Копируем каждую строку
// // // 	for i := range original {
// // // 		copyGrid[i] = make([]int, len(original[i])) // Создаем новый срез для каждой строки
// // // 		copy(copyGrid[i], original[i])              // Копируем данные с помощью copy()
// // // 	}

// // // 	return copyGrid
// // // }

// // // // deleteSquare(state.field, state.lastSquare, *state.usedCoords[len(*state.usedCoords)-1])
// // // func deleteSquare(field [][]int, x, y, squareSize int) {
// // // 	for i := x; i < x+squareSize; i++ {
// // // 		for j := y; j < y+squareSize; j++ {
// // // 			field[i][j] = 0
// // // 		}
// // // 	}
// // // }

// // // // Замощаяет поле размером size*size, обязательно используя квадраты из массива squares
// func backtracking(squares [][]int, size int) ([][]int, error) {

// 	field := make([][]int, size)
// 	for i := 0; i < len(field); i++ {
// 		field[i] = make([]int, size)
// 	}

// 	var stack []ExtraState
// 	var bestPlacement [][]int
// 	var placedSquares [][]int
// 	var unplacedSquaresSizes []int
// 	spinup := true

// 	freeArea := size * size
// 	// Размещаем квадраты, позиции которых заранее заданы
// 	for _, sq := range squares {
// 		if sq[0] != -1 {
// 			if checkPlace(field, sq[0], sq[1], sq[2]) {
// 				placeSquare(field, sq[0], sq[1], sq[2])
// 				placedSquares = append(placedSquares, sq)
// 				freeArea -= sq[2] * sq[2]

// 			} else {
// 				return nil, fmt.Errorf("%s", "Невозможно разместить квадраты таким образом")
// 			}
// 		} else {
// 			unplacedSquaresSizes = append(unplacedSquaresSizes, sq[2])
// 		}
// 	}

// 	stack = append(stack, ExtraState{field, placedSquares, &[][]int{}, -1, freeArea})

// 	for len(stack) > 0 {
// 		if len(stack) == 2 && !spinup {
// 			return bestPlacement, nil
// 		}

// 		state := stack[len(stack)-1]

// 		if state.freeArea == 0 {
// 			if len(bestPlacement) > len(state.squares) || len(bestPlacement) == 0 {
// 				bestPlacement = copyArr(state.squares)
// 			}
// 			// fmt.Println(bestPlacement)
// 			// раскрутка закончилась
// 			spinup = false
// 			// удаляем все квадраты размер 1, потому перебор их перестановок не изменит результата
// 			for state.lastSquare == 1 {
// 				// !Нужно включить случай, когда расстановка полностью состоит из единичных квадратов (n = 2)!
// 				if len(stack) == 0 {
// 					break
// 				}
// 				stack = stack[:len(stack)-1]
// 				state = stack[len(stack)-1]
// 			}
// 		}
// 		// если идёт скрутка, то нужно удалить поставленный на текущем шаге квадрат, чтобы подобрать для него другое место, имея возможность использовать часть занимаемых им ранее клеток
// 		if !spinup {

// 			// если количество использованных координат на данном шаге - 0, значит, это самый первый шаг, у которого ни одного квадрата на поле выставлено не было
// 			if len(*state.usedCoords) == 0 {
// 				break
// 			}
// 			deleteSquare(state.field, (*state.usedCoords)[len(*state.usedCoords)-1][0], (*state.usedCoords)[len(*state.usedCoords)-1][1], state.lastSquare)

// 			// После удаления квадрата, эти координаты должны быть перемещены в начало списка, чтобы в его конце гарантированно находился квадрат, который точно размещён на поле
// 			last := (*state.usedCoords)[len(*state.usedCoords)-1]                         // Сохраняем последний элемент
// 			copy((*state.usedCoords)[1:], (*state.usedCoords)[:len(*state.usedCoords)-1]) // Сдвигаем всё вправо
// 			(*state.usedCoords)[0] = last

// 		}

// 		/* Теперь, когда либо мы раскручиваем рекурсию, либо, когда мы ищем новую позицию для уже некогда поставленного квадрата,
// 		перебираем все доступные размеры квадратов, начиная с поставленного ранее (брать больше нет смысла, потому что поставленный в прошлы раз квадрат занимал не менее четверти всей площади)
// 		*/
// 		stopFlag := false
// 		// startStackLen := len(stack)
// 		startSize := state.lastSquare
// 		if state.lastSquare == -1 {
// 			// startSize = (size - 1)
// 			if size%3 == 0 {
// 				startSize = (size * 2) / 3
// 			} else {
// 				startSize = (size + 1) / 2
// 			}
// 		}
// 		for k := startSize; k > 0; k-- {

// 			if stopFlag {
// 				// если стоит остановка, то выходим для удаления
// 				break
// 			}

// 			if !spinup && k == 1 {
// 				stopFlag = true
// 				continue
// 			}
// 			// Перебираем все возможные координаты для поиска подходящего места для квадрата

// 			for s := 0; s <= 2*(size-1); s++ {
// 				if stopFlag {
// 					// если стоит остановка, то переходим на цикл выше
// 					break
// 				}
// 				for x := 0; x < size; x++ {
// 					y := s - x
// 					// Если j находится в допустимых пределах [0, size)
// 					if y >= 0 && y < size {

// 						if stopFlag {
// 							// если стоит остановка, то переходим на цикл выше
// 							break
// 						}
// 						if checkPlace(state.field, x, y, k) {
// 							// Просматриваем, что текущие координаты ранее не использовались на этом шаге
// 							used := false
// 							for _, coords := range *state.usedCoords {
// 								// Если текущие координаты были исопозованы, то переходим к следующим
// 								if x == coords[0] && y == coords[1] && k == state.lastSquare {
// 									used = true
// 									break
// 								}
// 							}
// 							if used {
// 								continue
// 							}

// 							if len(state.squares)+1 > len(bestPlacement) && len(bestPlacement) != 0 {
// 								spinup = false
// 								// После того, как стало ясно, что дальнейшее замощение бессмысленно, необходимо выйти из текущей цикла и начать процесс скрутки
// 								// После остановки текущий шаг удалится из стек, потому что в него ничего не было добавлено
// 								stopFlag = true
// 								break
// 							}
// 							// Если текущими квадратами нельзя заполнить площадь более эффективно, то можно не продолжать перебор
// 							if len(bestPlacement) != 0 && (state.freeArea/(k*k)+len(state.squares)) > len(bestPlacement) {
// 								spinup = false
// 								stopFlag = true
// 								break
// 							}

// 							if spinup {
// 								// Получаем копии объектов хранения, чтобы иметь возможность хранить разные их состояния на разных шагах
// 								copyField := copyArr(state.field)
// 								copyPlacedSquares := copyArr(state.squares)

// 								placeSquare(copyField, x, y, k)
// 								copyPlacedSquares = append(copyPlacedSquares, []int{x, y, k})

// 								if k == state.lastSquare {

// 									*state.usedCoords = append(*state.usedCoords, []int{x, y})
// 									stack = append(stack, ExtraState{copyField, copyPlacedSquares, state.usedCoords, k, state.freeArea - (k * k)})
// 								} else {
// 									stack = append(stack, ExtraState{copyField, copyPlacedSquares, &[][]int{{x, y}}, k, state.freeArea - (k * k)})

// 								}
// 								// После добавления квадрата необходимо продолжить заглублять рекурсию, поэтому из текущего цикла нужно выйти

// 							} else { // Если скрутка
// 								// Если k изменилось
// 								if k != state.lastSquare {
// 									state.usedCoords = &[][]int{{x, y}} // использованные координаты больше не актуальны, т.к. рассматривается новый квадрат, заменяем на массив с актуальными для текущего квалрата значениями
// 									// Пересчитывается площадь, потому что ставиться квадрат другого размера
// 									state.freeArea += state.lastSquare * state.lastSquare
// 									state.freeArea -= k * k

// 									// Мменяем последний квадрат на новый
// 									state.lastSquare = k
// 								} else {
// 									// Если k не изменилось, то использованные координаты пополняются новой позицией
// 									*state.usedCoords = append(*state.usedCoords, []int{x, y})
// 								}
// 								state.squares[len(state.squares)-1] = []int{x, y, k} // Убираем последний использованный квадрат, потому что теперь там будет другой

// 								// квадрат размещается независимо от k
// 								placeSquare(state.field, x, y, k) // размещаем квадрат (старый квадрат был удалён до входа в цикл)

// 								stack[len(stack)-1] = state
// 							}
// 							// После остановки продолжится раскрутка, потому что размер стека изменился (устар)
// 							stopFlag = true
// 							// продолжаем раскрутку, потому началась новая ветвь рекурсии
// 							spinup = true
// 						}
// 					}
// 				}
// 			}
// 		}

// 		if !spinup && stopFlag {
// 			stack = stack[:len(stack)-1]

// 		}
// 	}
// 	return bestPlacement, nil
// }

// // package main

// // import (
// // 	"bufio"
// // 	"fmt"
// // 	"os"
// // 	"time"
// // 	// "time"
// // 	// "strings"
// // )

// // type ExtraState struct {
// // 	field      [][]int
// // 	squares    [][]int
// // 	usedCoords *[][]int
// // 	lastSquare int
// // 	freeArea   int
// // }

// // func main() {

// // 	reader := bufio.NewReader(os.Stdin)
// // 	// ioContr := libs.NewIoController()
// // 	var n int
// // 	fmt.Fscan(reader, &n)
// // 	start := time.Now()
// // 	for i := 2; i < 21; i += 1 {
// // 		var squares [][]int

// // 		field := make([][]int, i)
// // 		for j := 0; j < len(field); j++ {
// // 			field[j] = make([]int, i)
// // 		}
// // 		// fmt.Println(sqs, len(sqs))
// // 		sqs, _ := backtracking(squares, i)

// // 		fmt.Println(i, ": ", len(sqs))
// // 		// for i := 0; i < len(sqs); i++ {
// // 		// 	sqs[i][0]++
// // 		// 	sqs[i][1]++
// // 		// 	// str := strings.ReplaceAll(fmt.Sprint(sqs[i]), " ", "") // Убираем пробелы
// // 		// 	fmt.Println(strings.Trim(fmt.Sprint(sqs[i]), "[]")) // Убираем квадратные скобки
// // 		// }
// // 	}
// // 	fmt.Println(time.Since(start))

// // 	// 	// field := make([][]int, n)
// // 	// 	// for j := 0; j < len(field); j++ {
// // 	// 	// 	field[j] = make([]int, n)
// // 	// 	// }
// // 	// 	// var squares [][]int

// // 	// 	// sqs, _ := backtracking(squares, n)

// // 	// 	// fmt.Println(len(sqs))
// // 	// 	// for i := 0; i < len(sqs); i++ {
// // 	// 	// 	sqs[i][0]++
// // 	// 	// 	sqs[i][1]++
// // 	// 	// 	// str := strings.ReplaceAll(fmt.Sprint(sqs[i]), " ", "") // Убираем пробелы
// // 	// 	// 	fmt.Println(strings.Trim(fmt.Sprint(sqs[i]), "[]")) // Убираем квадратные скобки
// // 	// 	// }
// // }

// // // func main() {

// // // 	reader := bufio.NewReader(os.Stdin)
// // // 	// ioContr := libs.NewIoController()
// // // 	var n int
// // // 	fmt.Fscan(reader, &n)

// // // 	field := make([][]int, n)
// // // 	for i := 0; i < len(field); i++ {
// // // 		field[i] = make([]int, n)
// // // 	}

// // // 	var squares [][]int

// // // 	sqs, _ := backtracking(squares, n)

// // // 	// fmt.Println(sqs, len(sqs))
// // // 	fmt.Println(len(sqs))
// // // 	for i := 0; i < len(sqs); i++ {
// // // 		sqs[i][0]++
// // // 		sqs[i][1]++
// // // 		// str := strings.ReplaceAll(fmt.Sprint(sqs[i]), " ", "") // Убираем пробелы
// // // 		fmt.Println(strings.Trim(fmt.Sprint(sqs[i]), "[]")) // Убираем квадратные скобки
// // // 	}

// // // }

// // // Проверяет, можно ли разместить квадрат размером size*size на поле field, расположив верхний левый угол квадрата в координатах x, y поля
// // func checkPlace(field [][]int, x, y, squareSize int) bool {

// // 	if x+squareSize > len(field) || y+squareSize > len(field) {
// // 		// Если квадрат при наложении выходит за пределы поля, то его нельзя разместить
// // 		return false
// // 	}

// // 	for i := x; i < x+squareSize; i++ {
// // 		for j := y; j < y+squareSize; j++ {
// // 			// Если хотя бы одна клетка поля в рассматриваемой области занята, то размещение невозможно
// // 			if (field)[i][j] != 0 {
// // 				return false
// // 			}
// // 		}
// // 	}
// // 	return true
// // }

// // // Помечает клетки области поля, как "занятые", иметируя размещения квадрата на поле
// // func placeSquare(field [][]int, x, y, squareSize int) {
// // 	for i := x; i < x+squareSize; i++ {
// // 		for j := y; j < y+squareSize; j++ {
// // 			field[i][j] = 1
// // 		}
// // 	}
// // }

// // // Копирует двумерный массив и возвращает его копию
// // func copyArr(original [][]int) [][]int {
// // 	// Создаем новый срез с тем же количеством строк
// // 	copyGrid := make([][]int, len(original))

// // 	// Копируем каждую строку
// // 	for i := range original {
// // 		copyGrid[i] = make([]int, len(original[i])) // Создаем новый срез для каждой строки
// // 		copy(copyGrid[i], original[i])              // Копируем данные с помощью copy()
// // 	}

// // 	return copyGrid
// // }

// // // deleteSquare(state.field, state.lastSquare, *state.usedCoords[len(*state.usedCoords)-1])
// // func deleteSquare(field [][]int, x, y, squareSize int) {
// // 	for i := x; i < x+squareSize; i++ {
// // 		for j := y; j < y+squareSize; j++ {
// // 			field[i][j] = 0
// // 		}
// // 	}
// // }

// // // Замощаяет поле размером size*size, обязательно используя квадраты из массива squares
// // func backtracking(squares [][]int, size int) ([][]int, error) {

// // 	field := make([][]int, size)
// // 	for i := 0; i < len(field); i++ {
// // 		field[i] = make([]int, size)
// // 	}

// // 	var stack []ExtraState
// // 	var bestPlacement [][]int
// // 	var placedSquares [][]int
// // 	var unplacedSquaresSizes []int
// // 	spinup := true

// // 	freeArea := size * size
// // 	// Размещаем квадраты, позиции которых заранее заданы
// // 	for _, sq := range squares {
// // 		if sq[0] != -1 {
// // 			if checkPlace(field, sq[0], sq[1], sq[2]) {
// // 				placeSquare(field, sq[0], sq[1], sq[2])
// // 				placedSquares = append(placedSquares, sq)
// // 				freeArea -= sq[2] * sq[2]

// // 			} else {
// // 				return nil, fmt.Errorf("%s", "Невозможно разместить квадраты таким образом")
// // 			}
// // 		} else {
// // 			unplacedSquaresSizes = append(unplacedSquaresSizes, sq[2])
// // 		}
// // 	}

// // 	stack = append(stack, ExtraState{field, placedSquares, &[][]int{}, -1, freeArea})

// // 	// cnt := 0

// // 	for len(stack) > 0 {
// // 		state := stack[len(stack)-1]
// // 		// time.Sleep(1000 * time.Millisecond)
// // 		// for i := range size {
// // 		// fmt.Println(state.field[i])
// // 		// }
// // 		// fmt.Println(len(stack), spinup, stack)

// // 		if state.freeArea == 0 {
// // 			if len(bestPlacement) > len(state.squares) || len(bestPlacement) == 0 {
// // 				bestPlacement = copyArr(state.squares)
// // 			}
// // 			// fmt.Println(bestPlacement)
// // 			// раскрутка закончилась
// // 			spinup = false
// // 			// удаляем все квадраты размер 1, потому перебор их перестановок не изменит результата
// // 			for state.lastSquare == 1 {
// // 				// !Нужно включить случай, когда расстановка полностью состоит из единичных квадратов (n = 2)!
// // 				if len(stack) == 0 {
// // 					break
// // 				}
// // 				stack = stack[:len(stack)-1]
// // 				state = stack[len(stack)-1]
// // 			}
// // 		}
// // 		// если идёт скрутка, то нужно удалить поставленный на текущем шаге квадрат, чтобы подобрать для него другое место, имея возможность использовать часть занимаемых им ранее клеток
// // 		if !spinup {

// // 			// если количество использованных координат на данном шаге - 0, значит, это самый первый шаг, у которого ни одного квадрата на поле выставлено не было
// // 			if len(*state.usedCoords) == 0 {
// // 				break
// // 			}
// // 			deleteSquare(state.field, (*state.usedCoords)[len(*state.usedCoords)-1][0], (*state.usedCoords)[len(*state.usedCoords)-1][1], state.lastSquare)

// // 			// После удаления квадрата, эти координаты должны быть перемещены в начало списка, чтобы в его конце гарантированно находился квадрат, который точно размещён на поле
// // 			last := (*state.usedCoords)[len(*state.usedCoords)-1]                         // Сохраняем последний элемент
// // 			copy((*state.usedCoords)[1:], (*state.usedCoords)[:len(*state.usedCoords)-1]) // Сдвигаем всё вправо
// // 			(*state.usedCoords)[0] = last

// // 		}
// // 		if len(stack) == 2 && !spinup {
// // 			state.usedCoords = &[][]int{}
// // 			state.freeArea += state.lastSquare * state.lastSquare
// // 			state.lastSquare -= 1
// // 			state.freeArea -= state.lastSquare * state.lastSquare

// // 			if state.lastSquare == size/2-1 {
// // 				return bestPlacement, nil
// // 			}
// // 			// cnt++
// // 			// 	if cnt == size {

// // 			// 	}
// // 		}

// // 		/* Теперь, когда либо мы раскручиваем рекурсию, либо, когда мы ищем новую позицию для уже некогда поставленного квадрата,
// // 		перебираем все доступные размеры квадратов, начиная с поставленного ранее (брать больше нет смысла, потому что поставленный в прошлы раз квадрат занимал не менее четверти всей площади)
// // 		*/
// // 		stopFlag := false
// // 		// startStackLen := len(stack)
// // 		startSize := state.lastSquare
// // 		if state.lastSquare == -1 {
// // 			if size%3 == 0 && size%2 != 0 {
// // 				startSize = (size * 2) / 3
// // 			} else {
// // 				startSize = (size - 1)
// // 				// startSize = (size + 1) / 2
// // 			}
// // 		}
// // 		for k := startSize; k > 0; k-- {

// // 			if stopFlag {
// // 				// если стоит остановка, то выходим для удаления
// // 				break
// // 			}

// // 			if !spinup && k == 1 {
// // 				stopFlag = true
// // 				continue
// // 			}
// // 			// Перебираем все возможные координаты для поиска подходящего места для квадрата

// // 			for s := 0; s <= 2*(size-1); s++ {
// // 				if stopFlag {
// // 					// если стоит остановка, то переходим на цикл выше
// // 					break
// // 				}
// // 				for x := 0; x < size; x++ {
// // 					y := s - x
// // 					// Если j находится в допустимых пределах [0, size)
// // 					if y >= 0 && y < size {

// // 						if stopFlag {
// // 							// если стоит остановка, то переходим на цикл выше
// // 							break
// // 						}
// // 						if checkPlace(state.field, x, y, k) {
// // 							// Просматриваем, что текущие координаты ранее не использовались на этом шаге
// // 							used := false
// // 							for _, coords := range *state.usedCoords {
// // 								// Если текущие координаты были исопозованы, то переходим к следующим
// // 								if x == coords[0] && y == coords[1] && k == state.lastSquare {
// // 									used = true
// // 									break
// // 								}
// // 							}
// // 							if used {
// // 								continue
// // 							}

// // 							if len(state.squares)+1 > len(bestPlacement) && len(bestPlacement) != 0 {
// // 								spinup = false
// // 								// После того, как стало ясно, что дальнейшее замощение бессмысленно, необходимо выйти из текущей цикла и начать процесс скрутки
// // 								// После остановки текущий шаг удалится из стек, потому что в него ничего не было добавлено
// // 								stopFlag = true
// // 								break
// // 							}
// // 							// Если текущими квадратами нельзя заполнить площадь более эффективно, то можно не продолжать перебор
// // 							if len(bestPlacement) != 0 && (state.freeArea/(k*k)+len(state.squares)) > len(bestPlacement) {
// // 								spinup = false
// // 								stopFlag = true
// // 								break
// // 							}

// // 							if spinup {
// // 								// Получаем копии объектов хранения, чтобы иметь возможность хранить разные их состояния на разных шагах
// // 								copyField := copyArr(state.field)
// // 								copyPlacedSquares := copyArr(state.squares)

// // 								placeSquare(copyField, x, y, k)
// // 								copyPlacedSquares = append(copyPlacedSquares, []int{x, y, k})

// // 								if k == state.lastSquare {

// // 									*state.usedCoords = append(*state.usedCoords, []int{x, y})
// // 									stack = append(stack, ExtraState{copyField, copyPlacedSquares, state.usedCoords, k, state.freeArea - (k * k)})
// // 								} else {
// // 									stack = append(stack, ExtraState{copyField, copyPlacedSquares, &[][]int{{x, y}}, k, state.freeArea - (k * k)})

// // 								}
// // 								// После добавления квадрата необходимо продолжить заглублять рекурсию, поэтому из текущего цикла нужно выйти

// // 							} else { // Если скрутка
// // 								// Если k изменилось
// // 								if k != state.lastSquare {
// // 									state.usedCoords = &[][]int{{x, y}} // использованные координаты больше не актуальны, т.к. рассматривается новый квадрат, заменяем на массив с актуальными для текущего квалрата значениями
// // 									// Пересчитывается площадь, потому что ставиться квадрат другого размера
// // 									state.freeArea += state.lastSquare * state.lastSquare
// // 									state.freeArea -= k * k

// // 									// Мменяем последний квадрат на новый
// // 									state.lastSquare = k
// // 								} else {
// // 									// Если k не изменилось, то использованные координаты пополняются новой позицией
// // 									*state.usedCoords = append(*state.usedCoords, []int{x, y})
// // 								}
// // 								state.squares[len(state.squares)-1] = []int{x, y, k} // Убираем последний использованный квадрат, потому что теперь там будет другой

// // 								// квадрат размещается независимо от k
// // 								placeSquare(state.field, x, y, k) // размещаем квадрат (старый квадрат был удалён до входа в цикл)

// // 								stack[len(stack)-1] = state
// // 							}
// // 							// После остановки продолжится раскрутка, потому что размер стека изменился (устар)
// // 							stopFlag = true
// // 							// продолжаем раскрутку, потому началась новая ветвь рекурсии
// // 							spinup = true
// // 						}
// // 					}
// // 				}
// // 			}
// // 		}

// // 		if !spinup && stopFlag {
// // 			stack = stack[:len(stack)-1]

// // 		}
// // 	}
// // 	return bestPlacement, nil
// // }
