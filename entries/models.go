package entries

import "encoding/xml"

type DictionaryEntry struct {
	Headword    string   `xml:"lg>l" json:"headword"`
	Stems       []string `xml:"lg>stem" json:"stems"`
	Definitions []string `xml:"mg>tg>t" json:"definitions"`
	Synonyms    []string `xml:"mg>tg>syng>syn" json:"synonyms"`
}

type DictionaryEntries struct {
	XMLName xml.Name          `xml:"r"`
	Entries []DictionaryEntry `xml:"e"`
}
