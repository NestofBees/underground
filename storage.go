package underground

// Storage is a storage interface
type Storage interface {
	Write(p []byte) (n int, err error)
	GetData(ip string) []string
}
