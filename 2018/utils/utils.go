package utils

import (
	"io/ioutil"
	"strings"
)

// ReadFileLines got path of file and read it into string slice
func ReadFileLines(path string) ([]string, error) {
	c, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	data := strings.Split(string(c), "\n")

	return data, nil
}
