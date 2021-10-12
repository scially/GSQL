package main

import (
	"os"

	"cn.gzpi/gsql/gsql"

	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	Config string `short:"f" long:"config" required:"true" description:"config file" `
	Port   int    `short:"p" long:"port" description:"port" `
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)

	if _, err := parser.Parse(); err != nil {
		parser.WriteHelp(os.Stdout)
		logs.Critical(parser.Usage)
		return
	}

	httpServer := web.NewHttpSever()
	httpServer.Cfg.Listen.HTTPPort = opts.Port

	client, err := gsql.NewDbClient(opts.Config)
	if err != nil {
		logs.Critical("Client init error [%s]", err)
		return
	}

	client.Init(httpServer)
	httpServer.InsertFilter("*", web.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))

	httpServer.Run("0.0.0.0")

}
