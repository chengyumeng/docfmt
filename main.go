package main
import (
	"github.com/chengyumeng/docfmt/pkg/format"
)


func main() {
	d := format.NewBasicDocument(".",[]string{"READ"},[]string{`^\..*`,})
	d.Format()
}
