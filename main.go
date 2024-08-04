package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	var wordSettingXml *zip.File
	fuminaDocx, err := zip.OpenReader("test.docx")
	if err != nil {
		panic(err)
	}
	defer fuminaDocx.Close()

	wordDocx, err := zip.OpenReader("test2.docx")
	if err != nil {
		panic(err)
	}	
	defer wordDocx.Close()
	

	fmt.Println("Word docx")
	for _, el := range wordDocx.File {
		// fmt.Println(el.Name)
		if el.Name == "word/settings.xml" {
			wordSettingXml = el
		}
	}

	fuminaDocx.File = append(fuminaDocx.File, wordSettingXml)

	fmt.Println("Fumina docx")
	for _, el := range fuminaDocx.File {
		fmt.Printf("%#v\n", el.Name)
	}
}

func printBodyOfXml(zipFile *zip.File) {
	file, err := zipFile.Open()
	if err != nil {
		panic(err)
	}
	decoder := xml.NewDecoder(file)
	var inElement string
	for {
		t, err := decoder.Token()
		if err != nil {
				if err == io.EOF {
						break
				}
				fmt.Println("Error:", err)
				return
		}

		switch se := t.(type) {
		case xml.StartElement:
				inElement = se.Name.Local
				fmt.Printf("Start element: %s\n", inElement)
		case xml.CharData:
				fmt.Printf("Char data: %s\n", se)
		case xml.EndElement:
				fmt.Printf("End element: %s\n", se.Name.Local)
		}
	}
}

func saveModifiedZip(zip *zip.ReadCloser) {
	tempDir, err := os.MkdirTemp("", "zip-temp")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(tempDir)

	for _, zipElement := range zip.File {
		filePath := filepath.Join(tempDir, zipElement.Name)

		if !zipElement.FileInfo().IsDir() {
			rc, err := zipElement.Open()
			if err != nil {
				panic(err)
			}

			defer rc.Close()

			w, err := os.Create(filePath)
			
		}

	}
}