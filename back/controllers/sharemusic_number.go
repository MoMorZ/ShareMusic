package controllers

import (
	"encoding/json"
	"sharemusic/models"
	"sharemusic/models/api"
)

func (c *MusicController) GetNumber() {
	id, _ := c.GetInt64("id")
	res, _ := models.GetOneById(id)
	models.UpdateRoomInfoUserName(id, res.UserNum+1)
	query := map[string]interface{}{"id": res.ListNum}
	answer := api.PlayListDetail(query)
	body := answer["body"].(string)
	m := make(map[string]interface{})
	_ = json.Unmarshal([]byte(body), &m)
	newBody := map[string]interface{}{
		"list":   m["privileges"],
		"code":   m["code"],
		"number": id,
		"mstate": res.State,
		"nownum": res.NowNum,
	}
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = newBody
	c.ServeJSON()
}
