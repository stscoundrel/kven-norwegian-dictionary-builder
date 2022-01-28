package main

import (
	"fmt"
	"testing"
)

func TestGetsKvenNorwegianDictionary(t *testing.T) {
	const expected = "kansalinen"
	result := GetKvenNorwegianDictionary()

	// for _, entry := range result.Entries {
	// 	fmt.Println("Headword: " + entry.Headword)
	// 	fmt.Println("Stems: ")
	// 	for _, stem := range entry.Stems {
	// 		fmt.Println(stem)
	// 	}
	// 	fmt.Println("Definitions: ")
	// 	for _, definition := range entry.Definitions {
	// 		fmt.Println(definition)
	// 	}
	// 	fmt.Println("Synonymes: ")
	// 	for _, synonyme := range entry.Synonymes {
	// 		fmt.Println(synonyme)
	// 	}
	// 	fmt.Println("---------------")
	// }

	fmt.Println(result.Entries[0].Headword)

	if result.Entries[0].Headword != "kansalinen" {
		t.Error("Did not return expected content. Received", result, "expected kansalinen")
	}
}
