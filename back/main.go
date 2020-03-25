package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
	"sharemusic/controllers"

	_ "sharemusic/routers"
)

func init() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/sharemusic?charset=utf8")
}

// 1914998204
func main() {
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		AllowCredentials: true,
	}))
	beego.Router("/song/url:id", &controllers.ApiController{}, "get:SongUrl")
	beego.Router("/song/detail:ids", &controllers.ApiController{}, "get:SongDetail")
	beego.Router("/sharemusic/list", &controllers.MusicController{}, "get:GetListDetail")
	beego.Router("/sharemusic/number", &controllers.MusicController{}, "get:GetNumber")
	beego.Router("/sharemusic/heart", &controllers.MusicController{}, "get:Heart")
	beego.Router("/sharemusic/left:id", &controllers.MusicController{}, "get:LeftRoom")
	beego.Run()
}
