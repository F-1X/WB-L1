package main

import (
	"fmt"
)

// Реализовать паттерн «адаптер» на любом примере.
func main() {
	// допустим у нас китайская вилка, и нам нужен переходник на европейскую розетку.

	// китайская вилка
	chinesePlug := &ChinaPlug{}

	// адаптер китайской вилки (переходник на европейскую розетку)
	adapter := NewEuroAdapter()
	adapter.ConnectToAdapter(chinesePlug) // также возможен вариант подсоединения через конструктор

	// создаем розетку
	euroSocket := &EuroSocket{}

	// используем адаптер как европейскую розетку
	euroSocket.PlugIn(adapter.PlugIn()) // adapter.PlugIn() - для наглядности использования из-за вывода в консоль
	euroSocket.PlugIn(adapter) // композиция
	
	// Chinese fork connected to adapter
	// Using European plug
	// Powering with European socket

}

type EuroAdapterer interface {
	ConnectToAdapter()
}

// EuroAdapter - адаптер, позволяющий использовать китайскую вилку с европейской розеткой
type EuroAdapter struct {
	euroPlug    EuroPlug
	chinesePlug ChinesePlug
}

func NewEuroAdapter() *EuroAdapter {
	return &EuroAdapter{
		euroPlug: EuroPlug{}, // создается новая евро вилка из переходника
		//chinesePlug: nil, // китайская вилка встроена в переходник
	}
}

func (e *EuroAdapter) ConnectToAdapter(chinesePlug ChinesePlug) {
	fmt.Println("Chinese fork connected to adapter")
	e.chinesePlug = chinesePlug
}

func (e *EuroAdapter) PlugIn() EuroPlug {
	return e.euroPlug.PlugIn()
}

// ChinesePlug - интерфейс китайской вилки
type ChinesePlug interface {
	PlugIn()
}

// ChinaPlug - китайская вилки
type ChinaPlug struct{}

func (c *ChinaPlug) PlugIn() {
	fmt.Println("Using Chinese plug")
}

// EuropeanSocket - возможности розетки
type EuropeanSocket interface {
	PlugIn(plug EuroPlugger) // подключиться к розетке могут только евро вилки
}

// EuroSocket - евро розетка
type EuroSocket struct{}

// Подключиться к розетке
func (e *EuroSocket) PlugIn(plug EuroPlugger) {
	fmt.Println("Powering with European socket")
}

// Возможности евро вилок
type EuroPlugger interface {
	PlugIn() EuroPlug
}

// евро вилка
type EuroPlug struct{}

// вставить евро вилку
func (e EuroPlug) PlugIn() EuroPlug {
	fmt.Println("Using European plug")
	return e
}
