package main

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestCreatesJsonBuild(t *testing.T) {
	ToJson()

	_, err1 := os.Stat("./build/kven-norwegian.json")
	_, err2 := os.Stat("./build/norwegian-kven.json")

	fmt.Println(err1)
	fmt.Println(err2)

	if errors.Is(err1, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for Kven - Norwegian dictionary")
	}

	if errors.Is(err2, os.ErrNotExist) {
		t.Error("JSON output file not found in build directory for Norwegian - Kven dictionary")
	}
}
