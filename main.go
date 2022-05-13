package main

import(
  "fmt"
  "log"
  "net/http"
)

func main() {
  fileServer := http.fileServer(http.Dir("./static"))
  http.Handle("/",fileServer)
  http.HandleFunc("/form",formHandler)
  http.HandleFunc("/hello",formHandler)
  fmt.Printf("Iniciando server en el puerto 8080\n")
  if err := http.ListenAndServe(":8080",nil)
}
