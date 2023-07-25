package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"template-single/internal/cmd"
	_ "template-single/internal/packed"
)

func main() {
	//代码生成后去除!!!
	//gencode.GenALl()

	cmd.Main.Run(gctx.GetInitCtx())
}
