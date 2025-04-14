package libs

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

// Структура для управления вводом/выводом
type InputOutputController struct {
	writer *bufio.Writer // указатель на структуру для вывода данных
	reader *bufio.Reader // указатель на структуру для чтения данных
}

// Создаёт новый экземпляр структуры управления вводом/выводом
func NewIoController() *InputOutputController {
	// возвращает указатель на созданных экземпляр структуры
	return &InputOutputController{
		// инициализация полей
		writer: bufio.NewWriter(os.Stdout),
		reader: bufio.NewReader(os.Stdin),
	}
}

// Считывает одно число из потока, предполаагается корректный ввод
func (ioC *InputOutputController) ReadNum() int {
	line, _, _ := ioC.reader.ReadLine()
	num, _ := strconv.Atoi(string(line))

	return num
}

// считывает строку из чисел и возвращает массив из этих чисел
func (ioC *InputOutputController) ReadLineOfInts() []int {
	line, _ := ioC.reader.ReadString('\n')
	return splitLineOfInt(line)
}

// преобразует строку чисел в массив чисел
func splitLineOfInt(input string) []int {

	fields := bytes.Fields([]byte(input))
	numbers := make([]int, len(fields))
	for i, field := range fields {
		numbers[i], _ = strconv.Atoi(string(field))
	}
	return numbers
}

// выводит данные из переменных
func (ioC *InputOutputController) Writeln(data ...any) {
	fmt.Fprintln(ioC.writer, data...)
	ioC.writer.Flush()
}

// вывод данных из переменных без переноса строки
func (ioC *InputOutputController) Write(data ...any) {
	fmt.Fprint(ioC.writer, data...)
	ioC.writer.Flush()
}
