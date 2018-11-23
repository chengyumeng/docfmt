package cmd

import (
	"github.com/chengyumeng/docfmt/pkg/format"
	"github.com/chengyumeng/docfmt/pkg/utils"
	"github.com/spf13/cobra"
)

var FormatCmd = &cobra.Command{
	Use:   "docfmt",
	Short: "中英文之间自动增加空格的工具",
	Long: `一款性感的格式化中英文混合编写的纯文本文档的工具，自动在东亚语言和拉丁语言之间添加空格，使得文档更优雅。
开发者 https://github.com/chengyumeng
微信公众号: 程天写代码
`,
	Example: `docfmt $(pwd) --debug`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		option.Ignore = append(option.Ignore, utils.Loadignore()...)
		return nil
	},
	RunE: doFormate,
}

var option = format.Option{
	Ignore: []string{"^\\..*$"},
}

func init() {
	FormatCmd.Flags().StringVarP(&option.Path, "path", "p", ".", "需要格式化文档的路径/文件")
	FormatCmd.Flags().StringVarP(&option.Match, "match", "m", "md$", "按照给定规则过滤文件名（正则表达式，不符合规则的文件不格式化）")
	FormatCmd.Flags().BoolVarP(&option.Debug, "debug", "d", false, "是否使用 debug 模式，debug 模式下，只输出将要修改的文字（类似 git diff）")
}

func doFormate(cmd *cobra.Command, args []string) error {
	doc := format.NewBasicDocument(option)
	return doc.Format()
}
