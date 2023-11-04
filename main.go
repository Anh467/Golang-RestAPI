package main

import (
	"encoding/json"
	"log"
	"main/common"
	"main/entities"
	"main/routers"

	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var configPath string = "config.json"

func setConfigPath() {
	path, err := common.GetRootProject(configPath)
	if err != nil {
		fmt.Println("err: " + err.Error())
	} else {
		fmt.Println("path: " + path)
	}
	configPath = path
}

func getConfigJson(path string) entities.Config {

	file, err := common.GetFile(path)

	var config entities.Config
	if err != nil {
		fmt.Print("[main/main.go/getConfigJson(path string)]: " + err.Error())
		return config
	}

	fileContent, err := common.GetFileContents(file)

	if err != nil {
		fmt.Print("[main/main.go/getConfigJson(path string)]: " + err.Error())
		return config
	}
	errr := json.Unmarshal([]byte(fileContent), &config)
	if errr != nil {
		fmt.Print("[main/main.go/getConfigJson(path string)]: " + err.Error())
		return config
	}
	return config
}

func main() {

	setConfigPath()
	configJson := getConfigJson(configPath)
	fmt.Print(configJson)
	//
	db, err := ConnectSqlServerGorm(configJson.SqlServer)
	if err != nil {
		fmt.Print("[main/main.go]: " + err.Error())
		return
	}
	//
	gin.SetMode(gin.ReleaseMode)
	//gin
	r := gin.Default()
	// setting
	r.Use(cors.New(cors.Config{
		AllowHeaders:    []string{"*"},
		AllowAllOrigins: true,
	}))
	//
	r.ForwardedByClientIP = true
	r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.2", "10.0.0.0/8"})
	// middleware recover

	// router
	routers.V1Router(r, db)
	// run
	r.Run()
}

func ConnectSqlServerGorm(Config entities.SqlServer) (*gorm.DB, error) {
	connString := fmt.Sprintf("server=localhost; user id=%s; password=%s; database=%s;",
		Config.User, Config.Pass, Config.DB)
	db, err := gorm.Open(sqlserver.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Print("connection wrong " + "\n")
		log.Fatal(err)
	}
	// Kiểm tra kết nối
	err = db.Error
	if err != nil {
		fmt.Print("check connection wrong ")
		log.Fatal(err)
	}
	fmt.Println("Kết nối thành công!")
	return db, err
}

func ConnectMySQLGorm(Config entities.MySql) (*gorm.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.User, Config.Pass, Config.DB)
	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		fmt.Print("connection wrong " + "\n")
		log.Fatal(err)
	}
	// Kiểm tra kết nối
	err = db.Error
	if err != nil {
		fmt.Print("check connection wrong ")
		log.Fatal(err)
	}
	fmt.Println("Kết nối thành công!")
	return db, err
}
