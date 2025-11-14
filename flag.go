package main

import "flag"

// Package flag pada golang digunakan untuk memparsing argumen command line yang diberikan saat menjalankan program
func main() {
	username := flag.String("username", "guest", "the username for login")
	password := flag.String("password", "", "the password for login")
	host := flag.String("host", "localhost", "the host address")
	port := flag.Int("port", 8080, "the port number")

	flag.Parse()

	println("Username:", *username)
	println("Password:", *password)
	println("Host:", *host)
	println("Port:", *port)
}
