package main

import "fmt"

const (
	englishHelloPrefix = "Hello "
	spanishHelloPrefix = "Hola "
	koreanHelloPrefix  = "안녕하세요 "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}
	return getGreetings(language) + name
}

func getGreetings(language string) (prefix string) {
	switch language {
	case "korean":
		prefix = koreanHelloPrefix
	case "spanish":
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Joe", ""))
}
