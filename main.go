package main

import (
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/dictionary"
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/entries"
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/writer"
)

func GetKvenNorwegianDictionary() []entries.DictionaryEntry {
	return dictionary.GetKvenNorwegianDictionary()
}

func GetNorwegianKvenDictionary() []entries.DictionaryEntry {
	return dictionary.GetNorwegianKvenDictionary()
}

func ToJson() {
	kvenToNorwegian := GetKvenNorwegianDictionary()
	norwegianToKven := GetNorwegianKvenDictionary()

	writer.WriteJson("build/kven-norwegian.json", kvenToNorwegian)
	writer.WriteJson("build/norwegian-kven.json", norwegianToKven)
}

func main() {
	ToJson()
}
