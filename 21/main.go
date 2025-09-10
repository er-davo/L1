package main

import "fmt"

// Target - интерфейс, который ожидает клиент
type Logger interface {
	Log(msg string, args ...interface{})
}

// Adaptee — существующий код с несовместимым интерфейсом
type SomeOtherLogger struct{}

func (l *SomeOtherLogger) Print(msg string) {
	fmt.Println(msg)
}

// Adapter — реализует интерфейс Logger и
// внутри использует SomeOtherLogger
type LoggerAdapter struct {
	som *SomeOtherLogger
}

func (l *LoggerAdapter) Log(msg string, args ...interface{}) {
	l.som.Print(fmt.Sprintf(msg, args...))
}

// Client — работает только с интерфейсом Logger
func Client(l Logger) {
	l.Log("Logging message %s", "Hello World!")
}

func main() {
	som := &SomeOtherLogger{}
	somAdapter := &LoggerAdapter{som: som}

	Client(somAdapter)
}
