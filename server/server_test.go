package server

import (
	"errors"
	"net"
	"testing"

	"github.com/NestofBees/underground/storage"
)

func TestServer(t *testing.T) {
	storage := &storage.InMemoryStorage{}
	server := tcpServer{writer: storage, C: make(chan bool)}

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
		asserStringSliceEqual(t, []string{"Hello, World!"}, storage.GetData(conn.LocalAddr().String()))
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

func asserStringSliceEqual(t *testing.T, want, got []string) {
	t.Helper()
	if len(want) != len(got) {
		t.Fatalf("got %v, expected %v", got, want)
	}

	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			t.Fatalf("got %v, expected %v", got, want)
		}
	}
}
	