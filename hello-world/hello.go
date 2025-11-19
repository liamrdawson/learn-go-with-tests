package main

import "fmt"

const (
	welsh  = "Welsh"
	french = "French"

	englishHelloPrefix = "Hello, "
	welshHelloPrefix   = "Helo, "
	frenchHelloPrefix  = "Bonjour, "
)

func greetingPrefix(language string) (prefix string) {

	switch language {
	case welsh:
		prefix = welshHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return
}

func Hello(language, name string) string {

	if name == "" {
		if language == welsh {
			name = "Byd"
		} else {
			name = "World"
		}
	}

	return greetingPrefix(language) + name
}

func main() {
	fmt.Println(Hello("", "Liam"))
}
