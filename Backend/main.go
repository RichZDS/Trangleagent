package main

import (
	_ "TrangleAgent/internal/packed"

	_ "TrangleAgent/internal/logic"

	"TrangleAgent/internal/cmd"

	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
