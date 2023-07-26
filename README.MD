## 基于 GoFrame v2.5.0 自定义项目模板以及代码生成工具

GoFrame官方 Quick Start:
- https://goframe.org/pages/viewpage.action?pageId=1114399

当前自定义模板依赖的CodeGen 工具版本: github.com/Jiaru0314/go_gen_code v1.1.0

### 前置准备 确保Golang环境已经安装、MySQL数据库以及业务表已经 准备好
###  Step1
    修改  `/hack/config.yaml` 文件中数据库连接信息 以及配置需要生成的表 可参考GoFrame官方文档

```yaml
gfcli:
  gen:
    dao:
      - link: "mysql:root:123456@tcp(127.0.0.1:3306)/your_database"
        tables: "document"
        descriptionTag: true
        noModelComment: true
```

### Step2
    执行/main.go 中的 gencode.GenALl()方法生成dao层以及业务代码,生成结束后 将此行代码删除 OR 注释

    生成结果示例:
```
2023/07/26 10:30:01 generated: internal\dao\document.go
2023/07/26 10:30:01 generated: internal\dao\internal\document.go
2023/07/26 10:30:01 generated: internal\model/do\document.go
2023/07/26 10:30:01 generated: internal\model/entity\document.go
[document] Dao 层代码生成完毕

2023/07/26 10:30:01 generated: ./api/document.go
2023/07/26 10:30:01 generated: ./internal/model/document.go
2023/07/26 10:30:01 generated: ./internal/controller/document.go
2023/07/26 10:30:01 generated: ./internal/logic/document/document.go
2023/07/26 10:30:01 generated: ./internal/service/document.go

2023/07/26 10:30:01 [document] 业务代码生成完毕

```

### Step3
    修改 /internal/cmd/cmd.go 文件 进行业务路由注册。

```go
package cmd

import (
	"context"
	"github.com/Jiaru0314/template-single/internal/consts"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.GET("/swagger", func(r *ghttp.Request) {
					r.Response.Write(consts.SwaggerUIPageContent)
				})

				// TODO Step3 路由绑定
				group.Bind(
					controller.Document,
					controller.SysOperaLog,
				)
			})
			s.Run()
			return nil
		},
	}
)

```

### Step4
    修改 /manifest/config/config.yaml 配置文件
```yaml
    server:
    address:     ":8000"
    openapiPath: "/api.json"

    logger:
      level : "all"
      stdout: true

    database:
      logger:
        level: "all"
        stdout: true
      default:
        link: "mysql:root:123456@tcp(127.0.0.1:3306)/your_database"
        debug: true
```

### Step5
    修改 /main.go 去除代码生成函数; 启动main函数;
```go

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
```
