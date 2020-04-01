package main

import (
	employee "github.com/saurabhagg301/employee/server"
)

func main() {
	host := "0.0.0.0"
	port := 8080

	employee.StartWebServer(host, port)
}
