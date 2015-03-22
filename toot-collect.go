package main

import (
  "log"
  "net/http"
  "os"

  "github.com/erkl/robo"

  "./db"
  "./poller"
  "./routes"
)

func main() {
  db.Connect(os.Getenv("MONGO"))

  // Bootstrap existing sites
  poller.Bootstrap()

  mux := new(robo.Mux)
  routes.Attach(mux)

  log.Printf("Listening on %s", os.Getenv("HOST")+":"+os.Getenv("PORT"))
  log.Printf("Connecting to mongo on %s", os.Getenv("MONGO"))
  err := http.ListenAndServe(os.Getenv("HOST")+":"+os.Getenv("PORT"), mux)
  if err != nil {
    log.Fatal("http.ListenAndServe: %s", err)
  }
}
