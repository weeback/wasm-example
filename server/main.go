package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	addrs = "0.0.0.0:3000"
)

func init() {
	publicPort := os.Getenv("PORT")
	if len(publicPort) > 2 || publicPort == "80" {
		addrs = fmt.Sprintf("0.0.0.0:%s", publicPort)
	}
}

func main() {
	selectDir := func(path string) http.Handler {
		dir := http.Dir(path)
		return http.FileServer(dir)
	}
	defer func() {
		if err := http.ListenAndServe(addrs, selectDir("application")); err != nil {
			log.Printf("%+v", struct {
				Level  string `json:"level"`
				Method string `json:"method"`
				Addrs  string `json:"addrs"`
				Error  error  `json:"error"`
			}{
				Level:  "DEBUG",
				Method: "ListenAndServe",
				Addrs:  addrs,
				Error:  err,
			})
		}
		println("Server closed!")
	}()

	println("Start the server on", addrs, "...")
}
