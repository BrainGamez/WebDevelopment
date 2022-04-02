package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	request(conn)
}

func request(conn net.Conn) {
	i := 0

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			url := strings.Fields(ln)[1]
			fmt.Println("*** METHOD:", m)
			fmt.Println("*** URI:", url)
			if m == "GET" && url == "/" {
				responseOK(conn)
			}
			if m == "GET" && url != "/" {
				responseNOTFound(conn)
			}
		}
		if ln == "" {
			//headers are done
			break
		}
		i++
	}
}

func responseOK(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8">
			 <title></title></head><body><strong>This was a GET Message (200)</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text.txt/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func responseNOTFound(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8">
			 <title></title></head><body><strong>This URL Cannot be found (404)</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text.txt/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
