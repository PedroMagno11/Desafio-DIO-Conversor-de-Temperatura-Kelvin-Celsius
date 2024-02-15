package main

import "fmt"

func main() {
	menu()
}

func menu() {
	fmt.Println("Welcome to the temperature converter")
	tempCelsius := converter(enterTemp())
	fmt.Printf("This temperature is equivalent to %g degrees Celsius", tempCelsius)
}

func enterTemp() float32 {
	var tempKelvin float32
	fmt.Print("Enter a temperature in Kelvin: ")
	fmt.Scan(&tempKelvin)
	return tempKelvin
}

func converter(temp float32) float32 {
	return temp - 273.15
}
