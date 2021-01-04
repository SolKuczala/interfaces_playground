package cli

import (
	"fmt"
)

type NoOPtool struct {
	Core *CoreTool
}

func (noOPtool *NoOPtool) List() ([]byte, error) {
	return noOPtool.Core.List()
}

func (noOPtool *NoOPtool) Save(value string) error {
	fmt.Printf("[NOOP] Saving %s\n", value)
	return nil
}

func (noOPtool *NoOPtool) Delete(value string) error {
	fmt.Printf("[NOOP] Deleting %s\n", value)
	return nil
}
