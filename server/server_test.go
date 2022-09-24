package server

import (
	"bytes"
	"errors"
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
		assertErrorEqual(t, nil, err)
		defer conn.Close()
	})

	t.Run("Test send message", func(t *testing.T) {
		conn, err := net.Dial("tcp", ":8080")
		assertErrorEqual(t, nil, err)
		defer conn.Close()

		data := []byte("Hello world")
		n, err := conn.Write(data)
		assertErrorEqual(t, nil, err)
		assertIntEqual(t, len(data), n)
	})

	t.Run("Test save message", func(t *testing.T) {
		conn, err := net.Dial("tcp", ":8080")
		if err != nil {
			t.Fatalf("got %s, expecte err to be nil", err.Error())
		}
		defer conn.Close()

		data := []byte("Hello world")
		n, err := conn.Write(data)
		assertErrorEqual(t, nil, err)	
		assertIntEqual(t, len(data), n)
		asserStringEqual(t, "Hello, World!", buffer.String())
	})
}

func assertErrorEqual(t *testing.T, want, got error) {
	t.Helper()
	if !errors.Is(want, got) {
		t.Fatalf("got %s, expected %s", got, want)
	}
}

func assertIntEqual(t testing.TB, want, got int) {
	t.Helper()
	if want != got {
		t.Fatalf("got %d, expected %d", got, want)
	}
}

func asserStringEqual(t *testing.T, want, got string) {
	t.Helper()
	if reflect.DeepEqual(want, got) {
		t.Fatalf("got %s, expected %s", got, want)
	}
}
	
	
	
	
	
	
	
	