package main

import (
	"alanvivona/interchotas/pkg/cli"
	"alanvivona/interchotas/pkg/storage"
	"errors"
	"fmt"
)

var Tool cli.Tool

func main() {
	action, value, isNoOP, network, err := getInputFlags()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Got input: %+v %+v --no-op=%+v network=%+v\n", action, value, isNoOP, network)

	var Storage cli.Storage
	if network {
		Storage = &storage.NetStorage{}
	} else {
		Storage = &storage.FileStorage{}
	}

	coreFun := &cli.CoreTool{Storage: Storage}

	if isNoOP {
		Tool = &cli.NoOPtool{Core: coreFun}
	} else {
		Tool = coreFun
	}

	err = nil
	switch action {
	case "save":
		err = Tool.Save(value)
	case "delete":
		err = Tool.Delete(value)
	default:
		err = errors.New(fmt.Sprintf("Action %s not implemented\n", value))
	}

	if err != nil {
		panic(err)
	}

	fmt.Println("==============================")
	data, err := Tool.List()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}

func getInputFlags() (string, string, bool, bool, error) {
	noop := false
	network := false
	return "save", "pepe", noop, network, nil
}
