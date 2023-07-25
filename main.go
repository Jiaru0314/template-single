package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	_ "github.com/gogf/template-single/internal/packed"

	"github.com/gogf/template-single/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
