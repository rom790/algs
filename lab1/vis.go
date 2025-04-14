package main

import (
	// "fmt"
	libs "lab1/libs"
)

// ANSI escape-коды для цветов, соответствующих размерам квадратов
var sizeToColor = map[int]string{
	1:  "\033[31m",       // Красный
	2:  "\033[32m",       // Зеленый
	3:  "\033[33m",       // Желтый
	4:  "\033[34m",       // Синий
	5:  "\033[35m",       // Пурпурный
	6:  "\033[36m",       // Голубой
	7:  "\033[37m",       // Белый
	8:  "\033[90m",       // Темно-серый
	9:  "\033[38;5;208m", // Оранжевый
	10: "\033[38;5;46m",  // Ярко-зеленый
	11: "\033[38;5;27m",  // Глубокий синий
	12: "\033[38;5;201m", // Розовый
	13: "\033[38;5;220m", // Золотистый
	14: "\033[38;5;51m",  // Бирюзовый
	15: "\033[38;5;202m", // Темно-оранжевый
	16: "\033[38;5;129m", // Фиолетовый
	17: "\033[38;5;82m",  // Лаймовый
	18: "\033[38;5;226m", // Лимонный
	19: "\033[38;5;196m", // Алый
}

// Сброс цвета
const resetColor = "\033[0m"

// выбирает цвет, в зависимости от размера квадраата
func getColorForSize(size int) string {
	if color, exists := sizeToColor[size]; exists {
		return color
	}
	return "\033[97m" // Белый по умолчанию
}

// выводит поле с размещёнными на нём квадратами
func Show(squares [][]int, fieldSize int) {
	// структура ввода-вывод
	ioContr := libs.NewIoController()

	// Создаём пустое поле
	field := make([][]string, fieldSize)
	for i := range field {
		field[i] = make([]string, fieldSize)
		for j := range field[i] {
			field[i][j] = "." // Заполняем пустыми точками
		}
	}

	// Размещаем квадраты
	for _, square := range squares {
		// получаем координаты и размер квадратат
		x, y, size := square[0], square[1], square[2]
		color := getColorForSize(size)
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				if y+i < fieldSize && x+j < fieldSize {
					field[x+i][y+j] = color + "■" + resetColor
				}
			}
		}
	}

	// Выводим поле в консоль
	for _, row := range field {
		for _, cell := range row {
			ioContr.Write(cell)
			ioContr.Write(" ")
		}
		ioContr.Writeln()
	}
}
