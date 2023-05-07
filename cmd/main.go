package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-vhr/api/router"
	"go-vhr/pkg/tools"
)

var (
	port   string
	daemon bool
)
var serverCmd = &cobra.Command{
	Use:     "server",
	Short:   "启动http服务",
	Example: "go-fly server -c config/",
	Run: func(cmd *cobra.Command, args []string) {
		main()
	},
}

func main() {
	fmt.Println("hello world")
	tools.Logger().Infoln("this is a log")
	engine := gin.Default()
	engine.Use(tools.Session("go-vhr"))
	router.InitRouter(engine)
	engine.Run(":9000")
}
