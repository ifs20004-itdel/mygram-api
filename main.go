package main

import (
	"mygramapi/routers"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	routers.StartServer().Run(":8080")
}
