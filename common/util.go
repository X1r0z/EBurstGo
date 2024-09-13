package common

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func LoadDict(filePath string) []string {
	var dict []string

	fp, _ := os.Open(filePath)
	r := bufio.NewReader(fp)
	defer fp.Close()

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if string(line) != "" {
			dict = append(dict, string(line))
		}
	}

	return dict
}

func LoadAllDict(filePath string) map[string]string {
	var dict map[string]string

	fp, _ := os.Open(filePath)
	r := bufio.NewReader(fp)
	defer fp.Close()

	for {
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}

		if string(line) != "" {
			username, password, _ := strings.Cut(string(line), ":")
			dict[username] = password
		}
	}

	return dict
}
