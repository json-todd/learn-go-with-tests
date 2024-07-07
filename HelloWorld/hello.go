package main
import "fmt"

const englishHelloPrefix = "Hello, "
const finnishHelloPrefix = "Moi, "

func Hello(name string) string {
    return englishHelloPrefix + name
}

func main() {
    fmt.Println(Hello("world"))
}
