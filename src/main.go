package main

import (
	"flag"
	"fmt"

	"hassh/src/internal/config"
	"hassh/src/internal/handler"
	"hassh/src/internal/middleware"
	"hassh/src/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

func main() {
	ctx := initConfig()
	c := ctx.Config
	server := rest.MustNewServer(c.RestConf)
	logx.MustSetup(c.Log)
	defer server.Stop()

	svc.InitComponents(ctx)
	svc.InitQueues(ctx)
	handler.RegisterHandlers(server, ctx)
	middleware.ErrorHandling()

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	fmt.Printf("Admin Web Page: http://localhost:%d/admin.html\n", c.Port)
	server.Start()
}

func initConfig() *svc.ServiceContext {
	var configFile = flag.String("f", "etc/sshtask.yaml", "the config file")
	var configJson = flag.String("jconfig", "etc/componentConfig.json", "the config file")

	flag.Parse()

	var c config.Config
	var j svc.CustomConfigStruct
	conf.MustLoad(*configFile, &c)
	conf.Load(*configJson, &j)

	ctx := svc.NewServiceContext(c, j)
	return ctx
}
