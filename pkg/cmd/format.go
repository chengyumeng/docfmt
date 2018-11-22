package cmd

import (
	"github.com/chengyumeng/docfmt/pkg/format"
	"github.com/spf13/cobra"
)

var FormatCmd = &cobra.Command{
	Use:   "docfmt",
	Short: "Add one space between Chinese and English.",
	Long: `A tool for automatically adding one space between Chinese and English.
Power By https://github.com/chengyumeng
Wechat Public Accounta: 程天写代码
`,
	Example: `docfmt $(pwd) --debug`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: doFormate,
}

var option = format.Option{
	Ignore: []string{"^\\..*$"},
}

func init() {
	FormatCmd.Flags().StringVarP(&option.Path, "path", "p", ".", "")
	FormatCmd.Flags().BoolVarP(&option.Debug, "debug", "d", false, "")
}

func doFormate(cmd *cobra.Command, args []string) error {
	doc := format.NewBasicDocument(option)
	return doc.Format()
}
