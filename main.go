package main

import (
	cfg "auto-installation/config"
	router "auto-installation/handle"
	"auto-installation/repository"
	itasks "auto-installation/task"
	"fmt"
	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
	"os"
	"sort"
)

var (
	server *machinery.Server
	cnf    *config.Config
	app    *cli.App
	tasks  map[string]interface{}
)

func init() {

	cfg.InitConfig()

	//app := cli.NewApp()
	//app.Flags = []cli.Flag{
	//	cli.StringFlag{
	//		Name:  "c",
	//		Usage: "",
	//		Value: "Path to a configuration file",
	//	},
	//}

	tasks = map[string]interface{}{
		"add":      itasks.Add,
		"multiply": itasks.Multiply,
	}

	cnf, err := loadConfig(cfg.Cfg.ConfigPath)
	if err != nil {
		panic(err)
	}

	server, err = machinery.NewServer(cnf)
	if err != nil {
		// do something with the error
		fmt.Println(err.Error())
		panic(err)
	}

	//init result-backend client
	redis.InitRedisClient()
}

func main() {
	app := &cli.App{
		Commands: []cli.Command{
			{
				Name:  "worker",
				Usage: "launch machinery worker",
				Action: func(c *cli.Context) error {
					if err := runWorker(); err != nil {
						return cli.NewExitError(err.Error(), -1)
					}
					return nil
				},
			},
			{
				Name:  "sender",
				Usage: "send async tasks",
				Action: func(c *cli.Context) error {
					if err := runSender(); err != nil {
						return cli.NewExitError(err.Error(), -1)
					}
					return nil
				},
			},
		},
	}
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Run(os.Args)
	//app.Run([]string{"worker"})
}

func loadConfig(configPath string) (*config.Config, error) {
	if configPath != "" {
		return config.NewFromYaml(configPath, true)
	}

	return config.NewFromEnvironment()
}

func runWorker() (err error) {
	if err = server.RegisterTasks(tasks); err != nil {
		fmt.Println(err)
		panic(err)
		return
	}

	workers := server.NewWorker("worker_test", 10)
	if err = workers.Launch(); err != nil {
		panic(err.Error())
		return
	}
	return
}

// 对外提供接口
func runSender() (err error) {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/add", func(c *gin.Context) {
		router.Add(c, server)
	})

	r.GET("/add_chain", func(c *gin.Context) {
		router.AddChain(c, server)
	})

	err = r.Run(fmt.Sprintf(":%d", cfg.Cfg.AppPort))
	return
}

func startServer() (err error) {
	server, err = machinery.NewServer(cnf)
	if err != nil {
		return
	}

	err = server.RegisterTasks(tasks)
	return
}
