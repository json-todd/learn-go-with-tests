package main
import "fmt"

const finnish = "Finnish"
const germany = "Germany"
const englishHelloPrefix = "Hello, "
const finnishHelloPrefix = "Moi, "
const germanyHelloPrefix = "Hallo, "

func Hello(name, language string) string {
    if name == "" {
        name = "World"
    }

    if language == finnish {
        return finnishHelloPrefix + name
    }

    if language == germany {
        return germanyHelloPrefix + name
    }
    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("",""))
}
