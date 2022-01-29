package parser

import (
	"encoding/xml"

	"github.com/stscoundrel/kven-norwegian-dictionary-builder/entries"
)

func ParseDictionary(rawDictionary []byte) entries.DictionaryEntries {
	var entries entries.DictionaryEntries
	xml.Unmarshal(rawDictionary, &entries)

	return entries
}
