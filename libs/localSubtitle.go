package libs

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"

	"golang.org/x/text/language"
)

type Header struct {
	Version int `xml:"version,attr"`
	Time    int `xml:"time,attr"`
}

type Version struct {
	Major    int `xml:"major,attr"`
	Minor    int `xml:"minor,attr"`
	Revision int `xml:"revision,attr"`
	Build    int `xml:"build,attr"`
}

type Attribute struct {
	ID    string `xml:"id,attr"`
	Value string `xml:"value,attr"`
	Type  string `xml:"type,attr"`
}

type TextNode struct {
	NodeName   string      `xml:"node,attr"`
	Attributes []Attribute `xml:"attribute"`
	Children   []TextNode  `xml:"children>node"`
}

type Region struct {
	ID        string     `xml:"id,attr"`
	Subtitles []TextNode `xml:"node>children>node"`
}

type Save struct {
	XMLName xml.Name `xml:"save"`
	Header  Header   `xml:"header"`
	Version Version  `xml:"version"`
	Region  Region   `xml:"region"`
}

func LocalSubtitle(xmlData []byte, fileName string) error {
	pattern := `<attribute id="Text" value="([^"]+)" handle="([^"]+)" type="28" />`
	re := regexp.MustCompile(pattern)
	// Find matches
	input := string(xmlData)
	matches := re.FindAllStringSubmatch(input, -1)
	// fmt.Println(matches)
	for _, match := range matches {
		oldValue := match[1]
		fmt.Println(oldValue)
		newValue, err := TranslateText(language.Indonesian.String(), oldValue)
		if err != nil {
			return err
		}
		input = strings.ReplaceAll(input, oldValue, newValue)
	}

	err := CreateFileWithContent(fmt.Sprint("output/English/Subtitles/", fileName), []byte(input))
	if err != nil {
		return err
	}

	// fmt.Println(string(jsonData))
	return nil
}
