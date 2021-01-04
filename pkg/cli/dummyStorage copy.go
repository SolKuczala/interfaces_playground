package cli

type DummyStorage struct {
	WriteFN func(path string, data []byte) error
	ReadFN  func(path string) ([]byte, error)
}

func (s *DummyStorage) Write(path string, data []byte) error {
	return s.WriteFN(path, data)
}

func (s *DummyStorage) Read(path string) ([]byte, error) {
	return s.ReadFN(path)
}
