package utils

import (
	"io/ioutil"
	"os"
	"strings"
)

const (
	IGNOREFILE = ".docfmtignore"
)

func Loadignore() (lines []string) {
	if home := os.Getenv("HOME"); home != "" {
		if data, err := ioutil.ReadFile((home + string(os.PathSeparator) + IGNOREFILE)); err == nil {
			lines = append(lines, strings.Split(string(data), "\n")...)
		}
	}
	if data, err := ioutil.ReadFile(IGNOREFILE); err == nil {
		lines = append(lines, strings.Split(string(data), "\n")...)
	}
	return
}
