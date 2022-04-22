package main

import (
	"flag"
	"fmt"

	"datac/internal/config"
	"datac/internal/handler"
	"datac/internal/svc"
	"github.com/zeromicro/go-zero/core/service"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/datac-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)

	handler.RegisterHandlers(server, ctx)

	group := service.NewServiceGroup()
	defer group.Stop()

	group.Add(server)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	group.Start()
}
