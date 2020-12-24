package main

import (
	"fmt"
	"github.com/im-a-giraffe/hello-go/gitparser"
	"github.com/im-a-giraffe/hello-go/slicer"
)

func main() {
	//server.Run()
	fmt.Println(gitparser.GetVersion())

	slicer.SliceTest()
	//httpserver.RunHttpServer()
}
