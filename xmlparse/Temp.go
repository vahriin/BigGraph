package xmlparse

import (
	"encoding/xml"
	"io/ioutil"
)

func XMLRead(filename string) Meta {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var document Meta

	err = xml.Unmarshal(file, &document)
	if err != nil {
		panic(err)
	}

	return document
}


