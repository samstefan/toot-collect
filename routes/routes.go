package routes

import (
  "github.com/erkl/robo"
)

// Attach adds all routes to a robo.Mux instance.
func Attach(mux *robo.Mux) {

  // GET /tweet?account="test"
  mux.Add("GET", "/tweets/{screenName}", getTweet)

}