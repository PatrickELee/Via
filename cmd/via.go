package main

import "github.com/PatrickELee/Via/internal/server"

func main() {
	vs := server.CreateViaWebServers()
	vs.ListenAndServe()
}
