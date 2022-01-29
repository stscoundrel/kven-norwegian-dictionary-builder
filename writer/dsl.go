package writer

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/stscoundrel/kven-norwegian-dictionary-builder/dictionary"
	"github.com/stscoundrel/kven-norwegian-dictionary-builder/entries"
)

func formatEntry(entry entries.DictionaryEntry) []string {
	entryLines := []string{}

	entryLines = append(entryLines, entry.Headword)

	for _, definition := range entry.Definitions {
		entryLines = append(entryLines, "    "+definition)
	}

	if len(entry.Stems) > 0 {
		stemLine := "    Stem: " + strings.Join(entry.Stems, ", ")
		entryLines = append(entryLines, stemLine)
	}

	if len(entry.Synonyms) > 0 {
		synonymLine := "    Synonyms: " + strings.Join(entry.Synonyms, ", ")
		entryLines = append(entryLines, synonymLine)
	}

	return entryLines
}

func buildHeaders(variant string) []string {
	if variant == dictionary.KVEN_NORWEGIAN {
		return []string{
			"#NAME Kven-Norwegian Bokmål dictionary",
			"#INDEX_LANGUAGE \"Kven\"",
			"#CONTENTS_LANGUAGE \"Norwegian\"",
		}
	} else if variant == dictionary.NORWEGIAN_KVEN {
		return []string{
			"#NAME Norwegian Bokmål-Kven dictionary",
			"#INDEX_LANGUAGE \"Norwegian\"",
			"#CONTENTS_LANGUAGE \"Kven\"",
		}
	} else {
		log.Fatalf("Unkown dictionary variant: %s", variant)
	}

	return []string{}
}

func formatDSL(entries []entries.DictionaryEntry, variant string) []string {
	dsl := buildHeaders(variant)

	for _, entry := range entries {
		entryLines := formatEntry(entry)
		dsl = append(dsl, entryLines...)
	}

	return dsl
}

func WriteDsl(filePath string, entries []entries.DictionaryEntry, variant string) {
	dslLines := formatDSL(entries, variant)

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		log.Fatalf("Error creating DSL file: %s", err)
	}

	datawriter := bufio.NewWriter(file)

	for _, line := range dslLines {
		_, _ = datawriter.WriteString(line + "\n")
	}

	datawriter.Flush()
	file.Close()
}
