package main

import "strings"

func createHugeString(size int) string {
	return strings.Repeat("a", size)
}

// глобальная переменная
// живет в памяти до тех пор, пока программа жива
var justString string

func someFunc() {
	// создаем большую строку (размером 1024)
	v := createHugeString(1 << 10)
	// теперь justString ссылается на тот же массив байтов, что и v
	justString = v[:100]
	// после окончания работы someFunc() в памяти останется весь массив 1024 байтов,
	// так как justString ссылается на него и cборщик мусора не сможет удалить его
	// пока жива переменная justString

	// правильно будет сделать сделать копию
	// justString = string([]byte(v[:100]))
	// или
	// justString = strings.Clone(v[:100])
}

func main() {
	someFunc()
}
