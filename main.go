package main

import (
	"bubble/database"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func main() {
	InitConfig()            //读取配置文件
	db := database.InitDB() //初始化数据库
	defer db.Close()
	r := gin.Default()
	r = CollectRoute(r)
	port := viper.GetString("server.port")
	r.Run(":" + port)
}
func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
