package main
import "fmt"

const (
    finnish = "Finnish"
    german = "German"

    englishHelloPrefix = "Hello, "
    finnishHelloPrefix = "Moi, "
    germanHelloPrefix = "Hallo, "
)

func Hello(name, language string) string {
    if name == "" {
        name = "World"
    }

    return greetingPrefix(language) + name

}

func greetingPrefix(language string) (helloPrefix string) {
    switch language {
      case finnish:
          helloPrefix = finnishHelloPrefix
      case german:
          helloPrefix = germanHelloPrefix
      default:
          helloPrefix = englishHelloPrefix
    }
    return
}

func main() {
    fmt.Println(Hello("",""))
}
