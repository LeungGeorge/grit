package main

import (
	"log"
	"os"
	"sort"

	"github.com/LeungGeorge/grit/after"
	"github.com/LeungGeorge/grit/before"
	"github.com/LeungGeorge/grit/command/git"
	"github.com/LeungGeorge/grit/command/hexo"
	"github.com/LeungGeorge/grit/command/space"
	"github.com/LeungGeorge/grit/command/update"
	"github.com/LeungGeorge/grit/flag"
	"github.com/urfave/cli"
)

func main() {
	// 1. 创建 APP
	// 通过 cli.NewApp() 创建一个实例
	app := cli.NewApp()

	// 2. 基础配置
	// 配置 APP 的一些属性、动作，包括 name，usage 等等。
	app.Name = "grit"
	app.Usage = "grit is a tool of grimoire, batch execute commands."
	app.Version = "1.0.0"

	// 2.1
	// 配置 flags，一些公用变量标识，供后续逻辑（比如 action 中）使用
	// flag 信息是按照声明顺序打印的，如果想按照字典序打印，可以在 flags 声明后加上如下代码：
	// import "sort"
	// sort.Sort(cli.FlagsByName(app.Flags))
	// 以下为测试代码，仅做说明之用。
	app.Flags = []cli.Flag{
		flag.FlagMessage,
	}
	sort.Sort(cli.FlagsByName(app.Flags))

	// 2.2
	// app.Action = XXX

	// 2.3. Command 配置
	// 命令行程序除了有 flag 之外，还有 command（比如 git log, git commit 等等）。
	// 每个 command 对应一个 cli.Command 接口实例，入口函数通过 Action 指定。
	// 如果想在帮助信息里面实现分组展示，可通过设置 Category 实现。
	// 另外，command 可能还有 subcommand 命令，这就需要添加两个命令才能完成相应的操作。
	// 举个栗子：
	// grit git sync
	// 可使用 grit git -h 查看帮助信息

	app.Commands = []cli.Command{
		git.Git,
		hexo.Hexo,
		update.Update,
		space.Space,
	}
	sort.Sort(cli.CommandsByName(app.Commands))

	// 2.4
	// 如果想在 Action 之前/后 执行一些操作，可以通过 Before、After 实现。
	app.Before = before.Action
	app.After = after.Action

	// 3 总结：
	// github.com/urfave/cli 库只需关注 Action 实现逻辑即可实现一个命令行工具，简单易用。
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
