package test

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"sharemusic/models"
	"testing"
)

func TestGetAll(t *testing.T) {
	initSql()
	models.GetAll()
}

func TestGetOneById(t *testing.T) {
	initSql()
	fmt.Println(models.GetOneById(1258875077))
}

func TestDelRoomInfo(t *testing.T) {
	initSql()
	var id int64
	id = -660075787
	fmt.Println(models.GetOneById(id))
	models.DelRoomInfo(id)
}

func initSql() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/sharemusic?charset=utf8")
}
