package storage

import (
	"bytes"
)

// InMemoryStorage is a storage that stores data in memory
type InMemoryStorage struct {
	data []byte
	length []int
}

func (s *InMemoryStorage) Write(p []byte) (n int, err error) {
	n = len(p)
	s.data = append(s.data, p...)
	s.length = append(s.length, n)
	return
}

func (s *InMemoryStorage) GetData(ip string) (data []string) {
	pre := 0
	for _, length := range s.length {
		bs := bytes.Split(s.data[pre:pre+length], []byte(":-"))	
		if string(bs[1]) == ip {
			data = append(data, string(s.data)[pre:pre+length])
		}
		pre = pre+length
	}
	return
}
