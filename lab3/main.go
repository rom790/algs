package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"sort"
	"strings"
	"sync"
)

const (
	k = 5 // Количество символов в алфавите: {A, C, G, T, N}
)

type TrieVrtx struct {
	NextVrtx  [k]int `json:"next_vrtx"`  // Массив переходов по символам (ASCII)
	PatNum    int    `json:"pat_num"`    // Номер образца
	SuffLink  int    `json:"suff_link"`  // Суффиксная ссылка
	AutoMove  [k]int `json:"auto_move"`  // Автоматические переходы
	Par       int    `json:"par"`        // Родительский узел
	SuffFLink int    `json:"suff_flink"` // Ссылка на терминальную вершину
	Flag      bool   `json:"flag"`       // Флаг терминальной вершины
	Parr      byte   `json:"parr"`       // Символ перехода
}

type SearchStep struct {
	CurrentState int    `json:"current_state"`           // индекс текущего состояние поиска (текущая вершина бора)
	Symbol       byte   `json:"symbol"`                  // текущий символ
	NextState    int    `json:"next_state"`              // индекс следующей вершины бора
	FoundPattern string `json:"found_pattern,omitempty"` // содержит найденную подстроку
}

type SearchResponse struct {
	Steps      []SearchStep `json:"steps"`       // массив шагов поиска
	SearchText string       `json:"search_text"` // исследуемый текст
}

type TrieResponse struct {
	Steps [][]TrieVrtx `json:"steps"` // массив шагов построения бора
}

type ParallelSearchResponse struct {
	Threads []SearchResponse `json:"threads"` // массив результатов работы потоков
}

var (
	// переменные для хранения данных, используемых при визуализации
	TrieSteps       TrieResponse
	searchSteps     SearchResponse
	parallelVisData ParallelSearchResponse
	//===================================
	trie    []TrieVrtx // Бор (префиксное дерево)
	pattern []string   // Список образцов
)

// charToIndex преобразует символ в индекс (A -> 0, C -> 1, G -> 2, T -> 3, N -> 4)
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
		return -1
	}
}

// makeTrieVrtx создает новый узел бора
func makeTrieVrtx(p int, c byte) TrieVrtx {
	v := TrieVrtx{}
	for i := 0; i < k; i++ {
		v.NextVrtx[i] = -1
		v.AutoMove[i] = -1
	}
	v.Flag = false
	v.SuffLink = -1
	v.Par = p
	v.Parr = c
	v.SuffFLink = -1
	return v
}

// trieInit инициализирует бор, добавляя корневой узел
func trieInit() {
	trie = append(trie, makeTrieVrtx(0, '$'))
}

// addStringToTrie добавляет строку в бор
func addStringToTrie(s string, vis bool) {
	num := 0 // Начинаем с корневого узла
	for i := 0; i < len(s); i++ {
		ch := charToIndex(s[i])           // Преобразуем символ в индекс
		if trie[num].NextVrtx[ch] == -1 { // Если переход по символу отсутствует
			trie = append(trie, makeTrieVrtx(num, s[i])) // Создаем новый узел
			trie[num].NextVrtx[ch] = len(trie) - 1       // Добавляем переход
			if vis {
				TrieSteps.Steps = append(TrieSteps.Steps, trie)
			}
		}
		num = trie[num].NextVrtx[ch] // Переходим к следующему узлу
	}
	trie[num].Flag = true               // Помечаем узел как конец образца (обошли все символы строки)
	pattern = append(pattern, s)        // Добавляем образец в список
	trie[num].PatNum = len(pattern) - 1 // Сохраняем номер образца
	// TrieSteps.Steps = append(TrieSteps.Steps, trie)
	if vis {
		TrieSteps.Steps[len(TrieSteps.Steps)-1] = trie
	}
}

// Возвращает суффиксную ссылку для узла v (номер уровня)
func getSuffLink(v int) int {
	if trie[v].SuffLink == -1 { // Если суффиксная ссылка еще не вычислена
		if v == 0 || trie[v].Par == 0 { // Если узел корневой или его родитель корневой
			trie[v].SuffLink = 0
		} else {
			trie[v].SuffLink = getAutoMove(getSuffLink(trie[v].Par), trie[v].Parr)
		}
	}
	return trie[v].SuffLink
}

// возвращает переход автомата для узла v и символа ch
func getAutoMove(v int, ch byte) int {
	idx := charToIndex(ch)           // Преобразуем символ в индекс
	if trie[v].AutoMove[idx] == -1 { // Если автоматический переход еще не вычислен
		if trie[v].NextVrtx[idx] != -1 { // Если есть прямой переход
			trie[v].AutoMove[idx] = trie[v].NextVrtx[idx]
		} else {
			if v == 0 { // Если узел корневой
				trie[v].AutoMove[idx] = 0
			} else {
				trie[v].AutoMove[idx] = getAutoMove(getSuffLink(v), ch)
			}
		}
	}
	return trie[v].AutoMove[idx]
}

