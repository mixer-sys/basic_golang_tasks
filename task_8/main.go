package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

func ReadFile(path string) (string, error) {

	fileInfo, err := os.Stat(path)
	if os.IsNotExist(err) {
		return "", errors.New("file does not exist")
	}

	if err != nil {
		return "", err
	}

	if fileInfo.Size() > 1*1024*1024 {
		return "", errors.New("file is too large")
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), err
}

func main() {

	content, err := ReadFile("test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(content)
	}
}
