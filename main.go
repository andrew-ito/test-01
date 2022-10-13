package main
import (
        "net/http"
        "io"
        "os"
        "fmt"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
        io.WriteString(w, "Hello world!")
}


func main() {
    port, exists := os.LookupEnv("golang_port")

    if exists {
        fmt.Print("Port = "+port)
        http.HandleFunc("/hello", helloHandler)
        http.ListenAndServe(":"+port, nil)
    } else {
        fmt.Println("No any port in env variables.")
        fmt.Println("Port will be 8080")
        http.HandleFunc("/hello", helloHandler)
        http.ListenAndServe(":8080", nil)
    }
}
