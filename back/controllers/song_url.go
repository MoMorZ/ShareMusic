package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"sharemusic/models/api"
)

type ApiController struct {
	beego.Controller
}

// song/url?id=
func (c *ApiController) SongUrl() {
	id := c.GetString("id")
	query := map[string]interface{}{
		"id": id,
	}
	answer := api.SongUrl(query)
	body := answer["body"].(string)
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(body), &m)
	//fmt.Println(answer)
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = m
	c.ServeJSON()
}
