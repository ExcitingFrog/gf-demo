package dao

import (
	"fmt"
	"gf-init/app/model"

	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	_ "github.com/lib/pq"
)

var DB gdb.DB

func GetUsers() {
	var user model.Users
	DB = g.DB("default")
	DB.Table("users").Where("nickname = ?", "1").Scan(&user)
	fmt.Println(user)
}
