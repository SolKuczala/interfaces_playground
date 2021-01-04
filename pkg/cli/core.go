package cli

import (
	"fmt"
)

type CoreTool struct {
	Storage Storage
}

func (coreTool *CoreTool) List() ([]byte, error) {
	fmt.Printf("[CORE] Listing from : %s\n", outFilePath)
	data, err := coreTool.Storage.Read(outFilePath)
	fmt.Printf("[CORE] Bytes read: %d\n", len(data))
	return data, err
}

func (coreTool *CoreTool) Save(value string) error {
	fmt.Printf("[CORE] Saving %s\n", value)
	err := coreTool.Storage.Write(outFilePath, []byte(value))
	return err
}

func (coreTool *CoreTool) Delete(value string) error {
	fmt.Printf("[CORE] Deleting %s\n", value)
	return nil
}
