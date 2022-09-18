package server

import (
	"log"
	"net"
)

type Server interface {
	Run(addr string)
}

type tcpServer struct {
	C chan bool
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

		go func(c net.Conn) {
			log.Printf("Received connection from %s", c.RemoteAddr())
		}(conn)
	}
}
