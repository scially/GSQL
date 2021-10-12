package gsql

import (
	"encoding/json"
	"log"

	"os"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

type Route struct {
	Method   string
	Endpoint string
	Query    []string
}

func (r *Route) GetSQLResult(ctx *context.Context, db *orm.Ormer) {
	params := make([]string, len(r.Query)-1)
	for i := 1; i < len(r.Query); i++ {
		params[i-1] = ctx.Input.Query(r.Query[i])
	}

	var maps []orm.Params

	_, err := (*db).Raw(r.Query[0], params).Values(&maps)
	if err != nil {
		ctx.Output.SetStatus(400)
		ctx.WriteString(err.Error())
		return
	}

	ctx.Output.SetStatus(200)

	result := make(map[string][]orm.Params)
	result["result"] = maps
	data, _ := json.Marshal(result)
	ctx.WriteString(string(data))
}

type Database struct {
	Type       string
	Connection string
}

type DbRouter struct {
	Db     Database
	Routes []Route
}

type DbController struct {
	Method string
	Query  []string
}

type GSQLClient struct {
	db     *orm.Ormer
	Router DbRouter
}

func NewDbClient(config string) (GSQLClient, error) {
	var client GSQLClient
	file, err := os.Open(config)
	if err != nil {
		logs.Critical("Open file failed: %s \n", config)
		return client, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&client.Router); err != nil {
		logs.Critical("Parse file failed")
		return client, err
	}

	return client, nil
}

func (client *GSQLClient) Init(httpServer *web.HttpServer) error {
	if err := orm.RegisterDataBase("default", client.Router.Db.Type, client.Router.Db.Connection); err != nil {
		log.Fatalln(err.Error())
		return err
	}

	_db := orm.NewOrm()
	client.db = &_db

	for ind := 0; ind < len(client.Router.Routes); ind++ {
		router := client.Router.Routes[ind]

		switch client.Router.Routes[ind].Method {
		case "get", "Get", "GET":
			func(router Route) {
				httpServer.Get(router.Endpoint, func(ctx *context.Context) {
					logs.Info("[GET] %s", ctx.Request.RequestURI)
					router.GetSQLResult(ctx, client.db)
				})
			}(router)
		case "post", "Post", "POST":
			func(router Route) {
				httpServer.Post(router.Endpoint, func(ctx *context.Context) {
					logs.Info("[POST] %s", ctx.Request.RequestURI)
					router.GetSQLResult(ctx, client.db)
				})
			}(router)
		}
	}
	return nil
}
