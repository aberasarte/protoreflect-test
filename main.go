package main

import (
	"fmt"

	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/desc/protoprint"
)

func main() {
	fds, _ := protoparse.Parser{}.ParseFiles("test.proto")

	printer := &protoprint.Printer{
		SortElements: true,
	}

	origstring, _ := printer.PrintProtoToString(fds[0])

	fmt.Println(origstring)
}
