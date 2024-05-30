package main

import "fmt"

const englishHelloPrefix = "Hello "
const spanish = "Spanish"
const spanishHelloPrefix = "Hola "
const french = "French"
const frenchHelloPrefix = "Bonjour "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("", ""))
}
