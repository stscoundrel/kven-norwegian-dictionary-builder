package dictionary

import (
	"testing"

	"github.com/bradleyjkemp/cupaloy"
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

func TestKvenNorwegianDictionaryMatchesSnapshot(t *testing.T) {
	result := GetKvenNorwegianDictionary()

	cupaloy.SnapshotT(t, result)
}

func TestNorwegianKvenDictionaryHasExpectedAmountOfEntries(t *testing.T) {
	result := GetNorwegianKvenDictionary()

	expected := 7715

	if len(result) != expected {
		t.Error("Did not return expected content. Received", len(result), "expected ", expected)
	}
}

func TestNorwegianKvenDictionaryHasExpectedContent(t *testing.T) {
	result := GetNorwegianKvenDictionary()

	expected := DictionaryEntry{
		Headword:    "flittig",
		Stems:       []string{},
		Definitions: []string{"ahkera", "uuttera", "virree", "ahkerasti"},
		Synonymes:   []string{},
	}

	if result[2].Headword != expected.Headword {
		t.Error("Did not return expected content. Received", result[2].Headword, "expected ", expected.Headword)
	}

	if len(result[2].Stems) != len(expected.Stems) {
		t.Error("Did not return expected content. Received", len(result[2].Stems), "expected ", len(expected.Stems))
	}

	for index, definition := range result[2].Definitions {
		if definition != expected.Definitions[index] {
			t.Error("Did not return expected content. Received", definition, "expected ", expected.Definitions[index])
		}
	}

	if len(result[2].Synonymes) != len(expected.Synonymes) {
		t.Error("Did not return expected content. Received", len(result[2].Synonymes), "expected ", len(expected.Synonymes))
	}
}

func TestNorwegianKvenDictionaryMatchesSnapshot(t *testing.T) {
	result := GetNorwegianKvenDictionary()

	cupaloy.SnapshotT(t, result)
}
