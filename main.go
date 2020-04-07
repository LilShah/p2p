//better to build before running with `go build main.go` then `./main.exe`
package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func dialOthers(c net.Conn, port string) {

	conn, err := net.Dial("tcp", "localhost:2020")
	if err != nil {
		// handle error
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

}

func main() {
	args := os.Args[1:]
	if len(args) == 1 {
		ln, err := net.Listen("tcp", ":"+args[0])
		if err != nil {
			log.Fatal(err)
		}
		log.Println(ln)
	} else if len(args) > 1 {
		ln, err := net.Listen("tcp", ":"+args[len(args)-1]) //listen to last arg
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < len(args)-1; i++ {
			conn, err := ln.Accept()
			if err != nil {
				log.Println(err)
				continue
			}
			go dialOthers(conn, args[i])
		}
	}
}
