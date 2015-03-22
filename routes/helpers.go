package routes

import (
  "strconv"

  "github.com/erkl/robo"
)

func send(w robo.ResponseWriter, mime string, buf []byte) {
  h := w.Header()
  h.Set("Content-Type", mime)
  h.Set("Content-Length", strconv.Itoa(len(buf)))
  w.WriteHeader(200)
  w.Write(buf)
}