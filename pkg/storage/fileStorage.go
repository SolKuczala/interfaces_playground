package storage

import (
	"fmt"
	"io/ioutil"
	"os"
)

type FileStorage struct {
}

func (s *FileStorage) Write(path string, data []byte) error {
	fmt.Printf("[FILE STORAGE] Write to %s\n", path)
	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}
	f.Sync()

	return nil
}

func (s *FileStorage) Read(path string) ([]byte, error) {
	fmt.Printf("[FILE STORAGE] Read %s\n", path)
	return ioutil.ReadFile(path)
}
