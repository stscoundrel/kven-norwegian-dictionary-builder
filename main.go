package main

import (
	"fmt"

	"github.com/stscoundrel/kven-norwegian-dictionary-builder/dictionary"
)

func GetKvenNorwegianDictionary() dictionary.DictionaryEntries {
	return dictionary.GetKvenNorwegianDictionary()
}

func main() {
	fmt.Println(GetKvenNorwegianDictionary())
}
