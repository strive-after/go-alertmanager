package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	model "webhook/module"
	"webhook/notifier"
)

var (
	//h            bool
	defaultRobot string
	//configpath  string
	//cfg *goconfig.ConfigFile
)

//func init() {
//	flag.BoolVar(&h, "h", false, "help")
//	flag.StringVar(&configpath, "config", "./config.ini", "config path")
//}

//func GetConfigIni(filepath string) (err error) {
//	config, err := goconfig.LoadConfigFile(filepath)
//	if err != nil {
//		fmt.Println("配置文件读取错误,找不到配置文件", err)
//		return err
//	}
//	cfg = config
//	return nil
//}


func main() {
	//flag.Parse()
	//
	//if h {
	//	flag.Usage()
	//	return
	//}
	//err := GetConfigIni(configpath)
	//if err != nil {
	//	fmt.Println(err,"配置文件路径错误")
	//	return
	//}
	//token ,err := cfg.GetValue("Ding","token")
	//if err != nil {
	//	fmt.Println(err,"token不存在")
	//	return
	//}
	token := os.Getenv("token")
	defaultRobot = "https://oapi.dingtalk.com/robot/send?access_token=" + token
	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)
		//fmt.Printf("%#v",notification)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": " successful receive alert notification message!"})

		err = notifier.Send(notification, defaultRobot)
		fmt.Println(err)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to dingtalk successful!"})
	})
	router.Run()
}