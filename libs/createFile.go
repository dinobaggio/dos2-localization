package libs

import (
	"fmt"
	"log"
	"os"
)

func CreateFileWithContent(fileName string, content []byte) error {
	err := os.WriteFile(fileName, content, 0644)
	if err != nil {
		log.Fatal("Tidak dapat membuat file:", err)
		return err
	}

	fmt.Println("File berhasil dibuat:", fileName)
	return nil
}
