package server

import (
	"fmt"
	"log"
	"net"

	"github.com/NestofBees/underground/storage"
)

type Server interface {
	Run(addr string)
}

type tcpServer struct {
	storage storage.Storage
	C      chan bool
}

func (s *tcpServer) Run(addr string) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	s.C <- true

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 512)
	_, err := conn.Read(buff)
	if err != nil {
		fmt.Printf("read error %s\n", err.Error())
		return
	}
	fmt.Fprintf(conn, "%s", string(buff))
}