// Возвращает суффиксную ссылку на терминальную вершину (fast link)
func getSuffFlink(v int) int {
	if trie[v].SuffFLink == -1 { // Если суффиксная ссылка еще не вычислена
		u := getSuffLink(v)
		if u == 0 { // Если суффиксная ссылка ведет в корень
			trie[v].SuffFLink = 0
		} else {
			if trie[u].Flag { // Если узел является концом образца
				trie[v].SuffFLink = u
			} else {
				trie[v].SuffFLink = getSuffFlink(u) // двигаемся вверх по дереву
			}
		}
	}
	return trie[v].SuffFLink
}

// check проверяет, является ли узел концом образца, и выводит результат
func check(v int, i int, res *[][]int) []string {
	var patterns []string
	for u := v; u != 0; u = getSuffFlink(u) {
		if trie[u].Flag { // Если узел является концом образца
			pos := i - len(pattern[trie[u].PatNum]) + 1
			patNum := trie[u].PatNum + 1
			*res = append(*res, []int{pos, patNum})
			patterns = append(patterns, pattern[trie[u].PatNum])
		}
	}
	return patterns
}

// findAllPos ищет все вхождения образцов в строке s
func findAllPos(s string, vis bool) [][]int {
	var res [][]int
	thread := SearchResponse{
		SearchText: s,
		Steps:      make([]SearchStep, 0),
	}
	u := 0 // Начинаем с корневого узла
	for i := 0; i < len(s); i++ {
		prevState := u
		u = getAutoMove(u, s[i]) // Переходим по автоматической ссылке
		step := SearchStep{
			CurrentState: prevState,
			Symbol:       s[i],
			NextState:    u,
		}
		found := check(u, i+1, &res) // Проверяем, является ли узел концом образца
		if len(found) > 0 && vis {
			step.FoundPattern = strings.Join(found, ", ")
		}
		if vis {
			thread.Steps = append(thread.Steps, step)
		}
	}

	parallelVisData.Threads = append(parallelVisData.Threads, thread)
	return res
}

// разбивает строку по заданному разделителю на массив подстрок
func splitWithPositions(input string, delimiter rune) ([]string, []int) {
	var substrings []string
	var positions []int
	var buffer []rune

	ind := 0
	for _, r := range input {
		if r == delimiter {
			// Если текущий символ — разделитель, добавляем буфер в результат
			if len(buffer) > 0 {
				substrings = append(substrings, string(buffer))
				positions = append(positions, ind-len(buffer)+1)
				buffer = []rune{} // Очищаем буфер
			}
		} else {
			// Иначе добавляем символ в буфер
			buffer = append(buffer, r)
		}
		ind += 1
	}
	// Добавляем последний буфер, если он не пустой
	if len(buffer) > 0 {
		substrings = append(substrings, string(buffer))
		positions = append(positions, ind-len(buffer)+1)
	}
	return substrings, positions
}

func main() {
	numFlag := flag.Int("n", 1, "Номер используемой функции")
	threadsFlag := flag.Int("k", 1, "Количество используемых потоков")
	visFlag := flag.Bool("vis", false, "Использование визуализации")
	debugFlag := flag.Bool("d", false, "Использовать вывод работы потоков в консоль")

	flag.Parse()

	if *numFlag != 1 && *numFlag != 2 {
		fmt.Println("Неверный аргумент флага запуска -n:", *numFlag)
		return
	}
	if *threadsFlag <= 0 {
		fmt.Println("Неверный аргумент флага запуска -k:", *threadsFlag)
		return
	}

	trieInit() // Инициализация бора
	var t string
	fmt.Scan(&t)
	if *numFlag == 1 {
		var buffer string
		var n int
		fmt.Scan(&n)
		for i := 0; i < n; i++ {
			fmt.Scan(&buffer)
			addStringToTrie(buffer, *visFlag)
		}

		var res [][]int
		if *threadsFlag > 1 {
			maxPLen := 0
			for _, p := range pattern {
				if len(p) > maxPLen {
					maxPLen = len(p)
				}
			}
			if *threadsFlag > len(t) {
				fmt.Println("Количество потоков должно не превышать длину текста")
				return
			} else if (len(t) / *threadsFlag) < maxPLen {
				fmt.Println("Количество потоков слишком велико или длины шаблонов слишком большие")
				return
			}
			res = parallelFindAllPos(t, *threadsFlag, *visFlag, *debugFlag)
		} else {
			res = findAllPos(t, *visFlag)
		}
		sort.Slice(res, func(i, j int) bool {
			if res[i][0] < res[j][0] {
				return true
			}
			if res[i][0] == res[j][0] {
				return res[i][1] < res[j][1]
			}
			return false
		})

		for i := 0; i < len(res); i++ {
			fmt.Println(res[i][0], res[i][1])
		}

	} else {
		var buffer, jocker string

		fmt.Scan(&buffer, &jocker)

		bufferParts, startPositions := splitWithPositions(buffer, rune(jocker[0]))
		PartToPos := make(map[string][]int)
		for i, Part := range bufferParts {
			addStringToTrie(Part, *visFlag)
			PartToPos[Part] = append(PartToPos[Part], startPositions[i])
		}

		c := make([]int, len(t)+1)

		result := findAllPos(t, *visFlag)

		for _, r := range result {
			for _, ind := range PartToPos[bufferParts[r[1]-1]] {
				if r[0]-ind >= 0 {
					c[r[0]-ind] += 1
				}
			}
		}
		for i := 0; i < len(c)-len(buffer); i++ {
			if c[i] == len(bufferParts) {
				fmt.Println(i + 1)
			}
		}
	}
	if *visFlag {
		searchSteps.SearchText = t
	}

	if *visFlag {
		fmt.Println("Запущено на localhost:8080")
		http.HandleFunc("/search-data", searchHandler)
		http.HandleFunc("/trie-data", trieDataHandler)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "static/index.html")
		})

		http.ListenAndServe(":8080", nil)
	}
}

