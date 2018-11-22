package format

import (
	"fmt"
	"os"
	"io/ioutil"
	"regexp"
)

const (
	Latin string = "[A-Za-z0-9\u00C0-\u00FF\u0100-\u017F\u0180-\u024F\u1E00-\u1EFF]"
	//Asia string = "'[\u4E00-\u9FFF]', '[\u3400-\u4DB5\u9FA6-\u9FBB\uFA70-\uFAD9\u9FBC-\u9FC3\u3007\u3040-\u309E\u30A1-\u30FA\u30FD\u30FE\uFA0E-\uFA0F\uFA11\uFA13-\uFA14\uFA1F\uFA21\uFA23-\uFA24\uFA27-\uFA29]', '[\uD840-\uD868][\uDC00-\uDFFF]|\uD869[\uDC00-\uDEDF]', '\uD86D[\uDC00-\uDF3F]|[\uD86A-\uD86C][\uDC00-\uDFFF]|\uD869[\uDF00-\uDFFF]', '\uD86D[\uDF40-\uDFFF]|\uD86E[\uDC00-\uDC1F]', '[\u31C0-\u31E3]'"
	//Panc string = "'[@&=_\,\.\?\!\$\%\^\*\-\+\/]', '[\(\\[\'"<‘“]', '[\)\\]\'">”’]'"
)


type Document interface {
	Format() error
	ListFile() ([]string,error)
}

func NewBasicDocument(path string,match,ignore []string) Document {
	return &BasicDoc{
		path,match,ignore,[]string{}, string(os.PathSeparator),
	}
}

type BasicDoc struct {
	Path string
	Match []string
	Ignore []string
	files []string
	pathSeparator string
}

func (b *BasicDoc)Format() error {
	files,err := b.ListFile()
	if err != nil {
		return err
	}
	for _,f := range files {
		b.format(f)
	}
	return nil
}

func (b *BasicDoc) ListFile() ([]string,error) {
	f, err := os.Stat(b.Path)
	if err != nil {
		return nil,err
	}
	if !f.IsDir() {
		return []string{b.Path},nil
	}

	files,err := b.getAllFiles(b.Path)
	return files,err
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

func (b *BasicDoc)format(path string) error {
	fmt.Println(path)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	text := string(data)
	reg := regexp.MustCompile("([\u4E00-\u9FFF])([A-Za-z0-9\u00C0-\u00FF\u0100-\u017F\u0180-\u024F\u1E00-\u1EFF])")
	updateText := reg.ReplaceAllString(text, "$1 $2")
	if updateText != text {
		f, err := os.Create(path)
		if err != nil {
			return err
		} else {
			defer f.Close()
			f.Write([]byte(updateText))
		}
	}
	return nil
}

func (b *BasicDoc) edit(f os.File) error {
	return nil
}

func (b *BasicDoc) isAvaliable(file os.FileInfo) bool {
	right := true
	if len(b.Match) > 0 {
		for _,s := range b.Match {
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
		for _,s := range b.Ignore {
			i, err = regexp.MatchString(s, file.Name())
			if err != nil {
				return false
			}

		}
		right = !i  && right
	}
	return right
}


