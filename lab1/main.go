package main

import (
	"flag"
	libs "lab1/libs"
	"sort"
	"time"
)

type State struct {
	field             [][]int
	squares           [][]int
	usedCoords        [][]int
	symmetricalCoords [][]int
	lastSquare        int
	freeArea          int
	isRequired        bool
}

type ExtraState struct {
	field      [][]int
	squares    [][]int
	usedCoords *[][]int
	lastSquare int
	freeArea   int
}

func main() {

	ioContr := libs.NewIoController()

	useIndividualization := flag.Bool("ind", false, "Использовать ввод для индивидуального задания")
	useDebug := flag.Int("d", 0, "Использовать вывод промежуточных состояний")

	flag.Parse()
	debugStackDepth := *useDebug

	var n int
	var squares []int
	if !(*useIndividualization) {
		n = ioContr.ReadNum()
	} else {
		ioContr.Writeln("Введите число - размер стороны поля")
		n = ioContr.ReadNum()
		ioContr.Writeln("Введите размеры квадратов через пробел")

		squares = ioContr.ReadLineOfInts()
	}
	sort.Slice(squares, func(i, j int) bool {
		return squares[i] < squares[j]
	})

	sqs := advancedBacktracking(squares, n, debugStackDepth, ioContr)

	ioContr.Writeln(len(sqs))
	if *useIndividualization {
		ioContr.Writeln(sqs)
		Show(sqs, n)
	} else {
		for _, sq := range sqs {
			ioContr.Writeln(sq[0]+1, sq[1]+1, sq[2])
		}
	}
}

// Проверяет, можно ли разместить квадрат размером size*size на поле field, расположив верхний левый угол квадрата в координатах x, y поля
func checkPlace(field [][]int, x, y, squareSize int) bool {

	if x+squareSize > len(field) || y+squareSize > len(field) {
		// Если квадрат при наложении выходит за пределы поля, то его нельзя разместить
		return false
	}

	for i := x; i < x+squareSize; i++ {
		for j := y; j < y+squareSize; j++ {
			// Если хотя бы одна клетка поля в рассматриваемой области занята, то размещение невозможно
			if (field)[i][j] != 0 {
				return false
			}
		}
	}
	return true
}

// Помечает клетки области поля, как "занятые", иметируя размещения квадрата на поле
func placeSquare(field [][]int, x, y, squareSize int) {
	for i := x; i < x+squareSize; i++ {
		for j := y; j < y+squareSize; j++ {
			field[i][j] = 1
		}
	}
}

// Копирует двумерный массив и возвращает его копию
func copyArr(original [][]int) [][]int {
	// Создаем новый срез с тем же количеством строк
	copyGrid := make([][]int, len(original))

	// Копируем каждую строку
	for i := range original {
		copyGrid[i] = make([]int, len(original[i])) // Создаем новый срез для каждой строки
		copy(copyGrid[i], original[i])              // Копируем данные с помощью copy()
	}

	return copyGrid
}

// удаляет квадрат с поля
func deleteSquare(field [][]int, x, y, squareSize int) {
	for i := x; i < x+squareSize; i++ {
		for j := y; j < y+squareSize; j++ {
			field[i][j] = 0
		}
	}
}