// записывает данные о процессе построения бора для отправки клиенту
func trieDataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(TrieSteps)
}

// записывает данные о процессе поиска для отправки клиенту
func searchHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(parallelVisData)
}

// Ищет все вхождения шаблонов, используя k потоков
func parallelFindAllPos(s string, k int, vis bool, debug bool) [][]int {

	parallelVisData.Threads = make([]SearchResponse, 0, k)

	maxPatternLen := 0
	for _, p := range pattern {
		if len(p) > maxPatternLen {
			maxPatternLen = len(p)
		}
	}

	chunks := splitTextWithOverlap(s, k, maxPatternLen)
	overlap := maxPatternLen - 1

	type position struct {
		pos    int
		patNum int
	}

	var wg sync.WaitGroup
	results := make(chan []position, k) // Новый тип для хранения позиций
	var finalResult [][]int

	// Запускаем горутины для обработки каждой части
	for j, chunk := range chunks {
		wg.Add(1)
		go func(chunk string, startPos int, chunkIndex int) {
			defer wg.Done()

			if debug {
				fmt.Printf("Поток %d начал обработку фрагмента текста %d (позиции %d - %d)",
					chunkIndex, chunkIndex+1, startPos, startPos+len(chunk))
			}

			threadVis := SearchResponse{
				SearchText: chunk,
			}

			var localRes []position
			u := 0

			for i := 0; i < len(chunk); i++ {
				step := SearchStep{
					CurrentState: u,
					Symbol:       chunk[i],
					NextState:    getAutoMove(u, chunk[i]),
				}
				u = getAutoMove(u, chunk[i])
				if debug {
					fmt.Printf("Поток %d: обработал %d/%d символов (%.1f%%), состояние %d->%d\n",
						chunkIndex, i+1, len(chunk), float64(i+1)/float64(len(chunk))*100, step.CurrentState, u)
				}
				absolutePos := startPos + i + 1

				// Временный сбор результатов
				var tmpRes [][]int
				check(u, absolutePos, &tmpRes)
				if len(tmpRes) > 0 {
					step.FoundPattern = pattern[tmpRes[0][1]-1]
					// Конвертируем во внутренний формат
					for _, r := range tmpRes {
						// Фильтруем дубликаты на границах чанков
						if chunkIndex == 0 || r[0] >= chunks[chunkIndex].start+overlap {
							localRes = append(localRes, position{r[0], r[1]})
						}
					}
				}
				if vis {
					threadVis.Steps = append(threadVis.Steps, step)
				}
			}
			if vis {
				parallelVisData.Threads = append(parallelVisData.Threads, threadVis)
			}

			results <- localRes
		}(chunk.text, chunk.start, j)
	}

	// Собираем и дедуплицируем результаты
	seen := make(map[position]bool)
	go func() {
		wg.Wait()
		close(results)
		if debug {
			fmt.Println("Все потоки завершили работу")
		}
	}()

	for res := range results {
		for _, r := range res {
			if !seen[r] {
				seen[r] = true
				finalResult = append(finalResult, []int{r.pos, r.patNum})
			}
		}
	}

	return finalResult
}

// разбивает строку на подстроки по заданному разделителю
func splitTextWithOverlap(text string, k int, maxPatternLen int) []struct {
	text  string
	start int
} {
	textLen := len(text)
	chunkSize := (textLen + k - 1) / k // Округляем вверх
	overlap := maxPatternLen - 1

	var chunks []struct {
		text  string
		start int
	}

	for i := 0; i < k; i++ {
		start := i * chunkSize
		if start >= textLen {
			break
		}

		end := start + chunkSize
		if i < k-1 {
			end += overlap
		}
		if end > textLen {
			end = textLen
		}

		// Для всех частей кроме первой, начинаем с перекрытия
		chunkStart := start
		if i > 0 {
			chunkStart = start - overlap
			if chunkStart < 0 {
				chunkStart = 0
			}
		}

		chunks = append(chunks, struct {
			text  string
			start int
		}{
			text:  text[chunkStart:end],
			start: chunkStart,
		})
	}

	return chunks
}
