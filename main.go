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
    port, exists := os.LookupEnv("GOLANG_PORT")

    if exists {
        http.HandleFunc("/hello", helloHandler)
        http.ListenAndServe(':'+port, nil)
    	fmt.Print("Port = "+port)
   }
}