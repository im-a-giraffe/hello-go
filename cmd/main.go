package main

import (
	"fmt"
	"github.com/im-a-giraffe/hello-go/gitparser"
)

func main() {
	//server.Run()
	fmt.Println(gitparser.GetVersion())
	gitparser.CloneSampleRepository()
}
