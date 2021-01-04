package storage

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type NetStorage struct {
}

func (s *NetStorage) Write(path string, data []byte) error {
	fullPath := path + "/" + string(data)
	fmt.Printf("[NET STORAGE] Write %s\n", fullPath)
	_, err := http.Get(fullPath)
	return err
}

func (s *NetStorage) Read(path string) ([]byte, error) {
	fmt.Printf("[NET STORAGE] Read %s\n", path)
	return ioutil.ReadFile(path)
}
