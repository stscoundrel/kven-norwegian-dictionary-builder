package dictionary

import (
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/parser"
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/reader"
)

var KVEN_NORWEGIAN = "kven-norwegian"
var NORWEGIAN_KVEN = "norwegian-kven"

func GetKvenNorwegianDictionary() DictionaryEntries {
	bytes := reader.ReadXmlDictionary(KVEN_NORWEGIAN)
	entries := parser.ParseDictionary(bytes)

	return entries.Entries
}

func GetNorwegianKvenDictionary() DictionaryEntries {
	bytes := reader.ReadXmlDictionary(NORWEGIAN_KVEN)
	entries := parser.ParseDictionary(bytes)

	return entries.Entries
}
