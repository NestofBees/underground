package storage

type Storage interface {
	Write(p []byte) (n int, err error)
	GetData(ip string) []string
}