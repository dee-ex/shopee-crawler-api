package utils

import (
  "strconv"
  "net/http"
)

func GetNumericQuery(r *http.Request, key string, _default int) int {
  value_query, ok := r.URL.Query()[key]
  if !ok {
    return _default
  }
  value, err := strconv.Atoi(value_query[0])
  if err != nil {
    return _default
  }
  return value
}
