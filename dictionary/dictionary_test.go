package dictionary

import (
	"testing"
)

func TestKvenNorwegianDictionaryHasExpectedAmountOfEntries(t *testing.T) {
	result := GetKvenNorwegianDictionary()

	expected := 8468

	if len(result) != expected {
		t.Error("Did not return expected content. Received", len(result), "expected ", expected)
	}
}

func TestKvenNorwegianDictionaryHasExpectedContent(t *testing.T) {
	result := GetKvenNorwegianDictionary()

	expected := DictionaryEntry{
		Headword:    "historialinen",
		Stems:       []string{},
		Definitions: []string{"historisk"},
		Synonymes:   []string{"histoorialinen", "histoorillinen"},
	}

	if result[1].Headword != expected.Headword {
		t.Error("Did not return expected content. Received", result[1].Headword, "expected ", expected.Headword)
	}

	if len(result[1].Stems) != len(expected.Stems) {
		t.Error("Did not return expected content. Received", len(result[1].Stems), "expected ", len(expected.Stems))
	}

	for index, definition := range result[1].Definitions {
		if definition != expected.Definitions[index] {
			t.Error("Did not return expected content. Received", definition, "expected ", expected.Definitions[index])
		}
	}

	for index, synonyme := range result[1].Synonymes {
		if synonyme != expected.Synonymes[index] {
			t.Error("Did not return expected content. Received", synonyme, "expected ", expected.Synonymes[index])
		}
	}
}
