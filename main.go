package main

import (
	"dos2/libs"
	"fmt"
	"os"

	"github.com/schollz/progressbar/v3"
)

func main() {
	defer fmt.Println("done")
	dirPath := "Localization/English/Subtitles"
	dirList, err := os.ReadDir(dirPath)
	if err != nil {
		panic(err)
	}

	bar := progressbar.Default(int64(len(dirList)))
	for _, item := range dirList {
		fileName := item.Name()
		byteContent, err := os.ReadFile(fmt.Sprint(dirPath, "/", fileName))
		if err != nil {
			panic(err)
		}
		err = libs.LocalSubtitle(byteContent, fileName)
		if err != nil {
			panic(err)
		}
		bar.Add(1)
	}

	fmt.Println("Done !")
}
