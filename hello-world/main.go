package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola "
const french = "French"
const frenchPrefix = "Bonjour "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == spanish {
		return spanishHelloPrefix + name
	}
	if language == french {
		return frenchPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
