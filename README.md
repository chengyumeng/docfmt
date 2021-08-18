# docfmt

[![Build Statue](https://travis-ci.com/chengyumeng/docfmt.svg?branch=master)](https://travis-ci.org/chengyumeng/docfmt)
[![Build Tag](https://img.shields.io/github/tag/chengyumeng/docfmt.svg)](https://github.com/chengyumeng/docfmt/releases)

有研究显示，打字的时候不喜欢在中文和英文之间加空格的人，感情路都走得很辛苦，有七成的比例会在 34 岁的时候跟自己不爱的人结婚，而其余三成的人最后只能把遗产留给自己的猫。

毕竟爱情跟书写都需要适时地留白。与大家共勉之。

其实这个问题只存在于工程师的视野里，人工智能的天才们是不需要考虑的，毕竟他们的世界里只需要写 Word 文档和做 PPT，并不需要写代码、文档和注释。

对东亚国家的工程师而言，编写文档和注释经常面对中英文混合编写的情形。正如我们总能感觉到那些前端页面少了两个像素的对齐会有点不舒服一样，有些洋洋洒洒的文档经常会被感觉别扭。

在经过很长时间的探索，我们总会发现他们在东亚文字和拉丁文字之间没有使用空格，正如你所知道的那样，这属于异端、邪教，不可饶恕！

# 稍微有点审美的人，绝对不允许中英文纯文本之间不留空格

但是，坦率的讲，我认为这应该被作为排版规范而不是根据个人习惯去选择。

事实上，Adobe InDesign、Microsoft Word 等对中文与西文（这里用「西文」来泛指用空格分词的外文）混合排版支持较好的软件，都默认增大汉字和西文的间距。

InDesign 的默认设置是 1/4 的全角空格宽度（遵从 JIS），也就是约等于一个半角空格。  

但大多数情况下我们没有这样专业排版软件的支持，只能手动控制，比如我编写这个 README.md 的时候。

# 本工具主要解决自动在东西方文字之间加空格的问题。

### 使用方法
```bash
$ go get github.com/chengyumeng/docfmt
$ docfmt --help
$ docfmt -p=/home/chengyumeng/docfmt 
```

### 特性
- 添加忽略文件
> 由于无法从技术上完全分辨文本文件和二进制文件，因此用户可以通过在 HOME 目录或者工作目录下增加一个命名为 .docfmtignore 的文件按行添加可忽略文件正则。
- 指定匹配规则
> 可以更暴力的通过 -m 参数指定匹配文件名的正则形式，只对规则内文件内容格式化
- 调试模式
> 通过 --debug 参数，用户可以不立刻对文档格式化，而是打印出需要格式化的内容，让用户确保格式化规则没有疏漏。

### TODO
- 新增更多中规则的格式化
- 提升性能，增加性能测试
