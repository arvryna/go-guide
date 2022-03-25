package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gomarkdown/markdown"
)

func main() {
	fileName := ""

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	} else {
		panic("No file passed, please pass a filename")
	}

	markDownFileName := fileName + ".md"
	htmlOutFileName := fileName + ".html"

	fmt.Println("Converting file.. ", markDownFileName)
	data, err := ioutil.ReadFile(markDownFileName)
	if err != nil {
		panic(err)
	}

	// create a output file directory to store the generated HTML files
	// if err = os.Mkdir("out", os.ModePerm); err != nil {
	// 	panic(err)
	// }

	html := markdown.ToHTML(data, nil, nil)

	if err = ioutil.WriteFile(htmlOutFileName, html, 0644); err != nil {
		panic(err)
	}

	fmt.Println("HTML file generated: " + htmlOutFileName)

}
