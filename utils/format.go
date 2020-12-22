package utils

import (
  "fmt"
  "net/http"
  "encoding/json"
)

func Respond(w http.ResponseWriter, status int, message string) {
  w.WriteHeader(status)
  fmt.Fprintf(w, message)
}

func RespondJSON(w http.ResponseWriter, status int, payload interface{}) {
  response, err := json.Marshal(payload)
  if err != nil {
    http.Error(w, err.Error(), 500)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
  w.Write([]byte(response))
}
