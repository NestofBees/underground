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
	l, err := net.Listen("tcp6", addr)
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

		go s.receiveMessage(conn)
	}
}

func (s *tcpServer) receiveMessage(conn net.Conn) {
	defer conn.Close()
	buff := make([]byte, 512)
	n, err := conn.Read(buff)
	if err != nil {
		fmt.Printf("read error %s\n", err.Error())
		return
	}
	_, _ = fmt.Fprintf(s.storage, "%s", string(buff[:n]))
}
