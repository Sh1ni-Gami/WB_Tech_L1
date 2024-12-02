package main

import "fmt"

// Интерфейс, который ожидает клиент
type TempFahrenheit interface {
	GetFahrenheit() float64
}

// структура Цельсиях
type TempCelsius struct {
	celsius float64
}

// Метод для получения температуры в Цельсиях
func (c *TempCelsius) GetCelsius() float64 {
	return c.celsius
}

// Адаптер для преобразования Цельсиев в Фаренгейты
type TempAdapter struct {
	celsiusTemp *TempCelsius
}

// Реализация интерфейса в адаптере
func (adapter *TempAdapter) GetFahrenheit() float64 {
	// Формула перевода: (C × 9/5) + 32
	celsius := adapter.celsiusTemp.GetCelsius()
	fahrenheit := (celsius * 9 / 5) + 32
	return fahrenheit
}

// Функция для печати температуры в Фаренгейтах
func printTemperature(temp TempFahrenheit) {
	fmt.Printf("Температура в Фаренгейтах: %.2f°F\n", temp.GetFahrenheit())
}

func main() {
	celsiusTemp := &TempCelsius{celsius: 25}
	fmt.Printf("Температура в Цельсиях: %.2f°C\n", celsiusTemp.GetCelsius())

	adapter := &TempAdapter{celsiusTemp: celsiusTemp}
	printTemperature(adapter)

}
