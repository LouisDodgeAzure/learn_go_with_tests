package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const australian = "Australian"
const englishHelloPrefix = "Hello, "
const spanishHelloPrexix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const australianHelloPrefix = "G'day, "

func Hello(language, name string) string {
	if name == "" {
		name = "World"
	}
	helloPrefix := greetingsPrefix(language)
	return helloPrefix + name + "!"
}

func greetingsPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrexix
	case australian:
		prefix = australianHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	g := Hello(englishHelloPrefix, "Louis")
	fmt.Println(g)
}