// Замощаяет поле размером size*size
func backtracking(size, db int, ioC *libs.InputOutputController) [][]int {

	field := make([][]int, size)
	for i := 0; i < len(field); i++ {
		field[i] = make([]int, size)
	}

	var stack []ExtraState
	var bestPlacement [][]int // храниться текущая наилучшая расстановка
	var placedSquares [][]int // текущие установленные квадраты
	spinup := true            // состояние работы функции: стек либо заполняется, либо освобождается

	freeArea := size * size // свободная плащадь, куда можно поставить квадраты

	// стартовое сотояние - пустое поле
	stack = append(stack, ExtraState{field, placedSquares, &[][]int{}, -1, freeArea})

	// если сторона квадрата - простое число, отличное от 2 и 3, то можно выставить 3 квадрата, которые сократят рассматриваемую область примерно на 75%
	if size%3 != 0 && size%2 != 0 {
		freeArea -= ((size + 1) / 2 * (size + 1) / 2)
		freeArea -= 2 * (((size+1)/2 - 1) * ((size+1)/2 - 1))

		placedSquares = append(placedSquares, []int{0, size/2 + 1, (size+1)/2 - 1})
		placedSquares = append(placedSquares, []int{size/2 + 1, 0, (size+1)/2 - 1})
		placedSquares = append(placedSquares, []int{size - (size/2 + 1), size - (size/2 + 1), (size + 1) / 2})

		placeSquare(field, size-(size/2+1), size-(size/2+1), (size+1)/2)
		placeSquare(field, 0, size/2+1, (size+1)/2-1)
		placeSquare(field, size/2+1, 0, (size+1)/2-1)

		// расстановка с тремя квадратами (базовая расстановка) записывается как единое состояние
		stack = append(stack, ExtraState{field, placedSquares, &[][]int{}, (size+1)/2 - 1, freeArea})
	}

	for len(stack) > 0 {
		// если программа вернулась к состоянию с базовой расстановкой - дальнейший обход нецелесообразен
		if len(stack) == 2 && !spinup {
			return bestPlacement
		}
		// последнее помещённое в стек состояние
		state := stack[len(stack)-1]

		// вывод поля при флаге -d
		if len(stack) < db && spinup {
			time.Sleep(1000 * time.Millisecond)
			Show(state.squares, size)
			ioC.Writeln()
		}
		// если состояние хранит заполненное поле
		if state.freeArea == 0 {
			// проверка, что расстановка из текущего состояния лучше, чем наилучшая на данный момент расстановка
			if len(bestPlacement) > len(state.squares) || len(bestPlacement) == 0 {
				bestPlacement = copyArr(state.squares)
				if db != 0 {
					ioC.Writeln("Новая наилучая расстановка:")
					ioC.Writeln(bestPlacement)
				}
			}
			// раскрутка закончилась
			spinup = false

			// удаляем все квадраты размера 1, потому перебор их перестановок не изменит результата
			for state.lastSquare == 1 {
				if len(stack) == 0 {
					break
				}
				// извлекаем состояние из стека
				stack = stack[:len(stack)-1]
				state = stack[len(stack)-1]
			}
		}

		// если идёт скрутка, то нужно удалить поставленный на текущем шаге квадрат,
		// чтобы подобрать для него другое место или поставить квадрат меньшего размера
		if !spinup {
			// если никакие координаты для текущего шага не использованы, то нужно переходить сразу к установке квадрата на поле
			if len(*state.usedCoords) == 0 {
				break
			}

			deleteSquare(state.field, (*state.usedCoords)[len(*state.usedCoords)-1][0], (*state.usedCoords)[len(*state.usedCoords)-1][1], state.lastSquare)
			// После удаления квадрата, эти координаты должны быть перемещены в начало списка, чтобы в конце этого списка гарантированно находился квадрат, размещённый на поле
			last := (*state.usedCoords)[len(*state.usedCoords)-1]                         // Сохраняем последний элемент
			copy((*state.usedCoords)[1:], (*state.usedCoords)[:len(*state.usedCoords)-1]) // Сдвигаем всё вправо
			(*state.usedCoords)[0] = last

		}

		stopFlag := false             // отвечает за то, нужно ли искать место для квадрата
		startSize := state.lastSquare // начальный размер устанавливаемого квадрата

		// если начальный размер не указан, то нужно либо начать с квадрата размером 2/3 от стороны поля (для кратных 3), либо 1/2 от стороны поля (для кратных 2)
		if state.lastSquare == -1 {
			if size%3 == 0 && size%2 != 0 {
				startSize = (size * 2) / 3
			} else {
				startSize = (size + 1) / 2
			}
		}

		for k := startSize; k > 0; k-- {
			if stopFlag {
				// если стоит остановка, то выходим для удаления
				break
			}
			/* Если стек сворачиваестя и мы пытаемся поставить квадрат с размером 1, то это бессмысленно, так как
			это точно не уменьшит количество используемых квадратов, поэтому эти состояния можно отбросить, пока не будет найден больший квадрат*/
			// сворачивание стек гарантирует, что какая-то расстановка уже найдена
			if !spinup && k == 1 {
				stopFlag = true
				continue
			}

			// Перебираем все возможные координаты для поиска подходящего места для квадрата
			for x := 0; x < size-k+1; x++ {
				if stopFlag {
					// если стоит остановка, то переходим на цикл выше
					break
				}

				for y := 0; y < size-k+1; y++ {
					if stopFlag {
						// если стоит остановка, то переходим на цикл выше
						break
					}
					if checkPlace(state.field, x, y, k) {
						// Просматриваем, что текущие координаты ранее не использовались на этом шаге
						used := false
						for _, coords := range *state.usedCoords {
							// Если текущие координаты были исопользованы, то переходим к следующим
							if x == coords[0] && y == coords[1] && k == state.lastSquare {
								used = true
								break
							}
						}
						if used {
							continue
						}

						// если дальнейшее размещение не уменьшит количество квадратов относительно текущей лучшей расстановки, то можно начинать скручивать стек
						if len(state.squares)+1 > len(bestPlacement) && len(bestPlacement) != 0 {
							spinup = false
							// После остановки текущий шаг удалится из стек, потому что в него ничего не было добавлено
							stopFlag = true
							break
						}
						// Если идёт раскрутка
						if spinup {
							// Получаем копии объектов хранения, чтобы иметь возможность хранить разные их состояния на разных шагах
							copyField := copyArr(state.field)
							copyPlacedSquares := copyArr(state.squares)

							placeSquare(copyField, x, y, k)
							copyPlacedSquares = append(copyPlacedSquares, []int{x, y, k})

							if k == state.lastSquare {
								// если размер установленного только что квадрата равен установленному ранее, то к списку использованных координат добавляются текущие
								*state.usedCoords = append(*state.usedCoords, []int{x, y})

								stack = append(stack, ExtraState{copyField, copyPlacedSquares, state.usedCoords, k, state.freeArea - (k * k)})
							} else {
								// если размеры отличаются, то нужно создать новый список
								stack = append(stack, ExtraState{copyField, copyPlacedSquares, &[][]int{{x, y}}, k, state.freeArea - (k * k)})

							}

						} else { // Если скрутка

							// если стек скручивался, и был найден квадрат, который можно установить
							if k != state.lastSquare {
								state.usedCoords = &[][]int{{x, y}} // использованные координаты больше не актуальны, т.к. рассматривается новый квадрат
								// Пересчитывается площадь, потому что ставиться квадрат другого размера
								state.freeArea += state.lastSquare * state.lastSquare
								state.freeArea -= k * k

								// Мменяем последний квадрат на новый
								state.lastSquare = k
							} else {
								// Если k не изменилось, то использованные координаты пополняются новой позицией
								*state.usedCoords = append(*state.usedCoords, []int{x, y})
							}
							state.squares[len(state.squares)-1] = []int{x, y, k} // Убираем последний использованный квадрат, потому что теперь там будет установленный только что

							placeSquare(state.field, x, y, k) // размещаем квадрат на поле

							stack[len(stack)-1] = state
						}
						// квадрат установлен, поэтому нужно перейти к проверке нового состояния из стека
						stopFlag = true
						// после установки квадрата стек обязательно должен начать раскручиваться
						spinup = true

						/* если послений установленный квадрат имеет размер 1 (это значит, что и остальные квадраты будут не больше 1),
						площадь которую нужно ими замостить + количество уже установленных квадратов превысит количество квадрато из лучшей расстановки,
						то можно начинать скручивать стек*/
						if len(bestPlacement) != 0 && ((state.freeArea + len(state.squares)) >= len(bestPlacement)) && k == 1 {
							spinup = false
							break
						}
					}

				}
			}

		}
		/*
			Если после всех попыток поставить квадрат на поле, это не произошло (размер стека не изменился),
			то нужно возвращаться на шаг назад
		*/

		if !spinup && stopFlag {
			stack = stack[:len(stack)-1]
		}
	}

	return bestPlacement
}

