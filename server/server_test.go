package server

import (
	"errors"
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/NestofBees/underground/storage"
)

func TestServer(t *testing.T) {
	storage := &storage.InMemoryStorage{}
	server := tcpServer{storage: storage, C: make(chan bool)}

	go server.Run(":8080")

	if <-server.C {
		t.Log("Server started")
	}

	t.Run("Test connection", func(t *testing.T) {
		conn, err := net.Dial("tcp6", ":8080")
		assertErrorEqual(t, nil, err)
		defer conn.Close()
	})

	t.Run("Test send message", func(t *testing.T) {
		conn, err := net.Dial("tcp6", ":8080")
		assertErrorEqual(t, nil, err)
		defer conn.Close()

		data := []byte(fmt.Sprintf("%d:-%s:-%s", time.Now().Unix(), conn.LocalAddr().String(), "Hello world"))
		n, err := conn.Write(data)
		assertErrorEqual(t, nil, err)
		assertIntEqual(t, len(data), n)
	})

	t.Run("Test save message", func(t *testing.T) {
		conn, err := net.Dial("tcp6", ":8080")
		if err != nil {
			t.Fatalf("got %s, expecte err to be nil", err.Error())
		}
		defer conn.Close()

		data := []byte(fmt.Sprintf("%d:-%s:-%s", time.Now().Unix(), conn.LocalAddr().String(), "Hello world"))
		n, err := conn.Write(data)
		assertErrorEqual(t, nil, err)
		assertIntEqual(t, len(data), n)
		// wait data save to storage
		time.Sleep(10 * time.Millisecond)
		asserStringSliceEqual(t, []string{string(data)}, storage.GetData(conn.LocalAddr().String()))
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
		t.Fatalf("got %v, expected %v, length not same", got, want)
	}

	for i := 0; i < len(want); i++ {
		if want[i] != got[i] {
			for k := 0; k < len(want[i]); k++ {
				t.Logf("equal: %v", want[i][k] == got[i][k])
			}
		}
	}
}
