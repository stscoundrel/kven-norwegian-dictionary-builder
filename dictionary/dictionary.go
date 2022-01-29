package dictionary

import (
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/entries"
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/parser"
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/reader"
)

var KVEN_NORWEGIAN = "kven-norwegian"
var NORWEGIAN_KVEN = "norwegian-kven"

// Expose local versions of entries.
type DictionaryEntry = entries.DictionaryEntry

func GetKvenNorwegianDictionary() []DictionaryEntry {
	bytes := reader.ReadXmlDictionary(KVEN_NORWEGIAN)
	entries := parser.ParseDictionary(bytes)

	return entries.Entries
}

func GetNorwegianKvenDictionary() []DictionaryEntry {
	bytes := reader.ReadXmlDictionary(NORWEGIAN_KVEN)
	entries := parser.ParseDictionary(bytes)

	return entries.Entries
}
