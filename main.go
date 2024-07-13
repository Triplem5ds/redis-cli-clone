package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	l, err := net.Listen("tcp", ":6379")

	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()
	defer conn.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		buff := make([]byte, 1<<10)

		_, err := conn.Read(buff)

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from the client: ", err)
			os.Exit(1)
		}

		conn.Write([]byte("+OK\r\n"))
	}

}
