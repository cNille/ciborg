package main
import (
  "net/http"
  "strings"
  "log"
)

//

func githubWebhook(w http.ResponseWriter, r *http.Request) {
  // Log path
  log.Println("Path visited: " + r.URL.Path)

  // Log body
  r.ParseForm()
  log.Println(r.Form)

  // Echo the path
  message := r.URL.Path
  message = strings.TrimPrefix(message, "/")
  message = "Hello " + message
  w.Write([]byte(message))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
  // Log path
  log.Println("Path visited: " + r.URL.Path)
  message := "Hello, my name is ciBorg and im your newbie CI server."
  w.Write([]byte(message))
}

func main() {
  // Handle updates from github
  http.HandleFunc("/github/webhook", githubWebhook)

  http.HandleFunc("/", sayHello)

  // http.HandleFunc("/", githubWebhook)
  if err := http.ListenAndServe(":9876", nil); err != nil {
    panic(err)
  }
}
