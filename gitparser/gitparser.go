package gitparser

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"os"
)

func GetVersion() string {
	return "1.0.0"
}

func CloneSampleRepository() {
	fmt.Println("git clone https://github.com/go-git/go-git")

	_, err := git.PlainClone("/tmp/foo", false, &git.CloneOptions{
		URL:      "https://github.com/go-git/go-git",
		Progress: os.Stdout,
	})

	fmt.Println(err)
}
