package main

import (
	"fmt"
	"io/ioutil"

	"github.com/h2non/filetype"
)

func main() {
	file, _ := ioutil.ReadFile("./testload.png")

	kind, err := filetype.Match(file)
	if err != nil {
		fmt.Println(err)
		panic("500::Read file type is failed")
	}

	if kind == filetype.Unknown {
		panic("500::Unknown file type")
	}

	fmt.Println(kind.MIME.Value, kind.MIME.Subtype)
}
