package controllers

import "sharemusic/models"

func (c *MusicController) LeftRoom() {
	id, _ := c.GetInt64("id")
	res, _ := models.GetOneById(id)
	var err error
	if res.UserNum == 1 {
		err = models.DelRoomInfo(id)
	} else {
		err = models.UpdateRoomInfoUserName(id, res.UserNum-1)
	}
	c.Ctx.Output.SetStatus(200)
	newBody := map[string]interface{}{
		"code": 200,
	}
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		newBody["code"] = 500
	}
	c.Ctx.Output.SetStatus(200)
	c.Data["json"] = newBody
	c.ServeJSON()
}
