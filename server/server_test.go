package server

import (
	"bytes"
	"net"
	"reflect"
	"testing"
)

func TestServer(t *testing.T) {
	buff := make([]byte, 512)
	buffer := bytes.NewBuffer(buff)
	server := tcpServer{writer: buffer, C: make(chan bool)}

	go server.Run(":8080")

	if <-server.C {
		t.Log("Server started")
	}

	t.Run("Test connection", func(t *testing.T) {
		conn, err := net.Dial("tcp", ":8080")
		if err != nil {
			t.Fatalf("got %s, expecte err to be nil", err.Error())
		}
		defer conn.Close()
	})

	t.Run("Test send message", func(t *testing.T) {
		conn, err := net.Dial("tcp", ":8080")
		if err != nil {
			t.Fatalf("got %s, expecte err to be nil", err.Error())
		}
		defer conn.Close()

		_, err = conn.Write([]byte("Hello, World!"))
		if err != nil {
			t.Fatalf("got %s, expecte err to be nil", err.Error())
		}
		asserStringEqual(t, "Hello, World!", buffer.String())
	})
}

func asserStringEqual(t *testing.T, want, got string) {
	if reflect.DeepEqual(want, got) {
		t.Fatalf("got %s, expected %s", got, want)
	}
}
	
	
	
	
	
	
	
	