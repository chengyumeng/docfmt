package format

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

const (
	AsiaLatin  = "([\u4E00-\u9FFF])([A-Za-z0-9\u00C0-\u00FF\u0100-\u017F\u0180-\u024F\u1E00-\u1EFF])"
	LatingAsia = "([A-Za-z0-9\u00C0-\u00FF\u0100-\u017F\u0180-\u024F\u1E00-\u1EFF])([\u4E00-\u9FFF])"
)

type Document interface {
	Format() error
	ListFile() ([]string, error)
}

func NewBasicDocument(opt Option) Document {
	return &BasicDoc{
		opt.Path, opt.Match, opt.Ignore, opt.Debug, []string{}, string(os.PathSeparator),
	}
}

type BasicDoc struct {
	Path          string
	Match         []string
	Ignore        []string
	Debug         bool
	files         []string
	pathSeparator string
}

func (b *BasicDoc) Format() error {
	files, err := b.ListFile()
	if err != nil {
		return err
	}
	if b.Debug == true {
		for _, f := range files {
			if err := b.preFormat(f); err != nil {
				return err
			}
		}
	} else {
		for _, f := range files {
			if err := b.format(f); err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *BasicDoc) ListFile() ([]string, error) {
	f, err := os.Stat(b.Path)
	if err != nil {
		return nil, err
	}
	if !f.IsDir() {
		if b.isAvaliable(f) {
			return []string{b.Path}, nil
		} else {
			return []string{}, nil
		}
	}

	files, err := b.getAllFiles(b.Path)
	return files, err
}

func (b *BasicDoc) getAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if !b.isAvaliable(fi) {
			continue
		}
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+b.pathSeparator+fi.Name())
		} else {
			// 过滤指定格式
			files = append(files, dirPth+b.pathSeparator+fi.Name())
		}
	}

	for _, table := range dirs {
		temp, _ := b.getAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}
	return files, nil
}

func (b *BasicDoc) format(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	text := string(data)
	reg1 := regexp.MustCompile(AsiaLatin)
	updateText := reg1.ReplaceAllString(text, "$1 $2")
	reg2 := regexp.MustCompile(LatingAsia)
	updateText = reg2.ReplaceAllString(updateText, "$1 $2")
	if updateText != text {
		f, err := os.Create(path)
		if err != nil {
			return err
		} else {
			defer f.Close()
			_, err := f.Write([]byte(updateText))
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (b *BasicDoc) preFormat(path string) error {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	txts := strings.Split(string(data), "\n")
	for line, txt := range txts {
		updateTxt := txt
		reg1 := regexp.MustCompile(AsiaLatin)
		arr := reg1.FindAllString(txt, -1)
		for _, a := range arr {
			updateTxt = strings.Replace(txt, a, color.RedString(a), -1)
		}
		reg2 := regexp.MustCompile(LatingAsia)
		arr = reg2.FindAllString(txt, -1)
		for _, a := range arr {
			updateTxt = strings.Replace(txt, a, color.RedString(a), -1)
		}
		if updateTxt != txt {
			fmt.Printf("%s %d %s\n", path, line, updateTxt)
		}
	}

	return nil
}

func (b *BasicDoc) isAvaliable(file os.FileInfo) bool {
	right := true
	if len(b.Match) > 0 {
		for _, s := range b.Match {
			m, err := regexp.MatchString(s, file.Name())
			if err != nil {
				return false
			}

			right = right && m
		}
	}
	if len(b.Ignore) > 0 {
		i := false
		var err error
		for _, s := range b.Ignore {
			i, err = regexp.MatchString(s, file.Name())
			if err != nil {
				return false
			}

		}
		right = !i && right
	}
	return right
}
