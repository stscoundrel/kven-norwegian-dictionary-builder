package parser

import "encoding/xml"

type DictionaryEntry struct {
	Headword    string   `xml:"lg>l"`
	Stems       []string `xml:"lg>stem"`
	Definitions []string `xml:"mg>tg>t"`
	Synonymes   []string `xml:"mg>tg>syng>syn"`
}

type DictionaryEntries struct {
	XMLName xml.Name          `xml:"r"`
	Entries []DictionaryEntry `xml:"e"`
}
