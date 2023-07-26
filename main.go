package main

import (
	"github.com/Jiaru0314/go_gen_code/codeGenUtil"
	"github.com/Jiaru0314/template-single/internal/cmd"
	_ "github.com/Jiaru0314/template-single/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	//代码生成后去除!!!
	codeGenUtil.GenALl()

	cmd.Main.Run(gctx.GetInitCtx())
}
