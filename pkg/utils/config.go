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
			for _,l := range strings.Split(string(data), "\n") {
				if len(l) > 0 && !strings.HasPrefix(l,"#") {
					lines = append(lines, l)
				}
			}
		}
	}
	if data, err := ioutil.ReadFile(IGNOREFILE); err == nil {
		for _,l := range strings.Split(string(data), "\n") {
			if len(l) > 0 && !strings.HasPrefix(l,"#") {
				lines = append(lines, l)
			}
		}
	}
	return
}
