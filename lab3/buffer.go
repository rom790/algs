package main

/*
import (
	"bufio"
	"fmt"
	"os"
)

const (
	k = 6 // Размер алфавита: {A, C, G, T, N}
)

type Vertex struct {
	children        [k]int // Дети узла
	ix              int    // Номер образца, если узел терминальный
	suffixIx        int    // Суффиксная ссылка
	advanceVertexes [k]int // Автоматические переходы
	par             int    // Родительский узел
	goodSuffixIx    int    // Ссылка на ближайший терминальный узел
	flag            bool   // Флаг терминального узла
	symbol          byte   // Символ, по которому перешли в этот узел
}

var (
	bohr     []Vertex       // Бор
	patterns []string       // Список образцов
	joker    byte     = '$' // Символ-джокер
)

// Преобразует символ в индекс
func charToIndex(c byte) int {
	switch c {
	case 'A':
		return 0
	case 'C':
		return 1
	case 'G':
		return 2
	case 'T':
		return 3
	case 'N':
		return 4
	default:
		return k - 1 // Для джокера
	}
}

// Преобразует индекс в символ
func indexToChar(i int) byte {
	switch i {
	case 0:
		return 'A'
	case 1:
		return 'C'
	case 2:
		return 'G'
	case 3:
		return 'T'
	case 4:
		return 'N'
	default:
		// panic("Недопустимый индекс: " + fmt.Sprint(i))
		return k - 1
	}
}

// Создает новый узел бора
func createVertex(p int, c byte) Vertex {
	v := Vertex{}
	for i := 0; i < k; i++ {
		v.children[i] = -1
		v.advanceVertexes[i] = -1
	}
	v.flag = false
	v.suffixIx = -1
	v.par = p
	v.symbol = c
	v.goodSuffixIx = -1
	return v
}

// Инициализирует бор
func bohrInit() {
	bohr = append(bohr, createVertex(0, '#'))
}

// Добавляет строку в бор
func addString(s string) {
	curLast := 0
	for i := 0; i < len(s); i++ {
		ch := charToIndex(s[i])

		if bohr[curLast].children[ch] == -1 {
			bohr = append(bohr, createVertex(curLast, s[i]))
			bohr[curLast].children[ch] = len(bohr) - 1
		}
		curLast = bohr[curLast].children[ch]

	}
	bohr[curLast].flag = true
	patterns = append(patterns, s)
	bohr[curLast].ix = len(patterns) - 1
}

// Возвращает суффиксную ссылку для узла v
func getSuffixIx(v int) int {
	if bohr[v].suffixIx == -1 {
		if v == 0 || bohr[v].par == 0 {
			bohr[v].suffixIx = 0
		} else {
			bohr[v].suffixIx = getAdvanceVertex(getSuffixIx(bohr[v].par), bohr[v].symbol)
		}
	}
	return bohr[v].suffixIx
}

// Возвращает автоматический переход для узла v и символа ch
func getAdvanceVertex(v int, ch byte) int {

	// return bohr[v].advanceVertexes[idx]
	idx := charToIndex(ch)
	if bohr[v].advanceVertexes[idx] != -1 {
		return bohr[v].advanceVertexes[idx]
	}

	// Если символ - джокер, обрабатываем его как любой символ алфавита
	if ch == joker {
		// Ищем первый доступный переход для джокера
		for i := 0; i < k; i++ {
			if bohr[v].children[i] != -1 {
				bohr[v].advanceVertexes[idx] = bohr[v].children[i]
				return bohr[v].advanceVertexes[idx]
			}
		}
		// Если нет переходов и мы находимся в корне, возвращаем корневой узел
		if v == 0 {
			bohr[v].advanceVertexes[idx] = 0
		} else {
			suffix := getSuffixIx(v)
			bohr[v].advanceVertexes[idx] = getAdvanceVertex(suffix, ch)
		}
		return bohr[v].advanceVertexes[idx]
	}

	// Если символ не джокер, проверяем обычный переход
	if bohr[v].children[idx] != -1 {
		bohr[v].advanceVertexes[idx] = bohr[v].children[idx]
		return bohr[v].advanceVertexes[idx]
	}

	// Если обычного перехода нет, проверяем переход по джокеру
	jokerIdx := charToIndex(joker)
	if bohr[v].children[jokerIdx] != -1 {
		bohr[v].advanceVertexes[idx] = bohr[v].children[jokerIdx]
		return bohr[v].advanceVertexes[idx]
	}

	// Если переходов нет, возвращаем переход из суффиксной ссылки
	if v == 0 {
		bohr[v].advanceVertexes[idx] = 0
	} else {
		suffix := getSuffixIx(v)
		bohr[v].advanceVertexes[idx] = getAdvanceVertex(suffix, ch)
	}

	return bohr[v].advanceVertexes[idx]
}

// Возвращает ссылку на ближайший терминальный узел
func getGoodSuffixIx(v int) int {
	if bohr[v].goodSuffixIx == -1 {
		u := getSuffixIx(v)
		if u == 0 {
			bohr[v].goodSuffixIx = 0
		} else {
			if bohr[u].flag {
				bohr[v].goodSuffixIx = u
			} else {
				bohr[v].goodSuffixIx = getGoodSuffixIx(u)
			}
		}
	}
	return bohr[v].goodSuffixIx
}

// Проверяет, является ли узел терминальным, и добавляет вхождения
func follow(v int, i int, occ *[]pair) {
	for u := v; u != 0; u = getGoodSuffixIx(u) {
		if bohr[u].flag {
			*occ = append(*occ, pair{i - len(patterns[bohr[u].ix]) + 1, bohr[u].ix + 1})
		}
	}
}

// Находит все вхождения образцов в строке s
func findAllOccurrences(s string) []pair {
	occ := []pair{}
	u := 0
	for i := 0; i < len(s); i++ {
		u = getAdvanceVertex(u, s[i])
		follow(u, i+1, &occ)
	}
	return occ
}

// func qcheck(occ []pair, s string) []pair {
// 	res := []pair{}
// 	for _, p := range occ {
// 		match := true
// 		for i := 0; i < len(patterns[p.ix-1]); i++ {
// 			pos := i + p.pos - 1
// 			if pos < len(s) &&
// 				s[pos] != patterns[p.ix-1][i] &&
// 				patterns[p.ix-1][i] != joker {
// 				match = false
// 				break
// 			}
// 		}
// 		if match {
// 			res = append(res, p)
// 		}
// 	}
// 	return res
// }

// Структура для хранения пар (позиция, номер образца)
type pair struct {
	pos int
	ix  int
}

// Выводит результаты
func outPairs(occ []pair, writer *bufio.Writer) {
	for _, p := range occ {
		fmt.Fprintln(writer, p.pos)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()
	bohrInit()
	// var S, tmp string
	// var n int

	// fmt.Scan(&S)
	// fmt.Scan(&n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&tmp)
	// 	addString(tmp)
	// }

	var t, buffer string
	fmt.Fscan(reader, &t)
	for {
		fmt.Fscan(reader, &buffer)
		if buffer == "$" {
			break
		}
		addString(buffer)
	}

	occ := findAllOccurrences(t)
	outPairs(occ, writer)
}
*/
