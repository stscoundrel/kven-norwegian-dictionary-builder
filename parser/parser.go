package parser

import (
	"encoding/xml"
)

func ParseDictionary(rawDictionary []byte) DictionaryEntries {
	var entries DictionaryEntries
	xml.Unmarshal(rawDictionary, &entries)

	return entries
}
