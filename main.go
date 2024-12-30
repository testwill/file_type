package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/h2non/filetype"
	"flag"
)

var (
	fooType = filetype.NewType("clamav", "clamav/clamav")
	fileMagic  = "ClamAV-VDB:"
	magicLen = len(fileMagic)
)

func fooMatcher(buf []byte) bool {
	if  len(buf) < magicLen {
		return false
	}

	return strings.HasPrefix(string(buf[:magicLen]), fileMagic)
}

func main() {
	var (
		fName string
	)
	flag.StringVar(&fName, "file", "", "file name")
	flag.Parse()
	fmt.Println(fName)
	// Register the new matcher and its type
	filetype.AddMatcher(fooType, fooMatcher)

	// Check if the new type is supported by extension
	if filetype.IsSupported("clamav") {
		fmt.Println("New supported type: clamav")
	}

	// Check if the new type is supported by MIME
	if filetype.IsMIMESupported("clamav/clamav") {
		fmt.Println("New supported MIME type: clamav/clamav")
	}

	file, _ := os.Open(fName)

	// We only have to pass the file header = first 261 bytes
	head := make([]byte, 261)
	file.Read(head)
	defer file.Close()
	kind, _ := filetype.Match(head)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
	} else {
		fmt.Printf("File type matched: %s\n", kind.Extension)
	}
}
