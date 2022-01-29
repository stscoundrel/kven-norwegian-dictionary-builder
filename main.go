package main

import (
	"fmt"

	"github.com/stscoundrel/kven-norwegian-dictionary-builder/dictionary"
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/entries"
)

func GetKvenNorwegianDictionary() []entries.DictionaryEntry {
	return dictionary.GetKvenNorwegianDictionary()
}

func GetNorwegianKvenDictionary() []entries.DictionaryEntry {
	return dictionary.GetNorwegianKvenDictionary()
}

func main() {
	fmt.Println(GetKvenNorwegianDictionary())
	fmt.Println(GetNorwegianKvenDictionary())
}
