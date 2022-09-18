package server

import (
	"net"
	"testing"
)

func TestServer(t *testing.T) {
	server := tcpServer{C: make(chan bool)}

	go server.Run(":8080")

	if <-server.C {
		t.Log("Server started")
	}
	_, err := net.Dial("tcp", ":8080")
	if err != nil {
		t.Fatalf("got %s, expecte err to be nil", err.Error())
	}
}
