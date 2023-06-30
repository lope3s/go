package main

import "fmt"

const (
    spanish               = "Spanish"
    french                = "French"
    portuguese            = "Portuguese"

	englishHelloPrefix    = "Hello, "
	spanishHelloPrefix    = "Holla, "
	frenchHelloPrefix     = "Bonjour, "
	portugueseHelloPrefix = "Olá, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

    return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
    default:
        prefix = englishHelloPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", "Spanish"))
}
