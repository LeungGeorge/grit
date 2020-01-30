package hexo

import (
	"github.com/urfave/cli"
)

// Hexo blog 预览发布封装
var Hexo = cli.Command{
	Name:     "hexo",
	Usage:    "hexo blog commands: preview, publish.",
	Category: "hexo",
	Subcommands: []cli.Command{
		{
			Name:   "clean",
			Usage:  "clean hexo local files",
			Action: preview,
		},
		{
			Name:   "preview",
			Usage:  "preview hexo blog",
			Action: preview,
		},
		{
			Name:   "publish",
			Usage:  "publish hexo blog",
			Action: publish,
		},
	},
}
