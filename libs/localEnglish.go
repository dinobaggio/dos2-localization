package libs

import (
	"encoding/xml"
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/text/language"
)

type Content struct {
	ContentUID string `xml:"contentuid,attr"`
	Value      string `xml:",chardata"`
}

type ContentList struct {
	XMLName xml.Name  `xml:"contentList"`
	Items   []Content `xml:"content"`
}

func LocalEnglish() {
	args := os.Args[1:]
	byteContent, err := os.ReadFile(args[0])
	if err != nil {
		panic(err)
	}
	var contentList ContentList
	err = xml.Unmarshal(byteContent, &contentList)
	if err != nil {
		panic(err)
	}
	var translateContent []Content
	startNumb := 62454
	contents := contentList.Items[startNumb:]
	bar := progressbar.Default(int64(len(contents)))
	for _, item := range contents {
		text, err := TranslateText(language.Indonesian.String(), item.Value)
		if err != nil {
			fmt.Println(err)
			break
		}
		translateContent = append(translateContent, Content{
			item.ContentUID,
			text,
		})
		bar.Add(1)
	}

	transList, err := xml.MarshalIndent(ContentList{
		XMLName: contentList.XMLName,
		Items:   translateContent,
	}, "", "  ")

	if err != nil {
		panic(err)
	}

	err = CreateFileWithContent("output/English/english.xml", transList)
	if err != nil {
		panic(err)
	}

	fmt.Println("done")
}
