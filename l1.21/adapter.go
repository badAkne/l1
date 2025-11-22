package main

import "fmt"

// Интерфейс класса адаптера
type Target interface {
	Operation()
}

// адаптируемый класс
type Adaptee struct{}

// Метод адаптируемого класса
func (a *Adaptee) AdaptedOperation() { fmt.Println("I'am a adapted operation()") }

// класс конкретного адаптера
type ConcreteAdapter struct{ *Adaptee }

// реализация метода интерфейса, реализующего вызов адаптируемого класса
func (a *ConcreteAdapter) Operation() { a.AdaptedOperation() }

// конструктор адаптера
func NewAdapter(a *Adaptee) Target {
	return &ConcreteAdapter{Adaptee: &Adaptee{}}
}

func main() {
	adapter := NewAdapter(&Adaptee{})
	adapter.Operation()
}

/*
	Применимость:
1. Если у нас есть фреймфорк с плохим API
2. Создание унифицированного кода интерфейса для разнородныхсистем

*/

/*
	Плюсы:
	1. Переиспользование старого кода, то поддержка принципу DRY
	2. Легко тестировать
*/

/*
	Минусы:
	1. Усложнения кода из-за большей абстракции
	2. Уменьшения производительности из-за допю слоя
*/

/*
	Реальные способы применения:
	1. Допустим у нас есть сервис, который отправляет уведомления через смс, через некоторое время мы хотим отправлять только PUSH уведомления, тогда можно использовать адптер
	2. Адаптация внешних API к внутрениим интерфейсам
*/