// ищет расстановку с учётом квадратов, которые обязательно нужно расставить
func advancedBacktracking(squares []int, size int, db int, ioC *libs.InputOutputController) [][]int {
	field := make([][]int, size)
	for i := 0; i < len(field); i++ {
		field[i] = make([]int, size)
	}
	var stack []State
	var bestPlacement [][]int

	freeArea := size * size

	// получаем расстановку без учёта необходимых квадратов
	var cleanPlacement [][]int
	if len(squares) == 0 {
		cleanPlacement = backtracking(size, db, ioC)
	} else {
		for _, sq := range squares {
			if sq >= size {
				ioC.Writeln("Некорректный набор квадратов")
				return [][]int{}
			}
		}
		cleanPlacement = backtracking(size, 0, ioC)

	}

	if len(squares) == 0 {
		return cleanPlacement
	}
	if len(cleanPlacement) == 0 {
		return [][]int{}
	}
	// сортировка квадратов из расстановки по возрастанию
	sort.Slice(cleanPlacement, func(a, b int) bool {
		return cleanPlacement[a][2] < cleanPlacement[b][2]
	})

	// Сохранили расстановку идентичных с оптимальной расстановкой квадратов
	bufferedSquares := make([]int, len(squares))
	copy(bufferedSquares, squares)
	for _, sq := range cleanPlacement {
		if ind := bSearch(squares, sq[2], 0, len(squares)-1); ind != -1 {
			squares = append(squares[:ind], squares[ind+1:]...)
		}
	}
	// Если все обязательные квадраты состоят в оптимальной расстановке, то она и будет результатом
	if len(squares) == 0 {
		return cleanPlacement
	} else {
		squares = bufferedSquares
	}

	// Получили нерасcтавленные квадраты и количество каждого вида (ассоциативный массив)
	unplacedSq := make([]int, size)
	for _, sq := range squares {
		unplacedSq[sq]++
	}

	placement := make([][]int, 0, len(squares))

	// Стартовое состояние
	stack = append(stack, State{copyArr(field), placement, [][]int{}, [][]int{}, -1, freeArea, false})

	bestPlacement = [][]int{}

	spinup := true
	for len(stack) > 0 {

		state := stack[len(stack)-1]
		if len(stack) < db && spinup {
			time.Sleep(1000 * time.Millisecond)
			Show(state.squares, size)
			ioC.Writeln()
		}

		if len(state.usedCoords) == 0 && !spinup {
			return bestPlacement
		}

		if state.freeArea == 0 {
			if len(bestPlacement) > len(state.squares) || len(bestPlacement) == 0 {
				bestPlacement = copyArr(state.squares)
				if db != 0 {
					ioC.Writeln("Новая наилучая расстановка:")
					ioC.Writeln(bestPlacement)
				}
			}
			// раскрутка закончилась
			spinup = false
			// удаляем все квадраты размер 1, потому перебор их перестановок не изменит результата
			for state.lastSquare == 1 {
				// !Нужно включить случай, когда расстановка полностью состоит из единичных квадратов (n = 2)!
				if len(stack) == 0 {
					break
				}
				if state.isRequired {
					unplacedSq[state.lastSquare]++
				}
				stack = stack[:len(stack)-1]
				state = stack[len(stack)-1]
			}
		}

		if !spinup {
			// если количество использованных координат на данном шаге - 0, значит, это самый первый шаг, у которого ни одного квадрата на поле выставлено не было
			if len(state.usedCoords) == 0 {
				break
			}
			deleteSquare(state.field, (state.usedCoords)[len(state.usedCoords)-1][0], (state.usedCoords)[len(state.usedCoords)-1][1], state.lastSquare)
			if state.isRequired {
				unplacedSq[state.lastSquare]++
			}

			// После удаления квадрата, эти координаты должны быть перемещены в начало списка, чтобы в его конце гарантированно находился квадрат, который точно размещён на поле
			last := (state.usedCoords)[len(state.usedCoords)-1]                        // Сохраняем последний элемент
			copy((state.usedCoords)[1:], (state.usedCoords)[:len(state.usedCoords)-1]) // Сдвигаем всё вправо
			(state.usedCoords)[0] = last

		}

		startSize := state.lastSquare
		if state.isRequired || startSize == -1 {
			// в данной функции стартовый квадрат нельзя выбрать оптимально, необходимо проверять квадраты всех размеров
			startSize = size - 1
		}

		stopFlag := false
		for k := startSize; k > 0; k-- {
			if stopFlag {
				break
			}
			/* Если стек скручивается и мы пытаемся поставить квадрат с размером 1, то это бессмысленно, так как
			это точно не уменьшит количество используемых квадратов, поэтому эти состояния можно отбросить, пока не будет найден больший квадрат*/
			if !spinup && k == 1 && !(state.isRequired) {
				stopFlag = true
				continue
			}

			useUnplaced := false // указывает на то, используется ли сейчас квадрат, который обязательно нужно поставить
			if spinup || (!spinup && state.isRequired) {
				// Если идёт раскрутка и из обязательных остались нерасставленные квадраты - нужно поставить наибольший
				for s := len(unplacedSq) - 1; s > 0; s-- {
					if unplacedSq[s] > 0 {
						k = s
						useUnplaced = true
						unplacedSq[s]--
						break
					}
				}
			}
			// проходим по всем координатам, выбираем место для квадрата
			for x := 0; x < size+1-k; x++ {
				if stopFlag {
					break
				}
				for y := 0; y < size+1-k; y++ {
					if stopFlag {
						break
					}
					if checkPlace(state.field, x, y, k) {

						used := false
						for _, coords := range state.usedCoords {
							// Если текущие координаты были использованы, то переходим к следующим
							if x == coords[0] && y == coords[1] && k == state.lastSquare {
								used = true
								break
							}
						}
						// провекрка симметричных позиций возможно только при скрутке
						if !spinup {
							// проверка
							for _, coords := range state.symmetricalCoords {
								if x == coords[0] && y == coords[1] && k == state.lastSquare {
									used = true
									break
								}
							}
						}
						if used {
							continue
						}

						if len(state.squares)+1 > len(bestPlacement) && len(bestPlacement) != 0 {
							spinup = false
							// После того, как стало ясно, что дальнейшее замощение бессмысленно, необходимо выйти из текущей цикла и начать процесс скрутки
							// После остановки текущий шаг удалится из стека, потому что в него ничего не было добавлено
							stopFlag = true
							break
						}
						/* если сумма кол-ва установленных квадратов и квадратов, которые можно расставить на свободной площади
						больше к-ва квадратов из лучшей расстановки, то можно скручивать стек*/
						if len(bestPlacement) != 0 && (state.freeArea/(k*k)+len(state.squares)) > len(bestPlacement) && !useUnplaced && k > 1 {
							spinup = false
							stopFlag = true
							break
						}
						// если квадрат был поставлен при раскрутке
						if spinup {
							// получание копий
							copyField := copyArr(state.field)
							copyPlacement := copyArr(state.squares)

							// установка квадрата
							placeSquare(copyField, x, y, k)
							copyPlacement = append(copyPlacement, []int{x, y, k})

							// рассчёт симметричных координат
							symCoords := [][]int{}
							mirrowedX := size - (y + k)
							mirrowedY := size - (x + k)
							if x != y {
								symCoords = append(symCoords, []int{y, x})
							}
							if x != size-(y+k) {
								symCoords = append(symCoords, []int{mirrowedX, mirrowedY})
							}
							symCoords = append(symCoords, []int{mirrowedX, size - (mirrowedY + k)}, []int{size - (mirrowedX + k), mirrowedY})
							symCoords = append(symCoords, []int{x, size - y - k}, []int{size - x - k, y})

							// если размер квадрата такой же, как и у предыдущего, то нужно сохранить координаты предыдущего для текущего, что бы избеджать проверки перестановочных случаев
							if k == state.lastSquare {
								copyUsedCoords := copyArr(state.usedCoords)
								copyUsedCoords = append(copyUsedCoords, []int{x, y})

								stack = append(stack, State{copyField, copyPlacement, copyUsedCoords, symCoords, k, state.freeArea - (k * k), useUnplaced})
							} else { // иначе создаётся новый массив с координатами

								stack = append(stack, State{copyField, copyPlacement, [][]int{{x, y}}, symCoords, k, state.freeArea - (k * k), useUnplaced})
							}

						} else { // если квадрат был поставлен при скрутке
							// получение симметричных координат
							mirrowedX := size - (y + k)
							mirrowedY := size - (x + k)
							if k == state.lastSquare {
								if x != y {
									// symCoords = append(symCoords, []int{y, x})
									state.symmetricalCoords = append(state.symmetricalCoords, []int{y, x})
								}
								if x != size-(y+k) {
									state.symmetricalCoords = append(state.symmetricalCoords, []int{mirrowedX, mirrowedY})
								}
								state.symmetricalCoords = append(state.symmetricalCoords, []int{mirrowedX, size - (mirrowedY + k)}, []int{size - (mirrowedX + k), mirrowedY})
								state.symmetricalCoords = append(state.symmetricalCoords, []int{size - x - k, y}, []int{x, size - y - k})
								state.usedCoords = append(state.usedCoords, []int{x, y})
							} else {
								state.usedCoords = [][]int{{x, y}}
								state.symmetricalCoords = [][]int{{mirrowedX, size - (mirrowedY + k)}, {size - (mirrowedX + k), mirrowedY}, {y, x}, {x, size - y - k}, {size - (y + k), size - (x + k)}, {size - x - k, y}}

								state.freeArea += state.lastSquare * state.lastSquare
								state.freeArea -= k * k
								// Меняем последний квадрат на новый
								state.lastSquare = k
							}

							state.squares[len(state.squares)-1] = []int{x, y, k} // Убираем последний использованный квадрат, потому что теперь там будет другой

							placeSquare(state.field, x, y, k) // размещаем квадрат

							stack[len(stack)-1] = state
						}

						stopFlag = true // переход к следующему состоянию
						useUnplaced = false
						spinup = true

						// проверка, что имеет смысл ставить квадраты размером 1, если они не относятся к обязательным
						if len(bestPlacement) != 0 && ((state.freeArea/(k*k) + len(state.squares)) >= len(bestPlacement)) && k == 1 && !(state.isRequired) {
							spinup = false
							break
						}
					}

				}
			}

			if useUnplaced {
				// если был выбран квадрат обязательный к установке, и он не был поставлен, то нужно начать скручивать стек и вернуть квадрат в массив unplacedSq
				stopFlag = true
				if spinup || k != state.lastSquare {
					unplacedSq[k]++
				}
				spinup = false
			}
		}
		// Если идёт расркутка и ни один квадрат не был установлен, нужно начинать скручивать стек
		if spinup && !stopFlag {
			spinup = false
		}
		// используется, когда скрутка и мест для квадрата больше не осталось
		if !spinup && stopFlag {
			if state.isRequired {
				unplacedSq[state.lastSquare]++
			}
			stack = stack[:len(stack)-1]
		}
	}

	return bestPlacement
}

// бинарный поиск ключа key в массиве  nums
func bSearch(nums []int, key, lo, hi int) int {
	if lo > hi {
		return -1
	}
	mi := lo + (hi-lo)/2
	if key == nums[mi] {
		return mi
	}
	if key > nums[mi] {
		return bSearch(nums, key, mi+1, hi)
	}
	return bSearch(nums, key, lo, mi-1)
}
