package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handler(conn net.Conn) {
	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	i := 0

	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			fmt.Fprintf(conn, "I heard you say: %s\n", ln)
		}
		if ln == "" {
			break
		}
		i++
	}

	body := `<!DOCTYPE html><html lang="en"><head><title>TCP Server</title></head><body>
	<h1>hello world</h1></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)

	fmt.Println("Code got here.")
}

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()

	for {
		conn, err := li.Accept()

		if err != nil {
			log.Println(err)
		}

		go handler(conn)
	}
}
