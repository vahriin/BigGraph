package xmlparse

import (
	"encoding/xml"
	"io/ioutil"

	"github.com/vahriin/BigGraph/types"
)

func XMLRead(filename string) *types.Meta {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	document := new(types.Meta)

	err = xml.Unmarshal(file, &document)
	if err != nil {
		panic(err)
	}

	return document
}
