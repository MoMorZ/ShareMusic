package controllers

import (
	"encoding/json"
	"sharemusic/models/api"
)

// song/url?id=
func (c *ApiController) SongDetail() {
	ids := c.GetString("ids")
	query := map[string]interface{}{
		"ids": ids,
	}
	answer := api.SongDetail(query)
	body := answer["body"].(string)
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(body), &m)
	//fmt.Println(answer)
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = m
	c.ServeJSON()
}
