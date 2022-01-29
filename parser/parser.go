package parser

import (
	"encoding/xml"

	"github.com/stscoundrel/kven-norwegian-dictionary-builder/entries"
)

func ParseDictionary(rawDictionary []byte) entries.DictionaryEntries {
	var entries entries.DictionaryEntries
	xml.Unmarshal(rawDictionary, &entries)

	// For writing purposes, convert nil slices to empty slices.
	for index, entry := range entries.Entries {
		if entry.Stems == nil {
			entries.Entries[index].Stems = []string{}
		}

		if entry.Definitions == nil {
			entries.Entries[index].Definitions = []string{}
		}

		if entry.Synonyms == nil {
			entries.Entries[index].Synonyms = []string{}
		}
	}

	return entries
}
