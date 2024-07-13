package main

import (
	"fmt"
	"io"
	"net"
	"os"

	"main.go/resp"
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

		resp := resp.NewResp(conn)

		value, err := resp.Read()

		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Println("Error reading from the client: ", err)
			os.Exit(1)
		}

		fmt.Println(value)

		conn.Write([]byte("+OK\r\n"))
	}

}
