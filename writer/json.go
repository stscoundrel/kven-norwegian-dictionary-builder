package writer

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/stscoundrel/kven-norwegian-dictionary-builder/entries"
)

func WriteJson(filePath string, entries []entries.DictionaryEntry) {
	file, _ := json.MarshalIndent(entries, "", "  ")

	err := ioutil.WriteFile(filePath, file, 0644)

	if err != nil {
		log.Fatalf("Error creating JSON file: %s", err)
	}
}
