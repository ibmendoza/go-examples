package main

import (
	//"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, _ := ioutil.ReadFile("README.md")
	HTMLRender := blackfriday.MarkdownCommon(file)
	newFile, err := os.Create("README.html")
	if err != nil {
		log.Fatal(err)
	}
	_, err = newFile.Write(HTMLRender)
	if err != nil {
		log.Fatal(err)
	}
}
