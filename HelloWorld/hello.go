package main
import "fmt"

const finnish = "Finnish"
const german = "German"
const englishHelloPrefix = "Hello, "
const finnishHelloPrefix = "Moi, "
const germanHelloPrefix = "Hallo, "

func Hello(name, language string) string {
    if name == "" {
        name = "World"
    }

    // default to english
    helloPrefix := englishHelloPrefix

    switch language {
      case finnish:
          helloPrefix = finnishHelloPrefix
      case german:
          helloPrefix = germanHelloPrefix
    }
    return helloPrefix + name
}

func main() {
    fmt.Println(Hello("",""))
}
