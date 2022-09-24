package storage

type InMemoryStorage struct {
	data []byte
}

func (s *InMemoryStorage) Write(p []byte) (n int, err error) {
	s.data = append(s.data, p...)
	return len(p), nil
}

func (s *InMemoryStorage) GetData(ip string) []string {
	return []string{"Hello, World!"}
}