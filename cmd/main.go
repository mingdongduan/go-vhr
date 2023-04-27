package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"go-vhr/router"
	"go-vhr/tools"
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
	router.InitRouter(engine)
	engine.Run(":9000")
}
