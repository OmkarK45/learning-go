package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

// Says hello to the user
func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return prefix
}

func main() {
	fmt.Println(Hello("world", "English"))
}
