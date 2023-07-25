package main

import (
	"github.com/Jiaru0314/go_gen_code/gencode"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/template-single/internal/packed"
)

func main() {
	//代码生成后去除!!!
	gencode.GenALl()

	//cmd.Main.Run(gctx.GetInitCtx())
}
