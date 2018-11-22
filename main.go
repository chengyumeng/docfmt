package main

import (
	"fmt"
	"os"

	"github.com/chengyumeng/docfmt/pkg/cmd"
)

func main() {
	err := cmd.FormatCmd.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
