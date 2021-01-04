package cli

const outFilePath = "./out.txt"

// const outFilePath = "http://0.0.0.0:8000/"

type Tool interface {
	Save(string) error
	Delete(string) error
	List() ([]byte, error)
}

type Storage interface {
	Write(string, []byte) error
	Read(string) ([]byte, error)
}
