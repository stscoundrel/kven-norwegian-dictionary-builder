package dictionary

import (
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/parser"
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/reader"
)

var KVEN_NORWEGIAN = "kven-norwegian"

func GetKvenNorwegianDictionary() DictionaryEntries {
	bytes := reader.ReadXmlDictionary(KVEN_NORWEGIAN)
	entries := parser.ParseDictionary(bytes)

	return entries.Entries
}
