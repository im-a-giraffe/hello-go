package main

import (
	"fmt"
	"github.com/im-a-giraffe/hello-go/database"
	"github.com/im-a-giraffe/hello-go/gitparser"
)

func main() {
	//server.Run()
	fmt.Println(gitparser.GetVersion())

	//slicer.SliceTest()
	//httpserver.RunHttpServer()

	//slicer.HelloInterfaces()

	//concurrency.Run()
	database.RunSQLite("simple.sqlite")
}
